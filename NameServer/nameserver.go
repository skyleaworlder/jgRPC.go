package main

import (
	"fmt"
	"net"
	"os"
	"time"

	jgut "github.com/skyleaworlder/jgRPC.go/jgrpcUtils"
	sm "github.com/umpc/go-sortedmap"
)

// Name Server should implement several functions:
// 1. health probe
//	send HELLO(3) to Node Server,
// 	then receive HELLO(3) response from Node Server.
//
// 2. service discovery
//	receive NS(3) from load balance/cache Server.
//  send ip/port informations to load balance/cache Server.

const (
	cfgMaxNum  = 16
	nameMaxNum = 16
)

var (
	// Config is a map to store configuration
	Config = make(map[string]string, cfgMaxNum)
	// ip hash table (
	iptable = sm.New(nameMaxNum, ipComp)
)

func main() {
	jgut.Readcfg(Config, "NameServer.cfg")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", Config["Listen_Port"])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %s", err.Error())
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: Listen failed. %s", err.Error())
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err.Error())
			continue
		}
		go handleRequest(iptable, conn)
	}
}

func handleRequest(iptable *sm.SortedMap, conn net.Conn) {
	conn.SetDeadline(time.Now().Add(5 * time.Second))
	defer conn.Close()

	buf := make([]byte, 256)
	_, err := conn.Read(buf)
	if err != nil {
		msg := "Warning: Name Server handleRequest failed\n"
		fmt.Fprintf(os.Stderr, msg+"%s", err.Error())
		return
	}

	resp := make([]byte, 256)

	// for debug
	// fmt.Print(buf)
	resp, _ = getResponse(iptable, buf)
	conn.Write(resp)
}
