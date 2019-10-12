package main

import "fmt"

func changeArrayValue(a [3]string) [3]string {
	a[1] = "zhong"
	return a
}

func main() {
	a := [3]string{
		"z", "t", "y",
	}
	fmt.Printf("a %v\n", a)

	f := changeArrayValue(a) // 不会改变数组 a 中的值
	fmt.Printf("f %v\n", f)
	fmt.Printf("a %v\n", a)
}
