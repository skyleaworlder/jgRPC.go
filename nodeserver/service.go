package main

import (
	prtco "github.com/skyleaworlder/jgRPC.go/jgproto"
	jgrpcs "github.com/skyleaworlder/jgRPC.go/nodeserver/jgrpc"
)

// NodeServer should implement several functions:
// 1. service register
//	send REGISTER(1) to Name Server
//
// 2. health probe
//	receive HELLO(3) from Name Server,
// 	then send HELLO(3) response to Name Server.
//
// 3. communicate with load balance/cache Server
// 	receive CALL(0) request, then send execute results.

func getResponse(msg []byte) []byte {
	// GC might work after re-assign *Request?
	req := new(prtco.Request)
	req = prtco.ParseRequest(msg)

	CID := req.GetCID()
	var resp *prtco.Response
	switch req.GetType() {
	// HELLO(3)
	case 0x03:
		resp = prtco.ConstructResponse(CID)
		resp.SetType(0x03)

	// CALL(0)
	case 0x00:
		resp = prtco.ConstructResponse(CID)
		ReturnPart := calcuReturnPart(Calcu, req.GetFuncName(), req.GetParamNum(), req.GetParamPart())
		// send OK(254)
		req.SetType(0xfe)
		resp.SetReturnNum(uint8(len(ReturnPart)))
		resp.SetReturnPart(ReturnPart)
	}

	retMsg := resp.ComposeResponse()
	return retMsg
}

// a function that calls RPC.Call, provides parameters.
func calcuReturnPart(in jgrpcs.RPC, FuncName string, ParamNum uint8, ParamPart []prtco.TLV) []prtco.TLV {
	params := make([]interface{}, ParamNum)
	// for debug
	// fmt.Println("param tlv part:", ParamPart)
	for i, param := range ParamPart {
		_, _, val := prtco.ParseTLV(param.ComposeTLV())
		params[i] = val
	}
	res := in.Call(FuncName, params...)
	tlvs := make([]prtco.TLV, len(res))
	for i, v := range res {
		tlvs[i] = *prtco.ConstructTLV(v.Interface())
	}

	return tlvs
}
