package main

import "fmt"

func main() {
	// num := 1
	// switch num {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 2, 3:
	// 	fmt.Println("2 or 3")
	// }

	num := 2
	arr := [...]int{1, 2, 3}
	switch num {
	case arr[0], arr[1]:
		fmt.Println("index 0 or 1")
	case arr[1], arr[2]:
		fmt.Println("index 1 or 2")
	}

	// 类型判断要求必须转 interface{} 类型
	switch t := interface{}(num).(type) {
	case byte:
		fmt.Println(1)
	case uint8:
		fmt.Println(2)
	case float32:
		fmt.Println("float32")
	case int:
		fmt.Println("int")
	default:
		fmt.Printf("do not know %T\n", t)
	}
}