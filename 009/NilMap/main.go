package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m) // map[]

	if m == nil {
		fmt.Println("m = nil") // 执行
	}

	if v, ok := m["zhong"]; ok { // 不存报错
		fmt.Println(v)
	} else {
		fmt.Println(v) // 0
	}

	delete(m, "zhong") // 删除某个键-元素对，正常运行

	m["10"] = 10 // 报错，不允许添加键-元素对
}
