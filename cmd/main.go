package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/williamchanrico/xtest/cmd/xtest"
	"github.com/williamchanrico/xtest/log"
)

var (
	version            string
	showVersionAndExit bool
)

func main() {
	f := xtest.Flags{}
	flag.StringVar(&f.HTTPAddress, "http-address", "0.0.0.0:9000", "HTTP API Listener")
	flag.StringVar(&f.LogLevel, "log-level", "info", "App log level")
	flag.Parse()

	if showVersionAndExit {
		fmt.Printf("xtest %v\n", version)
		os.Exit(0)
	}

	log.SetLevelString(f.LogLevel)

	exitCode, err := xtest.Run(f)
	if err != nil {
		log.Error(err)
	}

	if exitCode != 0 {
		log.Errorf("xtest exited with code %d", exitCode)
	}
	os.Exit(exitCode)
}