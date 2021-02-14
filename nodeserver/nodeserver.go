package main

import "fmt"

// NodeServer should implement several functions:
// 1. service register
//	send REGISTER(1) to Name Server
//
// 2. health probe
//	receive HELLO(3) from Name Server,
// 	then send HELLO(3) response to Name Server.
//
// 3. communicate with load balance/cache Server
// 	receive CALL(0) request, then send execute results.

func main() {
	fmt.Println("Hello World!")
}
