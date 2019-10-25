package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var num int32 = 512
	// res := atomic.CompareAndSwapInt32(&num, 10, 0)
	res := atomic.CompareAndSwapInt32(&num, 512, 0)
	fmt.Println(res)
}
