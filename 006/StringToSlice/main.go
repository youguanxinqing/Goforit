package main

import "fmt"

func main() {
	var name = "钟庭英"

	fmt.Println([]rune(name))
	
	// fmt.Println([]string(name))  // 报错

	fmt.Println(string([]byte{'\xe4', '\xbd', '\xa0', '\xe5', '\xa5', '\xbd'}))  // 你好

	fmt.Println(string([]rune{'\u4F60', '\u597D'})) // 你好
}