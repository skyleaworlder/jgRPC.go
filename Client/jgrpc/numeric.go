package jgrpcclient

import (
	prtco "github.com/skyleaworlder/jgRPC.go/jgproto"
	jgut "github.com/skyleaworlder/jgRPC.go/jgrpcUtils"
)

func add(c *Calculator, a, b int8) int8 {
	const FuncName = "AddInt8"
	req := prtco.ConstructRequest()

	// generate request
	req.SetSrcAddrStr(c.Config["Local_IP"])
	req.SetFuncName(FuncName)
	req.SetLength(uint16(4 + len(FuncName)))
	req.SetParamNum(2)

	atlv, btlv := prtco.ConstructTLV(a), prtco.ConstructTLV(b)
	tlv := []prtco.TLV{*atlv, *btlv}
	req.SetParamPart(tlv)

	buf := req.ComposeRequest()

	// fmt.Println("request, buf:", buf)
	// send request and get result in buf
	buf, _ = jgut.Dial(c.Config["TS_Addr"], buf)
	// fmt.Println("response, buf:", buf)

	// process response
	_, _, res := prtco.ParseResponse(buf)
	return res[0].(int8)
}
