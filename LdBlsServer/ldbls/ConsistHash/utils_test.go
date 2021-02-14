package conhash

import (
	"crypto/sha1"
	"fmt"
	"io"
	"testing"
)

func Test_HIDComp(t *testing.T) {
	fmt.Println("1. HIDComp:")
	fmt.Println("1.1 simple test:")
	fmt.Println(HIDComp([]byte{0x01}, []byte{0x02}))
	fmt.Println(HIDComp([]byte{0x02}, []byte{0x02}))
	fmt.Println(HIDComp([]byte{0x03}, []byte{0x02}))

	fmt.Println("1.2 crypto/sha1 test:")
	ha, hb, hc := sha1.New(), sha1.New(), sha1.New()
	addra := "192.168.1.1:5000"
	addrb := "127.0.0.1:3000"
	addrc := "192.168.1.1:5001"
	io.WriteString(ha, addra)
	io.WriteString(hb, addrb)
	io.WriteString(hc, addrc)
	ares, bres, cres := ha.Sum(nil), hb.Sum(nil), hc.Sum(nil)
	fmt.Printf("ha's res: %x\n", ares)
	fmt.Printf("hb's res: %x\n", bres)
	fmt.Printf("hc's res: %x\n", cres)
	fmt.Println("ha's res and hb's res:", HIDComp(ares, bres))
	fmt.Println("hb's res and hc's res:", HIDComp(bres, cres))
	fmt.Println("hc's res and ha's res:", HIDComp(cres, ares))
}
