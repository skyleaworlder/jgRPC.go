package ldbls

import "net"

// Node is a struct
type Node struct {
	HID  uint64
	IP   net.IP
	PORT uint16
}

const (
	// M is maximum num of Nodes
	// 1. Solution ConsistHash:
	// SortedMap uses built-in map to store data, and any slice type is unhashable,
	// so []byte(generated by crypto/sha1) are forbidden to use as KEY.
	// In ConsistHash, I use binary.BigEndian.Uint64 to process sha1 result.
	// Thus, M need to be restricted in internal: (0, 2^64).
	M = 16
)

// Alg is an interface
type Alg interface {
	// Init is to initialize Data Structure in Algothrim
	Init() error

	// 4 "REST-ful" methods
	// GetNode is to fetch one Node suitable for situation.
	// In GetNode method, msg is HID given.
	GetNode(HID uint64) (*Node, error)
	// PostNode is to add a new Node into Data Structure in Algothrim
	// In PostNode method, msg is a struct including all data needed.
	PostNode(n *Node) (*Node, error)
	// PutNode is to update an existed Node.
	// In PutNode method, msg is a struct including all data needed.
	PutNode(n *Node) (*Node, error)
	// DeleteNode is to delete an existed Node.
	// In DeleteNode method, msg is HID.
	DeleteNode(HID uint64) (*Node, error)
}
