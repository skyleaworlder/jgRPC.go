package main

import (
	"fmt"
	"net"
	"os"
	"time"

	jgut "github.com/skyleaworlder/jgRPC.go/jgrpcUtils"
	jgrpcs "github.com/skyleaworlder/jgRPC.go/nodeserver/jgrpc"
)

// NodeServer should implement several functions:
// 1. service register
//	send REGISTER(1) to Name Server
//
// 2. health probe
//	receive HELLO(3) from Name Server,
// 	then send HELLO(3) response to Name Server.
//
// 3. communicate with load balance/cache Server
// 	receive CALL(0) request, then send execute results.

const (
	cfgMaxNum = 16
)

var (
	// Config is a map
	// e.g. "Local_IP" => "127.0.0.1"
	Config = make(map[string]string, cfgMaxNum)

	// Calcu is the obj of calculator
	Calcu = new(jgrpcs.Calculator)
)

func main() {
	jgut.Readcfg(Config, "NodeServer.cfg")
	Calcu.Init()
	Calcu.Config = Config

	listener, err := net.Listen("tcp4", Config["Listen_Port"])
	if err != nil {
		msg := "Warning: Node Server listener initialization failed\n"
		fmt.Fprint(os.Stderr, msg)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			msg := "Warning: Node Server listener failed\n"
			fmt.Fprint(os.Stderr, msg)
			continue
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// set 5s ddl
	conn.SetDeadline(time.Now().Add(5 * time.Second))
	defer conn.Close()

	buf := make([]byte, 256)
	_, err := conn.Read(buf)
	if err != nil {
		msg := "Warning: Node Server handleRequest failed\n"
		fmt.Fprint(os.Stderr, msg)
		return
	}

	resp := make([]byte, 256)
	resp = getResponse(buf)
	conn.Write(resp)
}
