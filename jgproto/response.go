package jgproto

// Response is a struct
type Response struct {
	Magic      uint16
	CID        uint16
	Type       uint8
	ReturnNum  uint8
	ReturnPart []TLV
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
