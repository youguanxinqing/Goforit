package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var once sync.Once
	var num uint32 = 512

	once.Do(func() {
		// 方法 1
		// delta := int32(-3)
		// atomic.AddUint32(&num, uint32(delta))

		// 方法 2
		fmt.Println(2)
		fmt.Println(uint32(2))
		fmt.Println(^1)
		fmt.Println(1 ^ 1)
		fmt.Println(^2)
		fmt.Println(^3)
		atomic.AddUint32(&num, ^uint32(2))
	})

	fmt.Printf("num = %d\n", num)
}
