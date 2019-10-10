package main

import (
	"fmt"
	"container/ring"
)

func main() {
	r := ring.New(5)
	fmt.Println(r.Len())

	p := r
	fmt.Println(p.Value)
	for p = p.Next(); p != r; p = p.Next() {
		fmt.Println(p.Value)
	}
}