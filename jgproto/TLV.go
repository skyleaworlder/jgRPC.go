package jgproto

// TLV is a struct
//                                  1  1  1  1  1  1
//    0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
// 	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
// 	|          TYPE         |         LENGTH        |
// 	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
// 	|                                               |
// 	/                     VALUE                     /
// 	|                                               |
// 	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
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
