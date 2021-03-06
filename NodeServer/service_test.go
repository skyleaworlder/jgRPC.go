package main

import (
	"fmt"
	"testing"

	"github.com/skyleaworlder/jgRPC.go/jgproto"
	jgrpcs "github.com/skyleaworlder/jgRPC.go/nodeserver/jgrpc"
)

func Test_calcuReturnPart(t *testing.T) {
	calcu := new(jgrpcs.Calculator)
	calcu.Init()

	var a int8 = 2
	var b int8 = 5
	tlv := []jgproto.TLV{*jgproto.ConstructTLV(a), *jgproto.ConstructTLV(b)}

	res := calcuReturnPart(calcu, "AddInt8", 2, tlv)

	resp := jgproto.ConstructResponse(123)
	fmt.Println(res)
	resp.SetReturnNum(1)
	resp.SetReturnPart(res)
	fmt.Println("resp:", resp)
	fmt.Println("resp(byte):", resp.ComposeResponse())
	for _, v := range res {
		_, _, r := jgproto.ParseTLV(v.ComposeTLV())
		fmt.Println(r)
	}
}
