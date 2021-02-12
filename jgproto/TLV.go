package jgproto

// TLV is a struct
type TLV struct {
	Type   uint8
	Length uint8
	Value  []byte
}

// GetType is a get-method
func (tlv *TLV) GetType() uint8 {
	if tlv != nil {
		return tlv.Type
	}
	return 0
}

// SetType is a set-method
func (tlv *TLV) SetType(Type uint8) uint8 {
	if tlv != nil {
		tlv.Type = Type
		return tlv.Type
	}
	return 0
}

// GetLength is a get-method
func (tlv *TLV) GetLength() uint8 {
	if tlv != nil {
		return tlv.Length
	}
	return 0
}

// SetLength is a set-method
func (tlv *TLV) SetLength(Length uint8) uint8 {
	if tlv != nil {
		tlv.Length = Length
		return tlv.Length
	}
	return 0
}

// GetValue is a get-method
func (tlv *TLV) GetValue() []byte {
	if tlv != nil {
		return tlv.Value
	}
	return []byte{}
}

// SetValue is a set-method
func (tlv *TLV) SetValue(Value []byte) []byte {
	if tlv != nil {
		tlv.Value = Value
		return tlv.Value
	}
	return []byte{}
}
