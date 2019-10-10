package main

import "fmt"

func main() {
	var badMap = map[interface{}]int{
		"1":      1,
		[]int{2}: 2, // 编译器不会报错，语法上认同该做法，但运行报错
		3:        3,
	}

	fmt.Println(badMap)
}
