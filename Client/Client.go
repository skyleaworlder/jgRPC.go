package main

import (
	"fmt"

	jgrpcc "github.com/skyleaworlder/jgRPC.go/Client/jgrpc"
	jgut "github.com/skyleaworlder/jgRPC.go/jgrpcUtils"
)

// Client should be able to implement:
// 1. interface of RPC
//	send CALL(0), receive OK(1) response

const (
	cfgMaxNum = 16
)

var (
	// Config is a map
	// e.g. "Local_IP" => "127.0.0.1"
	Config = make(map[string]string, cfgMaxNum)
)

func main() {
	jgut.Readcfg(Config, "client.cfg")
	calcu := new(jgrpcc.Calculator)
	calcu.Config = Config

	res := calcu.Add(1, 2)
	fmt.Println(res)
}
