package main

import "fmt"

func main() {
	x := [2][3]string{
		{"1", "2", "3"},
		{"1", "2", "3"},
	}

	fmt.Println(x)

	fMap := map[interface{}]string{
		[1][]string{{}}: "1", // 编译器允许，运行 panic
	}
	fmt.Println(fMap)
}
