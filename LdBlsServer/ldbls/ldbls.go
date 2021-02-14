package ldbls

import "net"

// Node is a struct
type Node struct {
	HID  []byte
	ip   net.IP
	port uint16
}

const (
	// M is maximum num of Nodes
	M = 16
)

// Alg is an interface
type Alg interface {
	// Init is to initialize Data Structure in Algothrim
	Init() error

	// 4 "REST-ful" methods
	// GetNode is to fetch one Node suitable for situation.
	// In GetNode method, msg is HID given.
	GetNode(HID []byte) (*Node, error)
	// PostNode is to add a new Node into Data Structure in Algothrim
	// In PostNode method, msg is a struct including all data needed.
	PostNode(n *Node) (*Node, error)
	// PutNode is to update an existed Node.
	// In PutNode method, msg is a struct including all data needed.
	PutNode(n *Node) (*Node, error)
	// DeleteNode is to delete an existed Node.
	// In DeleteNode method, msg is HID.
	DeleteNode(HID []byte) (*Node, error)
}
