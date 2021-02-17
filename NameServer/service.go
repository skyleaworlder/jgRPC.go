package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strconv"

	prtco "github.com/skyleaworlder/jgRPC.go/jgproto"
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

func getResponse(iptable *sm.SortedMap, msg []byte) ([]byte, error) {
	req := prtco.ParseRequest(msg)
	CID := req.GetCID()
	Type := req.GetType()

	switch Type {
	// REGISTER(1)
	case 0x01:
		// Parse from SrcAddr([]byte), port(uint16) to a.b.c.d:port
		req := prtco.ParseRequest(msg)

		SrcAddr := make([]byte, 4)
		binary.BigEndian.PutUint32(SrcAddr, req.GetSrcAddr())

		// only 1 tlv in REGISTER(1) request
		tlv := req.GetParamPart()[0]
		_, _, val := prtco.ParseTLV(tlv.ComposeTLV())
		port := val.(uint16)

		ip := net.IPv4(SrcAddr[0], SrcAddr[1], SrcAddr[2], SrcAddr[3])
		addr := ip.String() + ":" + strconv.Itoa(int(port))

		fmt.Println("ip:", ip, "; port:", port, "; addr:", addr)
		if iptable.Has(addr) {
			break
		}
		iptable.Insert(addr, addr)
		resp := prtco.ConstructResponse(CID)
		return resp.ComposeResponse(), nil

	// NS(3)
	case 0x03:
		iter, _ := iptable.IterCh()
		defer iter.Close()

		tlvs := []prtco.TLV{}
		var cnt uint8 = 0
		for rec := range iter.Records() {
			ip := rec.Val.(string)
			tlv := prtco.ConstructTLV(ip)
			tlvs = append(tlvs, *tlv)
			cnt++
		}
		resp := prtco.ConstructResponse(CID)
		resp.SetReturnNum(cnt)
		resp.SetReturnPart(tlvs)

		// for debug
		fmt.Println("NS resp:", resp.ComposeResponse())
		return resp.ComposeResponse(), nil

	default:
		msg := "Fatal Error: NameServer failed\n"
		return nil, errors.New(msg)
	}
	return nil, nil
}

// compare function used in SortedMap
func ipComp(i, j interface{}) bool {
	return i.(string) < j.(string)
}
