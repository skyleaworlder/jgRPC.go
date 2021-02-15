package jgproto

import (
	"fmt"
	"testing"
)

func Test_ParseResponse(t *testing.T) {
	fmt.Println("5.1: jgproto.ParseResponse test:")
	resp := ConstructResponse(12345)

	resp.SetReturnNum(13)

	var a int8 = 0
	var b uint8 = 1
	var c int16 = 2
	var d uint16 = 3
	var e int32 = 4
	var f uint32 = 5
	var g int64 = 6
	var h uint64 = 7
	var i bool = true
	var j float32 = 9.1
	var k float64 = 10.2
	var l byte = 0x6e
	var m string = "asdadasdwqewq"
	testData := []interface{}{
		a, b, c, d, e, f, g, h, i, j, k, l, m,
	}

	sls := []TLV{}
	for _, val := range testData {
		sls = append(sls, *ConstructTLV(val))
	}
	resp.SetReturnPart(sls)

	buf := resp.ComposeResponse()
	fmt.Println(ParseResponse(buf))
}
