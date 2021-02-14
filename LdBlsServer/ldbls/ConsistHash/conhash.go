package conhash

import (
	"errors"
	"fmt"
	"os"

	"github.com/skyleaworlder/jgRPC.go/LdBlsServer/ldbls"
	sm "github.com/umpc/go-sortedmap"
)

type conhash struct {
	HT *sm.SortedMap
}

// newConHash is a function to create an obj
func (c *conhash) Init() error {
	if c == nil {
		msg := "Fatal Error: conhash.Init try to init the space that a nil pointer points\n"
		fmt.Fprint(os.Stderr, msg)
		return errors.New(msg)
	}
	c.HT = sm.New(ldbls.M, nodeComp)
	if c.HT == nil {
		msg := "Fatal Error: conhash.Init fails to uses SortedMap as data-structure to init hash table\n"
		fmt.Fprint(os.Stderr, msg)
		return errors.New(msg)
	}
	return nil
}

// GetNode is a method that can select the most suitable node in hash ring.
// In Consist Hash, all nodes are put into a ring, identified by its KEY(HID).
// I define uint64 as the type of HID, so 2^64 nodes are max-supported.
// This method return the nearest node in a clockwise direction.
//
// e.g. there is 7 nodes in ring:
// 	o---o---o---o---o---o---o---...
// 	0   1   2   3   4   5   6
//
// And their key is:
//	"0": 0x10, "1": 0x23, "2": 0x45, "3": 0x78, "4": 0xaa, "5": 0xbd, "6": 0xfa
//
// And we suppose that M's maximum is 256, so the max length of keys is 8.
// (since ceiling(log_2(256)) = 8)
//
// 1. given HID is 0x56, then return "3";
// 2. given HID is 0x23, then return "1";
// 3. given HID is 0xfb, then return "0".
func (c *conhash) GetNode(HID uint64) (*ldbls.Node, error) {
	iter, err := c.HT.IterCh()
	if err != nil {
		msg := "Fatal Error: conhash.GetNode cannot generate an iter\n"
		fmt.Fprint(os.Stderr, err.Error()+msg)
		return &ldbls.Node{}, errors.New(err.Error() + msg)
	}
	defer iter.Close()

	// Try to find the nearest node on ring.
	var delta uint64 = ^uint64(0)
	var target *ldbls.Node
	for rec := range iter.Records() {
		tmpDelta, err := deltaUint64(rec.Key.(uint64), HID)
		if err != nil {
			msg := "Fatal Error: conhash.DeltaUint64 failed\n"
			fmt.Fprint(os.Stderr, err.Error()+msg)
			return &ldbls.Node{}, errors.New(err.Error() + msg)
		}

		// refresh
		if tmpDelta < delta {
			// for debug
			// fmt.Print("delta:", tmpDelta, "\n")
			delta = tmpDelta
			target = rec.Val.(*ldbls.Node)
		}
	}

	return target, nil
}

// PostNode is a method that insert a node into consist hash ring.
// because there is a restriction about hash ring, namely, M(ring's size),
// if insert a node when hash ring is full, an error will occur.
//
// But it has to be NOTICED that
// hash ring has an limit size theoretically(2^64 for uint64 type is used in HID),
// while M is another limit on ring. M must be smaller than 2^64.
// M even should be smaller than 2^63, from my perspective in theory.
// And it's recommanded that M should be smaller than 2^10 in reality.
//
// other HT solution might be taken into consideration in the future perhaps?
// e.g. chord-DHT?
func (c *conhash) PostNode(n *ldbls.Node) (*ldbls.Node, error) {
	ok := c.HT.Insert(n.HID, n)
	if !ok {
		msg := "Just Warning: conhash.PostNode failed, for HT is full perhaps\n"
		fmt.Fprint(os.Stderr, msg)
		return &ldbls.Node{}, errors.New(msg)
	}
	return n, nil
}

// DeleteNode is a method used to delete one node from hash ring.
// Do not delete an unexisted node from ring, which will generate an error.
func (c *conhash) DeleteNode(HID uint64) (*ldbls.Node, error) {
	node, ok := c.HT.Get(HID)
	if !ok {
		msg := "Fatal Error: conhash.DeleteNode cannot get node from HT\n"
		fmt.Fprint(os.Stderr, msg)
		return node.(*ldbls.Node), errors.New(msg)
	}
	// try to delete node using HID
	c.HT.Delete(HID)
	return node.(*ldbls.Node), nil
}

// PutNode method's 3 phases:
// 1. check whether node given is in HT.
// 2. delete the node in HT that own the same key as given node.
// 3. insert a new node which is with the same key as the deleted node,
//    but with different properties, like addr, port, etc.
func (c *conhash) PutNode(n *ldbls.Node) (*ldbls.Node, error) {
	// since sm.Replace don't have return val, I don't use c.HT.Replace(n.HID, n)
	if !c.HT.Has(n.HID) {
		msg := "Fatal Error: conhash.PutNode cannot use sm.Has function to find node from HT\n"
		fmt.Fprint(os.Stderr, msg)
		return &ldbls.Node{}, errors.New(msg)
	}

	_, err := c.DeleteNode(n.HID)
	if err != nil {
		msg := "Fatal Error: conhash.PutNode cannot use DeleteNode to delete node from HT\n"
		fmt.Fprint(os.Stderr, err.Error()+msg)
		return &ldbls.Node{}, errors.New(err.Error() + msg)
	}

	node, err := c.PostNode(n)
	if err != nil {
		msg := "Fatal Error: conhash.PutNode cannot use PostNode to insert node into HT\n"
		fmt.Fprint(os.Stderr, err.Error()+msg)
		return node, errors.New(err.Error() + msg)
	}

	return node, nil
}
