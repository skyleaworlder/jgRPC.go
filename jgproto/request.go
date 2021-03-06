package jgproto

import (
	"crypto/rand"
	"encoding/binary"
	"math/big"
	"strconv"
	"strings"
)

// Request is a struct
//                                  1  1  1  1  1  1
//    0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//  |           J           |           G           |
//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//  |                      CID                      |
//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//  |          TYPE         |       PARAM_NUM       |
//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//  |                     LENGTH                    |
//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//  |                    SRC_ADDR                   |
//  |                                               |
//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//  |                                               |
//  /                    FUNC_NAME                  /
//  |                                               |
//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//  |                   PARAM_PART                  |
type Request struct {
	Magic     uint16
	CID       uint16
	Type      uint8
	ParamNum  uint8
	Length    uint16
	SrcAddr   uint32
	FuncName  string
	ParamPart []TLV
}

// ConstructRequest is a default construct function
//
// Magic: JG
//
// CID: have generated using crypto/rand in Constructor
//
// Type: default CALL(0)
//
// ParamNum: default 0
//
// Length: default 4(only Src_Addr)
//
// SrcAddr: default 127.0.0.1
//
// FuncName: default ""
//
// ParamPart default []TLV{}
func ConstructRequest() *Request {
	// to generate CID
	const UINT16MAX = ^uint16(0)
	req := new(Request)

	// set MAGIC charactor JG
	req.SetMagic(0x4A47)

	// generate a true random number as CID
	rd, _ := rand.Int(rand.Reader, big.NewInt(int64(UINT16MAX)))
	CID := uint16(rd.Uint64())
	req.SetCID(CID)

	// default type is CALL(0), and void
	req.SetType(0)
	req.SetParamNum(0)

	// Length means len({ SRC_ADDR, FUNC_NAME })
	// SRC_ADDR + FUNC_NAME  = 4
	//    4     + 0(default) = 4
	req.SetLength(4)
	// set SRC_ADDR 127.0.0.1
	req.SetSrcAddrStr("127.0.0.1")
	// default FUNC_NAME is ""
	req.SetFuncName("")

	// default PARAM_PART is nil
	req.SetParamPart([]TLV{})
	return req
}

// ComposeRequest is a method to generate a request([]byte)
func (req *Request) ComposeRequest() (res []byte) {

	Magic := make([]byte, 2)
	CID := make([]byte, 2)
	TypeParamNum := make([]byte, 2)
	Length := make([]byte, 2)
	SrcAddr := make([]byte, 4)
	FuncName := []byte(req.FuncName)

	Type := uint16(req.GetType())
	ParamNum := uint16(req.GetParamNum())
	binary.BigEndian.PutUint16(Magic, req.Magic)
	binary.BigEndian.PutUint16(CID, req.CID)
	binary.BigEndian.PutUint16(TypeParamNum, Type<<8+ParamNum)
	binary.BigEndian.PutUint16(Length, req.Length)
	binary.BigEndian.PutUint32(SrcAddr, req.SrcAddr)

	fields := [][]byte{
		Magic, CID, TypeParamNum, Length, SrcAddr, FuncName,
	}
	for _, v := range fields {
		res = append(res, v...)
	}

	// process Param-Part
	ParamPart := req.GetParamPart()
	for _, tlv := range ParamPart {
		res = append(res, tlv.ComposeTLV()...)
	}
	return
}

// ParseRequest is a function to transfer []byte to Request
func ParseRequest(msg []byte) (req *Request) {
	req = new(Request)

	Magic := binary.BigEndian.Uint16(msg[0:2])
	CID := binary.BigEndian.Uint16(msg[2:4])
	Type := uint8(msg[4])
	ParamNum := uint8(msg[5])
	Length := binary.BigEndian.Uint16(msg[6:8])
	SrcAddr := binary.BigEndian.Uint32(msg[8:12])
	FuncNameLen := Length - 4
	FuncName := string(msg[12 : 12+FuncNameLen])

	req.SetMagic(Magic)
	req.SetCID(CID)
	req.SetType(Type)
	req.SetParamNum(ParamNum)
	req.SetLength(Length)
	req.SetSrcAddr(SrcAddr)
	req.SetFuncName(FuncName)

	var tlv []TLV
	var i uint8 = 0
	tlvBeg := 12 + FuncNameLen
	for ; i < ParamNum; i++ {
		t := new(TLV)
		t.SetType(msg[tlvBeg])
		t.SetLength(msg[tlvBeg+1])
		t.SetValue(msg[tlvBeg+2 : tlvBeg+2+uint16(t.GetLength())])
		tlv = append(tlv, *t)
		tlvBeg += uint16(t.GetLength()) + 2
	}

	req.SetParamPart(tlv)
	return
}

// GetMagic is a get-method
func (req *Request) GetMagic() uint16 {
	if req != nil {
		return req.Magic
	}
	return 0
}

// SetMagic is a set-method
func (req *Request) SetMagic(Magic uint16) uint16 {
	if req != nil {
		req.Magic = Magic
		return req.Magic
	}
	return 0
}

// GetCID is a get-method
func (req *Request) GetCID() uint16 {
	if req != nil {
		return req.CID
	}
	return 0
}

// SetCID is a set-method
func (req *Request) SetCID(CID uint16) uint16 {
	if req != nil {
		req.CID = CID
		return req.CID
	}
	return 0
}

// GetType is a get-method
func (req *Request) GetType() uint8 {
	if req != nil {
		return req.Type
	}
	return 0
}

// SetType is a set-method
func (req *Request) SetType(Type uint8) uint8 {
	if req != nil {
		req.Type = Type
		return req.Type
	}
	return 0
}

// GetParamNum is a get-method
func (req *Request) GetParamNum() uint8 {
	if req != nil {
		return req.ParamNum
	}
	return 0
}

// SetParamNum is a set-method
func (req *Request) SetParamNum(ParamNum uint8) uint8 {
	if req != nil {
		req.ParamNum = ParamNum
		return req.ParamNum
	}
	return 0
}

// GetLength is a get-method
func (req *Request) GetLength() uint16 {
	if req != nil {
		return req.Length
	}
	return 0
}

// SetLength is a set-method
func (req *Request) SetLength(Length uint16) uint16 {
	if req != nil {
		req.Length = Length
		return req.Length
	}
	return 0
}

// GetSrcAddr is a get-method
func (req *Request) GetSrcAddr() uint32 {
	if req != nil {
		return req.SrcAddr
	}
	return 0
}

// SetSrcAddr is a set-method
func (req *Request) SetSrcAddr(SrcAddr uint32) uint32 {
	if req != nil {
		req.SrcAddr = SrcAddr
		return req.SrcAddr
	}
	return 0
}

// SetSrcAddrStr is a set-method
func (req *Request) SetSrcAddrStr(SrcAddr string) uint32 {
	if req != nil {
		ip := parseIPv4(SrcAddr)
		req.SrcAddr = ip
		return req.SrcAddr
	}
	return 0
}

// only for IPv4, not support IPv6
func parseIPv4(IP string) (ret uint32) {
	ret = 0
	ipData := strings.Split(IP, ".")
	for _, v := range ipData {
		p, _ := strconv.Atoi(v)
		ret = ret<<8 + uint32(p)
	}
	return
}

// GetFuncName is a get-method
func (req *Request) GetFuncName() string {
	if req != nil {
		return req.FuncName
	}
	return ""
}

// SetFuncName is a set-method
func (req *Request) SetFuncName(FuncName string) string {
	if req != nil {
		req.FuncName = FuncName
		return req.FuncName
	}
	return ""
}

// GetParamPart is a get-method
func (req *Request) GetParamPart() []TLV {
	if req != nil {
		return req.ParamPart
	}
	return nil
}

// SetParamPart is a set-method
func (req *Request) SetParamPart(ParamPart []TLV) []TLV {
	if req != nil {
		req.ParamPart = ParamPart
		return req.ParamPart
	}
	return nil
}
