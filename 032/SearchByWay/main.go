package main

import (
	"context"
	"fmt"
)

func D(ctx context.Context) context.Context {
	Zctx := context.WithValue(ctx, "d", "ding")
	return Zctx
}

func G(ctx context.Context) context.Context {
	Gctx := context.WithValue(ctx, "g", "guan")
	return Gctx
}

func L(ctx context.Context) context.Context {
	Lctx := context.WithValue(ctx, "l", "liang")
	return Lctx
}

func main() {

	ctx := context.WithValue(context.Background(), "z", "zhong")
	D(ctx)
	gctx := G(ctx)
	lctx := L(gctx)

	fmt.Println(lctx.Value("k")) // nil
	//
	fmt.Println(lctx.Value("l"))
	fmt.Println(lctx.Value("g"))
	fmt.Println(lctx.Value("d")) // nil
	fmt.Println(lctx.Value("z"))
}
