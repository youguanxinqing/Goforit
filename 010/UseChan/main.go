package main

import "fmt"

func main() {
	intChan := make(chan int, 8)
	
	intChan <- 1
	intChan <- 2
	intChan <- 3

	fmt.Println(<-intChan)

}