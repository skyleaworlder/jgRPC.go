package jgrpc

import (
	"fmt"
	"net"
	"os"
)

func connect(cfg map[string]string, msg []byte) ([]byte, error) {
	conn, err := net.Dial("tcp4", cfg["NS_Addr"])
	if err != nil {
		fmt.Fprint(os.Stderr, "Warning: jgrpc.send, net.Dial failed\n")
		return []byte{}, err
	}
	defer conn.Close()

	// Call
	_, err = conn.Write(msg)
	if err != nil {
		fmt.Fprint(os.Stderr, "Warning: jgrpc.send, conn.Write failed\n")
		return []byte{}, err
	}

	buf := make([]byte, 256)
	// wait
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Fprint(os.Stderr, "Warning: jgrpc.send, conn.Read failed\n")
		return []byte{}, err
	}

	return buf, nil
}
