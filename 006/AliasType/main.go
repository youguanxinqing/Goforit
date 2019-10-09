package main

import "fmt"

type MyString1 = string
type MyString2 = string

func outStr(name string) {
	fmt.Println(name)
}

func main() {
	var name string = "zhong"
	outStr(name)

	var nameMySr MyString1 = "zhong"
	outStr(nameMySr)  // 可以这样使用，但类型再定义不能

	var guan MyString1 = "zhong"
	fmt.Println(MyString2(guan))
	fmt.Printf("%T\n", guan)  // string

	var names = []MyString1{"a", "b", "c"}
	fmt.Println(names)
	fmt.Println([]string(names))
}