package main

import (
	"flag"
	"fmt"
)

func hello(name string) {
	fmt.Println(name + " hello world!")
}

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "person object")
}

func main() {

	flag.Usage = func() {
		fmt.Println("Usage Of Greet:")
		flag.PrintDefaults()
	}

	flag.Parse()
	hello(name)
}
