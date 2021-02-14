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
func newConHash() *conhash {
	res := new(conhash)
	res.HT = sm.New(ldbls.M, NodeComp)
	return res
}

// GetNode is a method that can select the most suitable node in hash ring.
// In Consist Hash, all nodes are put into a ring, identified by its KEY(HID).
// I define uint64 as the type of HID, so 2^64 nodes are max-supported.
// This method return the nearest node.
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
		tmpDelta, err := DeltaUint64(rec.Key.(uint64), HID)
		if err != nil {
			msg := "Fatal Error: conhash.DeltaUint64 failed\n"
			fmt.Fprint(os.Stderr, err.Error()+msg)
			return &ldbls.Node{}, errors.New(err.Error() + msg)
		}

		// refresh
		if tmpDelta < delta {
			fmt.Print("delta:", tmpDelta, "\n")
			delta = tmpDelta
			target = rec.Val.(*ldbls.Node)
		}
	}

	return target, nil
}

func (c *conhash) PostNode(n *ldbls.Node) (*ldbls.Node, error) {
	c.HT.Insert(n.HID, n)
	return n, nil
}

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
