package main

import (
	"flag"
	stdlog "log"
	"os"
)

var log *stdlog.Logger

func main() {
	flag.Parse()
}

func init() {
	log = stdlog.New(os.Stderr, "", 0)

	flag.String("socket", "", "path to control socket")
	flag.String("state", "", "path to state directory")
}
