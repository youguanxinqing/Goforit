package main

import "fmt"

func main() {
	zty := 10
	_, ok := interface{}(zty).(int)
	if ok {
		fmt.Println("int 类型")
	} else {
		fmt.Println("非 int 类型")
	}

	guan := "zhong"
	_, ok = interface{}(guan).(int)
	if ok {
		fmt.Println("int 类型")
	} else {
		fmt.Println("非 int 类型")
	}
}
