package main

import "fmt"

type Pet interface {
	Walk()
}

type Dog struct {
	name string
}

func (d *Dog) Walk() {
	fmt.Println("dog walk ...")
}

func main() {
	dog := Dog{"little dog"}
	// p 是接口类型变量，&dog 是 p 的动态值，Dog 是 p 的动态类型
	var p Pet = &dog
	p.Walk()

	var p1 Pet
	fmt.Println(p1)        // <nil>
	fmt.Println(p1 == nil) // true

	var d2 *Dog
	var p2 Pet = d2
	fmt.Println(p2)        // <nil>
	fmt.Println(p2 == nil) // false
}
