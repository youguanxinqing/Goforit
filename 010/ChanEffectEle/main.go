package main

import "fmt"

func main() {
	// sliceChan := make(chan []int, 3)  // 切片的值会受到影响
	sliceChan := make(chan [3]int, 3) // 数组中的值不会收到影响

	// srcSlice := []int{1, 2, 3}
	srcSlice := [3]int{1, 2, 3}
	fmt.Printf("srcSlice %v\n", srcSlice)

	sliceChan <- srcSlice
	dstSlice := <-sliceChan
	dstSlice[1] = 512

	fmt.Printf("After handle ...\n")
	fmt.Printf("srcSlice %v and dstSlice %v\n", srcSlice, dstSlice) // srcSlice 中的值发生改变
}
