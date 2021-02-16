package consisthash

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/skyleaworlder/jgRPC.go/LdBlsServer/ldbls"
)

// InitConHash is a factory interface
// to initialize a conhash data structure
func InitConHash(cfg map[string]string) *Conhash {
	ht := new(Conhash)
	ht.Init()

	// fake initialze
	ha := sha1.New()
	hb := sha1.New()
	addra := "127.0.0.1:23331"
	addrb := "127.0.0.1:23332"
	io.WriteString(ha, addra)
	io.WriteString(hb, addrb)
	ares := binary.BigEndian.Uint64(ha.Sum(nil)[:8])
	bres := binary.BigEndian.Uint64(hb.Sum(nil)[:8])
	fmt.Printf("ares: %x\n", ha.Sum(nil))
	fmt.Printf("bres: %x\n", hb.Sum(nil))
	nodea := ldbls.Node{
		HID: ares, IP: net.ParseIP("127.0.0.1"), PORT: 23331,
	}
	nodeb := ldbls.Node{HID: bres, IP: net.ParseIP("127.0.0.1"), PORT: 23332}
	ht.PostNode(&nodea)
	ht.PostNode(&nodeb)

	// fake initialize end
	fmt.Println(ht)
	return ht
}
