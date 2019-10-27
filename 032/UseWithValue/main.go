package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx := context.WithValue(context.Background(), "z", "zhong")

	var once sync.Once
	once.Do(func() {
		ctx = context.WithValue(ctx, "g", "guan")
	})

	once.Do(func() {
		ctx = context.WithValue(ctx, "d", "ding")
	})

	v := ctx.Value("z")
	fmt.Println(v)

	fmt.Println(ctx.Value("g"))
	fmt.Println(ctx.Value("d")) // nil, 因为 Do 一生只会执行一次函数
}
