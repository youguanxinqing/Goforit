package main

import (
	"flag"
	"fmt"
	"os"
)

var cmdline = flag.NewFlagSet("", flag.ExitOnError)
var name = cmdline.String("name", "***", "your real name")

func loginprompt() {
	fmt.Println(*name + " login successfully.")
}

func main() {
	cmdline.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage Of %s: \n", "Question")
		cmdline.PrintDefaults()
	}
	cmdline.Parse(os.Args[1:])

	loginprompt()

}
