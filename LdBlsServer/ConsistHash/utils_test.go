package consisthash

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
	fmt.Println(nodeComp(&anode, &bnode))
	fmt.Println(nodeComp(&bnode, &cnode))
	fmt.Println(nodeComp(&cnode, &anode))

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

func Test_deltaUint64(t *testing.T) {
	fmt.Println("2. deltaUint64:")

	var tmp uint64
	tmp, _ = deltaUint64(1, 5)
	fmt.Println("2.1 (1,5):", tmp)

	tmp, _ = deltaUint64(1, ^uint64(0)-1)
	fmt.Println("2.2 (1, 2^64-2):", tmp)

	tmp, _ = deltaUint64(^uint64(0)-1, 1)
	fmt.Println("2.3 (2^64-2,1):", tmp)

	tmp, _ = deltaUint64(5, 1)
	fmt.Println("2.4 (5,1):", tmp)

}
