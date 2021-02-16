package jgrpcutils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Readcfg is a function to initialize configuration
func Readcfg(cfg map[string]string, cfgAddr string) {
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
		cfg[key] = val
	}
}
