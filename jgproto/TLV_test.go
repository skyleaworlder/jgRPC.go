package jgproto

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func Test_ConstructTLV(t *testing.T) {
	fmt.Println("4.1 TLV constructor test:")

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

	for idx, rec := range testData {
		fmt.Println("4.1." + strconv.Itoa(idx) + ": " + reflect.TypeOf(rec).String())
		tlv := ConstructTLV(rec)
		fmt.Printf("type:  %x\n", tlv.GetType())
		fmt.Printf("len:   %d\n", tlv.GetLength())
		fmt.Printf("value: %x\n", tlv.GetValue())
		fmt.Println("")
	}
}

func Test_ComposeTLV(t *testing.T) {
	fmt.Println("4.2 TLV Composer test:")

	var a float64 = 10.526
	tlv := ConstructTLV(a)
	tlvByte := tlv.ComposeTLV()

	fmt.Println(tlvByte)
}
