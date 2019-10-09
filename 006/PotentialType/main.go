package main

import "fmt"

func main() {
	type string1 string
	type string2 string

	var name string1 = "zhong"
	fmt.Println(string2(name))

	fmt.Printf("%T\n", name)  // main.string1

	var alphas = []string1{"a", "b", "c"}
	fmt.Println(alphas)
	fmt.Println([]string(alphas))  // 不允许操作
}