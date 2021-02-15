package jgproto

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

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

// ConstructTLV is a constructor for TLV
func ConstructTLV(param interface{}) *TLV {
	tlv := new(TLV)

	type tlvInfo struct {
		Type   uint8
		Length uint8
	}

	m := map[string]tlvInfo{
		"int8": {0x00, 1}, "uint8": {0x01, 1},
		"int16": {0x02, 2}, "uint16": {0x03, 2},
		"int32": {0x04, 4}, "uint32": {0x05, 4},
		"int64": {0x06, 8}, "uint64": {0x07, 8},
		"bool": {0x08, 1}, "float32": {0x09, 4},
		"float64": {0x0a, 8}, "byte": {0x0b, 1},
		"string": {0x0c, ^uint8(0)},
	}
	typeStr := reflect.TypeOf(param).String()
	tlv.SetType(m[typeStr].Type)

	switch typeStr {
	case "int8", "uint8", "int16", "uint16",
		"int32", "uint32", "int64", "uint64",
		"bool", "float32", "float64", "byte":
		// for TLV.Length
		tlv.SetLength(m[typeStr].Length)

		// for TLV.Value
		buf := bytes.NewBuffer([]byte{})
		// binary.Write calls intDataSize function,
		// the switch-case statement in it doesn't process type string
		binary.Write(buf, binary.BigEndian, param)
		tlv.SetValue(buf.Bytes())
	case "string":
		tlv.SetLength(uint8(len(param.(string))))
		tlv.SetValue([]byte(param.(string)))
	}

	return tlv
}

// ComposeTLV is a method to generate a TLV
func (tlv *TLV) ComposeTLV() (res []byte) {
	TypeLength := make([]byte, 2)
	Type := uint16(tlv.GetType())
	Length := uint16(tlv.GetLength())
	Value := tlv.GetValue()

	binary.BigEndian.PutUint16(TypeLength, Type<<8+Length)
	res = append(res, TypeLength...)
	res = append(res, Value...)
	return
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
