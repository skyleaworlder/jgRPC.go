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
	sm "github.com/umpc/go-sortedmap"
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

func discovery(iptable *sm.SortedMap) bool {
	resp := prtco.ConstructRequest()
	// NS(3)
	resp.SetType(0x03)

	msg := resp.ComposeRequest()
	buf, err := jgut.Dial(Config["NS_Addr"], msg)
	if err != nil {
		msg := "Warning: Service Discovery Failed\n"
		fmt.Fprint(os.Stderr, msg+err.Error())
		return false
	}

	// transfer []bytes from []TLV to []interface{}
	// ip is string
	_, _, ips := prtco.ParseResponse(buf)
	for _, v := range ips {
		ip := v.(string)
		iptable.Insert(ip, ip)
	}
	return true
}

// compare function used in SortedMap
func ipComp(i, j interface{}) bool {
	return i.(string) < j.(string)
}
