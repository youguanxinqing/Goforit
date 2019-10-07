package main

import (
	"flag"
	"fmt"
)

// MyStr 存放传进的参数与长度
type MyStr struct {
	v string // 存放值
	l int    // 存放长度
}

func (m *MyStr) String() string {
	return string(m.v)
}

// Set 赋值操作
func (m *MyStr) Set(value string) error {
	*m = MyStr{v: value, l: len(value)}
	return nil
}

// MyStrVar 解析参数
func MyStrVar(m *MyStr, name string, val string, usage string) {
	flag.CommandLine.Var(m, name, usage)
}

func main() {
	var name MyStr
	MyStrVar(&name, "name", "***", "your name")
	flag.Parse()

	fmt.Printf("%T\n", name)
	fmt.Println(name)
}
