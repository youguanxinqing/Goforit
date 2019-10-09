package main

import "fmt"

func main() {
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	// 实现切片右向扩展
	s5 := s4[0:cap(s4)]

	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
}
