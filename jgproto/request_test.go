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

func Test_ConstructRequest(t *testing.T) {
	fmt.Println("3.3 jgproto.Test_ConstructRequest:")
	defreq := ConstructRequest()
	fmt.Printf("Magic: %x\n", defreq.GetMagic())
	fmt.Printf("CID: %d\n", defreq.GetCID())
	fmt.Printf("Type: %x\n", defreq.GetType())
	fmt.Printf("Param_num: %d\n", defreq.GetParamNum())
	fmt.Printf("Length: %d\n", defreq.GetLength())
	fmt.Printf("Src_addr: %x\n", defreq.GetSrcAddr())
	fmt.Printf("Func_name: %s\n", defreq.GetFuncName())
}

func Test_ComposeRequest(t *testing.T) {
	fmt.Println("3.4 jgproto.Test_ComposeRequest:")

	defreq := ConstructRequest()
	defreq.SetFuncName("add")
	defreq.SetParamNum(2)
	defreq.SetLength(7)

	var a uint64 = 5
	var b uint64 = 6
	atlv, btlv := ConstructTLV(a), ConstructTLV(b)
	defreq.SetParamPart([]TLV{*atlv, *btlv})

	bytarr := defreq.ComposeRequest()
	fmt.Println("[]byte of request:", bytarr)
}

func Test_ParseRequest(t *testing.T) {
	fmt.Println("3.5 jgproto.Test_ParseRequest:")

	defreq := ConstructRequest()
	defreq.SetFuncName("add")
	defreq.SetParamNum(2)
	defreq.SetLength(7)

	var a uint64 = 5
	var b uint64 = 6
	atlv, btlv := ConstructTLV(a), ConstructTLV(b)
	defreq.SetParamPart([]TLV{*atlv, *btlv})

	bytarr := defreq.ComposeRequest()
	fmt.Println("[]byte of request:", bytarr)

	defreq = ParseRequest(bytarr)
	fmt.Printf("Magic: %x\n", defreq.GetMagic())
	fmt.Printf("CID: %d\n", defreq.GetCID())
	fmt.Printf("Type: %x\n", defreq.GetType())
	fmt.Printf("Param_num: %d\n", defreq.GetParamNum())
	fmt.Printf("Length: %d\n", defreq.GetLength())
	fmt.Printf("Src_addr: %x\n", defreq.GetSrcAddr())
	fmt.Printf("Func_name: %s\n", defreq.GetFuncName())
	fmt.Printf("Param_Part: %x\n", defreq.GetParamPart())
}
