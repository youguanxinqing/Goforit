package main

import (
	"context"
	"fmt"
	"sync/atomic"
)

func main() {
	var num int32 = 0
	total := 10
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < total; i++ {
		go func() {
			atomic.AddInt32(&num, int32(1))
			fmt.Printf("num(%d)++\n", num)
			if atomic.LoadInt32(&num) == int32(total) {
				cancel() // 发起取消消息
			}
		}()
	}

	<-ctx.Done() // 接收
	fmt.Printf("task finish , and num is currently %d\n", num)
}
