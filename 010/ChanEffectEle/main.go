package main

import "fmt"

func main() {
	sliceChan := make(chan []int, 3)

	srcSlice := []int{1, 2, 3}
	fmt.Printf("srcSlice %v\n", srcSlice)
	
	sliceChan <- srcSlice
	dstSlice := <-sliceChan
	dstSlice[1] = 512

	fmt.Printf("After handle ...\n")
	fmt.Printf("srcSlice %v and dstSlice %v\n", srcSlice, dstSlice)  // srcSlice 中的值发生改变
}