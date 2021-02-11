package jgproto

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

// Response is a struct
type Response struct {
	Magic      uint16
	CID        uint16
	Type       uint8
	ReturnNum  uint8
	ReturnPart []TLV
}

// TLV is a struct
type TLV struct {
	Type   uint8
	Length uint8
	Value  []byte
}
