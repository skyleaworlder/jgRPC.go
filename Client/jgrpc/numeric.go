package jgrpc

import prtco "github.com/skyleaworlder/jgRPC.go/jgproto"

func add(c *Calculator, a, b int) int {
	const FuncName = "add"
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

	// send request
	buf, _ = connect(c.Config, buf)

	// process response
	_, _, res := prtco.ParseResponse(buf)
	return res[0].(int)
}
