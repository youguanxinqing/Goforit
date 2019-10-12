package main

import (
	"errors"
	"fmt"
)

// 定义一个类型
type operate func(int, int) int

func oper(x, y int) int {
	return x + y
}

func calculate(x, y int, op operate) (int, error) {
	// 卫述语句检查参数
	if op == nil {
		return 0, errors.New("op is nil")
	}
	return op(x, y), nil
}

func main() {
	r, _ := calculate(1, 2, oper)
	fmt.Println(r)
}
