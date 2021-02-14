package conhash

import (
	"github.com/skyleaworlder/jgRPC.go/LdBlsServer/ldbls"
	sm "github.com/umpc/go-sortedmap"
)

type conhash struct {
	HT *sm.SortedMap
}

// newConHash is a function to create an obj
func newConHash() *conhash {
	res := new(conhash)
	res.HT = sm.New(ldbls.M, HIDComp)
	return res
}
