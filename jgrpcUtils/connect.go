package jgrpcutils

import (
	"fmt"
	"net"
	"os"
)

// Dial is a function, using cfg to send msg
func Dial(addr string, msg []byte) ([]byte, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		fmt.Fprint(os.Stderr, "Warning: jgrpc.Dial, net.ResolveTCPAddr failed\n")
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprint(os.Stderr, "Warning: jgrpc.Dial, net.Dial failed\n")
		return nil, err
	}
	defer conn.Close()

	// Call
	_, err = conn.Write(msg)
	if err != nil {
		fmt.Fprint(os.Stderr, "Warning: jgrpc.Dial, conn.Write failed\n")
		return nil, err
	}

	buf := make([]byte, 256)
	// wait
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Fprint(os.Stderr, "Warning: jgrpc.Dial, conn.Read failed\n")
		return nil, err
	}

	return buf, nil
}
