package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage Of %s: \n", "Question")
		flag.PrintDefaults()
	}

	flag.Parse()
}
