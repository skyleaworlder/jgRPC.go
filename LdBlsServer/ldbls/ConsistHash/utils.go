package conhash

import (
	"github.com/skyleaworlder/jgRPC.go/LdBlsServer/ldbls"
)

// NodeComp is a less than comparison function
// this function is used in SortedMap.New()
func NodeComp(i, j interface{}) bool {
	return i.(*ldbls.Node).HID <= j.(*ldbls.Node).HID
}
