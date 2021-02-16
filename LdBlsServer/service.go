package main

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"

	conhash "github.com/skyleaworlder/jgRPC.go/LdBlsServer/ConsistHash"
	prtco "github.com/skyleaworlder/jgRPC.go/jgproto"
	jgut "github.com/skyleaworlder/jgRPC.go/jgrpcUtils"
)

// LdBlsServer should implement several functions:
// 1. communicate with client directly
//	accept CALL(0) request, and send a response to client
//
// 2. get ip/port from Name Server
//	send NS(3), receive NS(3)
//
// 3. communicate with Node Server
// 	send CALL(0), receive OK(254) response

func getResponse(ht *conhash.Conhash, msg []byte) ([]byte, error) {
	req := prtco.ParseRequest(msg)
	Type := req.GetType()

	// make the result of sha1.Sum addressable
	sha1Sum := sha1.Sum(msg)
	HID := binary.BigEndian.Uint64(sha1Sum[:][:8])

	// GetNode
	node, err := ht.GetNode(HID)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return []byte{}, err
	}

	switch Type {
	// request CALL(0)
	case 0x00:
		addr := node.IP.String() + ":" + strconv.Itoa(int(node.PORT))
		fmt.Println("ldbls.service.go/getResponse:", addr)
		buf, _ := jgut.Dial(addr, msg)
		return buf, nil
	default:
		fmt.Println("Warning: ldbls.service.go/getResponse shouldn't go here")
		return []byte{}, err
	}
}
