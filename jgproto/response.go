package jgproto

import "encoding/binary"

// Response is a struct
//                                  1  1  1  1  1  1
//    0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//  |           J           |           G           |
//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//  |                      CID                      |
//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//  |          TYPE         |       RETURN_NUM      |
//  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//  |                   RETURN_PART                 |
type Response struct {
	Magic      uint16
	CID        uint16
	Type       uint8
	ReturnNum  uint8
	ReturnPart []TLV
}

// ConstructResponse is a default constructor method
//
// Magic: JG
//
// CID: the same as request
//
// Type: default OK(1)
//
// ReturnNum: default 0, void
//
// ReturnPart: default []TLV{}
func ConstructResponse(CID uint16) *Response {
	resp := new(Response)
	resp.SetMagic(0x4A47)
	resp.SetCID(CID)
	resp.SetType(1)
	resp.SetReturnNum(0)
	resp.SetReturnPart([]TLV{})
	return resp
}

// ComposeResponse is a constructor method
func (resp *Response) ComposeResponse() (res []byte) {

	Magic := make([]byte, 2)
	CID := make([]byte, 2)
	TypeReturnNum := make([]byte, 2)

	Type := uint16(resp.GetType())
	ReturnNum := uint16(resp.GetReturnNum())
	binary.BigEndian.PutUint16(Magic, resp.Magic)
	binary.BigEndian.PutUint16(CID, resp.CID)
	binary.BigEndian.PutUint16(TypeReturnNum, Type<<8+ReturnNum)

	fields := [][]byte{
		Magic, CID, TypeReturnNum,
	}
	for _, v := range fields {
		res = append(res, v...)
	}

	ReturnPart := resp.GetReturnPart()
	for _, tlv := range ReturnPart {
		res = append(res, tlv.ComposeTLV()...)
	}
	return
}

// ParseResponse is a function to process []byte into Response
func ParseResponse(msg []byte) (Type uint8, ReturnNum uint8, res []interface{}) {
	resp := new(Response)

	Magic := binary.BigEndian.Uint16(msg[0:2])
	CID := binary.BigEndian.Uint16(msg[2:4])
	Type, ReturnNum = uint8(msg[4]), uint8(msg[5])

	resp.SetMagic(Magic)
	resp.SetCID(CID)
	resp.SetType(Type)
	resp.SetReturnNum(ReturnNum)

	// process TLV
	var i uint8 = 0
	var tlvBeg uint8 = 6
	for ; i < ReturnNum; i++ {
		// get the length of value
		Length := uint8(msg[tlvBeg+1])
		// tlvBeg points to TYPE of TLV
		// e.g. tlvBeg = 6, Length = 10, Type = string
		//		msg[tlvBeg : tlvBeg+Length+2] = msg[6:18]
		_, _, val := ParseTLV(msg[tlvBeg : tlvBeg+Length+2])
		res = append(res, val)
		// then tlvBeg points to next TLV
		tlvBeg += Length + 2
	}
	return
}

// GetMagic is a get-method
func (resp *Response) GetMagic() uint16 {
	if resp != nil {
		return resp.Magic
	}
	return 0
}

// SetMagic is a set-method
func (resp *Response) SetMagic(Magic uint16) uint16 {
	if resp != nil {
		resp.Magic = Magic
		return resp.Magic
	}
	return 0
}

// GetCID is a get-method
func (resp *Response) GetCID() uint16 {
	if resp != nil {
		return resp.CID
	}
	return 0
}

// SetCID is a set-method
func (resp *Response) SetCID(CID uint16) uint16 {
	if resp != nil {
		resp.CID = CID
		return resp.CID
	}
	return 0
}

// GetType is a get-method
func (resp *Response) GetType() uint8 {
	if resp != nil {
		return resp.Type
	}
	return 0
}

// SetType is a set-method
func (resp *Response) SetType(Type uint8) uint8 {
	if resp != nil {
		resp.Type = Type
		return resp.Type
	}
	return 0
}

// GetReturnNum is a get-method
func (resp *Response) GetReturnNum() uint8 {
	if resp != nil {
		return resp.ReturnNum
	}
	return 0
}

// SetReturnNum is a set-method
func (resp *Response) SetReturnNum(ReturnNum uint8) uint8 {
	if resp != nil {
		resp.ReturnNum = ReturnNum
		return resp.ReturnNum
	}
	return 0
}

// GetReturnPart is a get-method
func (resp *Response) GetReturnPart() []TLV {
	if resp != nil {
		return resp.ReturnPart
	}
	return nil
}

// SetReturnPart is a set-method
func (resp *Response) SetReturnPart(ReturnPart []TLV) []TLV {
	if resp != nil {
		resp.ReturnPart = ReturnPart
		return resp.ReturnPart
	}
	return nil
}
