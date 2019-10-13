package main

import "fmt"

type Cat struct {
	name string
}

func (c *Cat) setName(newName string) {
	c.name = newName
}

func main() {
	c := Cat{"ding"}
	fmt.Println(c.name)

	c.setName("zhong") // == (&c).setName(...)
	fmt.Println(c.name)
}
