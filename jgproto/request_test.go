package jgproto

import (
	"fmt"
	"testing"
)

func Test_proto(t *testing.T) {
	req := Request{
		Magic: 0x4A47, CID: 1,
		Type: 1, ParamNum: 0,
		Length: 9,
		// 192.168.1.1
		SrcAddr:  0xc0a80101,
		FuncName: "Hello",
	}

	fmt.Println(req.GetType())
}

func Test_parseIPv4(t *testing.T) {
	fmt.Println("ip is:", parseIPv4("192.168.1.1"), 0xc0a80101)
}
