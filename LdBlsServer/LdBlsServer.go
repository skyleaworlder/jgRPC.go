package main

import (
	"fmt"
	"net"
	"os"
	"time"

	conhash "github.com/skyleaworlder/jgRPC.go/LdBlsServer/ConsistHash"
	jgut "github.com/skyleaworlder/jgRPC.go/jgrpcUtils"
	sm "github.com/umpc/go-sortedmap"
)

// LdBlsServer should implement several functions:
// 1. communicate with client directly
//	accept CALL(0) request, and send a response to client
//
// 2. get ip/port from Name Server
//	send NS(3), receive OK(1)
//
// 3. communicate with Node Server
// 	send CALL(0), receive data-OK(1) response

const (
	cfgMaxNum = 16
	ipMaxNum  = 16
)

var (
	// Config is a map
	// e.g. "Local_IP" => "127.0.0.1"
	Config  = make(map[string]string, cfgMaxNum)
	iptable = sm.New(ipMaxNum, ipComp)
)

func main() {
	// init LdBlsServer
	jgut.Readcfg(Config, "LdBlsServer.cfg")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", Config["Listen_Port"])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %s", err.Error())
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %s", err.Error())
		os.Exit(1)
	}

	// service discovery
	for !discovery(iptable) {

	}

	// init conhash.HT
	ht := conhash.InitConHash(iptable, Config)
	fmt.Println(ht)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal Error: %s", err.Error())
			continue
		}
		go handleRequest(ht, conn)
	}
}

func handleRequest(ht *conhash.Conhash, conn net.Conn) {
	conn.SetDeadline(time.Now().Add(5 * time.Second))
	defer conn.Close()

	buf := make([]byte, 256)
	_, err := conn.Read(buf)
	if err != nil {
		msg := "Warning: LdBls Server handleRequest failed\n"
		fmt.Fprint(os.Stderr, msg)
		return
	}

	resp := make([]byte, 256)
	resp, _ = getResponse(ht, buf)
	conn.Write(resp)
}
