package main

import "fmt"

type Pet interface {
	Walk()
}

type Dog struct {
}

func (d *Dog) Walk() {
	fmt.Println("dog walk...")
}

func (d *Dog) Fly() {
	fmt.Println("dog Fly...")
}

func main() {
	var d Dog
	var p Pet = &d
	p.Walk()
}
