package conhash

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"testing"

	"github.com/skyleaworlder/jgRPC.go/LdBlsServer/ldbls"
)

func Test_RESTfulAPI(t *testing.T) {
	fmt.Println("1. conhash.REST-API test:")
	ha, hb, hc := sha1.New(), sha1.New(), sha1.New()
	addra := "192.168.1.1:5000"
	addrb := "127.0.0.1:3000"
	addrc := "192.168.1.1:5001"
	io.WriteString(ha, addra)
	io.WriteString(hb, addrb)
	io.WriteString(hc, addrc)
	ares := binary.BigEndian.Uint64(ha.Sum(nil)[:8])
	bres := binary.BigEndian.Uint64(hb.Sum(nil)[:8])
	cres := binary.BigEndian.Uint64(hc.Sum(nil)[:8])
	fmt.Printf("ares: %x\n", ha.Sum(nil))
	fmt.Printf("bres: %x\n", hb.Sum(nil))
	fmt.Printf("cres: %x\n", hc.Sum(nil))

	ht := newConHash()
	nodea := ldbls.Node{
		HID: ares, IP: net.ParseIP("192.168.1.1"), PORT: 5000,
	}
	nodeb := ldbls.Node{
		HID: bres, IP: net.ParseIP("127.0.0.1"), PORT: 3000,
	}
	nodec := ldbls.Node{
		HID: cres, IP: net.ParseIP("192.168.1.1"), PORT: 5001,
	}
	ht.PostNode(&nodea)
	ht.PostNode(&nodeb)
	ht.PostNode(&nodec)

	fmt.Println("1.1 I try to search node a:")
	na, ok := ht.HT.Get(ares)
	if !ok {
		fmt.Println("fuck! fail!")
	} else {
		fmt.Println("wuhu! success find it!")
		fmt.Println(na)
	}

	fmt.Println("1.2 I try to delete node a:")
	ht.DeleteNode(nodea.HID)
	na, ok = ht.HT.Get(ares)
	if !ok {
		fmt.Println("wuhu! not found! success")
	} else {
		fmt.Println("fuck! fail!")
		fmt.Println(na)
	}

	fmt.Println("1.3 I try to modify node a:")
	ht.PutNode(&nodea)
	na, ok = ht.HT.Get(ares)
	if !ok {
		fmt.Println("wuhu! not found! success")
	} else {
		fmt.Println("fuck! fail!")
		fmt.Println(na)
	}

	fmt.Println("1.4 I try to modify node b:")
	ht.PutNode(&ldbls.Node{HID: nodeb.HID, IP: nodeb.IP, PORT: 3306})
	nb, ok := ht.HT.Get(bres)
	if !ok {
		fmt.Println("fuck! fail!")
	} else {
		fmt.Println("wuhu! I find it!")
		fmt.Println(nb)
	}

	fmt.Println("1.5 I try to search node b:")
	nb, _ = ht.GetNode(bres + 10000)
	if !ok {
		fmt.Println("fuck! fail!")
	} else {
		fmt.Println("wuhu! I find it!")
		fmt.Println(nb)
	}
}
