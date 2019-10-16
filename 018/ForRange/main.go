package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4}  // 数组
	for i, e := range arr {
		if i == len(arr) - 1 {
			arr[0] += e
		} else {
			arr[i + 1] += e
		}
	}
	fmt.Println(arr) // [5 3 5 7]

	arr2 := []int{1, 2, 3, 4}  // 列表
	for i, e := range arr2 {
		if i == len(arr2) - 1 {
			arr2[0] += e
		} else {
			arr2[i + 1] += e
		}
	}
	fmt.Println(arr2) // [11 3 6 10]
}