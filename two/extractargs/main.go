package main

import (
	"flag"
	"fmt"
)

func hello(name string) {
	fmt.Println(name + " hello world!")
}

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "person object")
	// flag.String 与 flag.StringVar 功能类似，用法存在些许差异
	// name := flag.String("name", "everyone", "person object")
}

func main() {
	flag.Parse() // 真正解析命令参数，并将他们的值赋给对应变量
	hello(name)
}
