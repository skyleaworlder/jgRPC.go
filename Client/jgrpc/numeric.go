package jgrpcclient

import (
	"errors"
	"fmt"
	"os"

	prtco "github.com/skyleaworlder/jgRPC.go/jgproto"
	jgut "github.com/skyleaworlder/jgRPC.go/jgrpcUtils"
)

func add(c *Calculator, a, b int8) (int8, error) {
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
	buf, err := jgut.Dial(c.Config["TS_Addr"], buf)
	// fmt.Println("response, buf:", buf)

	// process response
	if buf == nil {
		msg := "Fatal Error: Dial failed, buf is nil\n"
		fmt.Fprint(os.Stderr, msg+err.Error())
		return 0, errors.New(msg)
	}
	_, _, res := prtco.ParseResponse(buf)
	return res[0].(int8), nil
}
