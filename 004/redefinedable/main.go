package main

import "fmt"

func two_string() (string, string) {
	return "zhong", "ting"
}

func main() {

	z, t := two_string()
	t, n := two_string()
	fmt.Println(z, t, n)

	var y string
	g, y := two_string()
	fmt.Println(g, y)

	// 结论:
	//     都可以实现变量重声明

}
