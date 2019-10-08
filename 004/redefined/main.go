package main

import "fmt"

func two_string() (string, string) {
	return "zhong", "ting"
}

func main() {
	var name string
	n, name := two_string()
	fmt.Println(n, name)
}
