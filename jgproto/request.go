package jgproto

import (
	"strconv"
	"strings"
)

// Request is a struct
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

// SetCID is a get-method
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

// SetType is a get-method
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

// SetParamNum is a get-method
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

// SetLength is a get-method
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

// SetSrcAddr is a get-method
func (req *Request) SetSrcAddr(SrcAddr uint32) uint32 {
	if req != nil {
		req.SrcAddr = SrcAddr
		return req.SrcAddr
	}
	return 0
}

// SetSrcAddrStr is a get-method
func (req *Request) SetSrcAddrStr(SrcAddr string) uint32 {
	if req != nil {
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

// GetParamPart is a get-method
func (req *Request) GetParamPart() []TLV {
	if req != nil {
		return req.ParamPart
	}
	return nil
}
