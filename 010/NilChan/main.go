package main

import "fmt"

func main() {
	var c chan int
	c <- 1 // 编译器不会报错，但运行会出现 error: fatal error: all goroutines are asleep - deadlock!
	fmt.Println("programer over!")
}
