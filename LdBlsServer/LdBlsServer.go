package main

import (
	"fmt"
	"net"
	"os"
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

func handleRequest(conn net.Conn) {
	defer conn.Close()
}

func main() {
	addr := "127.0.0.1:30000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %s", err.Error())
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %s", err.Error())
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal Error: %s", err.Error())
			continue
		}
		go handleRequest(conn)
	}
}
