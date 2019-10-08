package main

import "fmt"

var container = []string{"0", "2", "2"}

func main() {
	container := map[int]string{0: "0", 1: "1", 2: "2"}
	fmt.Println(container[1])
}
