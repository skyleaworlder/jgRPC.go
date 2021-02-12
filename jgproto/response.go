package jgproto

// Response is a struct
type Response struct {
	Magic      uint16
	CID        uint16
	Type       uint8
	ReturnNum  uint8
	ReturnPart []TLV
}
