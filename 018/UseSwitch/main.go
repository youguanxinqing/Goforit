package main

import "fmt"

func main() {
	arr := [...]int8{1, 2, 3}
	// switch 1 + 2 {  // 1 + 2 属于无类型常量
	// case arr[0]:
	// 	fmt.Println(arr[0])
	// case arr[1]:
	// 	fmt.Println(arr[1])
	// case arr[2]:
	// 	fmt.Println(arr[2])
	// }
	// 不能正常被编译

	// 正常运行
	switch arr[1] {
	case 1:
		fmt.Println(1)	
	case 2:
		fmt.Println(2)	
	case 3:
		fmt.Println(3)	
	}
}