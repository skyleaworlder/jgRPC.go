package conhash

import (
	"errors"
	"reflect"

	"github.com/skyleaworlder/jgRPC.go/LdBlsServer/ldbls"
)

// NodeComp is a less than comparison function
// this function is used in SortedMap.New()
func nodeComp(i, j interface{}) bool {
	return i.(*ldbls.Node).HID <= j.(*ldbls.Node).HID
}

// DeltaUint64 is a function return |i-j| on a ring mod 2^64
// due to the type of i and j is unsigned int64,
// so it's important to test whether one parameter is greater than the other one.
func deltaUint64(inpHID, nodeHID uint64) (uint64, error) {
	const (
		UINT64MAX     = ^uint64(0)
		HALFUINT64MAX = UINT64MAX >> 1
	)
	if reflect.TypeOf(inpHID).String() != "uint64" ||
		reflect.TypeOf(nodeHID).String() != "uint64" {
		msg := "Fatal Error: The type of i or j is not uint64\n"
		return UINT64MAX, errors.New(msg)
	}

	if inpHID < nodeHID {
		// 2.1 inpHID = 1, nodeHID = 5
		// 		delta = 5 - 1 = 4
		//
		// 2.2 inpHID = 1, nodeHID = 2^64 - 2
		// 		delta = 2^64 - 2 - 1 = 2^64 - 3 = 18446744073709551613
		return nodeHID - inpHID, nil
	}

	// 2.3. inpHID = 2^64 - 2, nodeHID = 1
	// 		delta = (2^64-1) - ((2^64-2) - 1) + 1 = 3
	// NOTICE: why "+1" ? because UINT64MAX = 2^64 - 1, instead of 2^64
	//
	// 2.4 inpHID = 5, nodeHID = 1
	//		delta = (2^64-1) - (5 - 1) + 1 = 2^64 - 4 = 18446744073709551612
	return UINT64MAX - (inpHID - nodeHID) + 1, nil
}
