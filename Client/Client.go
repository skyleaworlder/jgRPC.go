package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	readcfg("client.cfg")
}

func readcfg(cfgAddr string) {
	fd, err := os.Open(cfgAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer fd.Close()

	scan := bufio.NewScanner(fd)
	for scan.Scan() {
		lineTxt := scan.Text()
		// comments or blank line
		// short-circuit operation
		if len(lineTxt) == 0 || lineTxt[0] == '#' {
			continue
		}

		// process line text of config file
		cfgSls := strings.Split(lineTxt, "=")
		key, val := cfgSls[0], cfgSls[1]
		Config[key] = val
	}
}
