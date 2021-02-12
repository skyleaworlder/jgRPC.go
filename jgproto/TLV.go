package jgproto

// TLV is a struct
type TLV struct {
	Type   uint8
	Length uint8
	Value  []byte
}
