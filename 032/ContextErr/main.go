package main

import (
	"context"
	"sync/atomic"
)

func main() {
	total := 10
	var sum int32 = 0
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < total; i++ {
		go func() {
			atomic.AddInt32(&sum, 1)
			if atomic.LoadInt32(&sum) == int32(32) {
				ctx.Err()
				cancel()
			}
		}()
	}
}
