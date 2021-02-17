package consisthash

import (
	sm "github.com/umpc/go-sortedmap"
)

// InitConHash is a factory interface
// to initialize a conhash data structure
func InitConHash(iptable *sm.SortedMap, cfg map[string]string) *Conhash {
	ht := new(Conhash)
	ht.Init(iptable)

	// for debug
	// fmt.Println(ht)
	return ht
}
