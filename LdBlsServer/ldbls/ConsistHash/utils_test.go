package conhash

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net"
	"testing"

	"github.com/skyleaworlder/jgRPC.go/LdBlsServer/ldbls"
)

func Test_HIDComp(t *testing.T) {
	fmt.Println("1. HIDComp:")
	fmt.Println("1.1 simple test:")
	addra := "192.168.1.1:5000"
	addrb := "127.0.0.1:3000"
	addrc := "192.168.1.1:5001"
	anode := ldbls.Node{HID: 162, IP: net.ParseIP(addra), PORT: 5000}
	bnode := ldbls.Node{HID: 163, IP: net.ParseIP(addrb), PORT: 3000}
	cnode := ldbls.Node{HID: 164, IP: net.ParseIP(addrc), PORT: 5001}
	fmt.Println(NodeComp(&anode, &bnode))
	fmt.Println(NodeComp(&bnode, &cnode))
	fmt.Println(NodeComp(&cnode, &anode))

	fmt.Println("1.2 crypto/sha1 test:")
	ha, hb, hc := sha1.New(), sha1.New(), sha1.New()
	io.WriteString(ha, addra)
	io.WriteString(hb, addrb)
	io.WriteString(hc, addrc)
	ares, bres, cres := ha.Sum(nil), hb.Sum(nil), hc.Sum(nil)
	fmt.Printf("ha's res: %x\n", ares)
	fmt.Printf("hb's res: %x\n", bres)
	fmt.Printf("hc's res: %x\n", cres)
}
