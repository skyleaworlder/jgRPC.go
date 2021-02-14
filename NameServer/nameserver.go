package main

import "fmt"

// Name Server should implement several functions:
// 1. health probe
//	send HELLO(3) to Node Server,
// 	then receive HELLO(3) response from Node Server.
//
// 2. service discovery
//	receive NS(3) from load balance/cache Server.
//  send ip/port informations to load balance/cache Server.

func main() {
	fmt.Println("Hello World!")
}
