package main

import "fmt"

// WriteData : send data to channel
func WriteData(c chan<- string, v string) {
	c <- v // 只能发
}

// ReadData : read data from channel
func ReadData(c <-chan string) {
	fmt.Printf("Read Data: %s\n", <-c) //  只能收
}

func main() {
	c := make(chan string, 20)
	for _, i := range [4]int{0, 1, 2, 3} {
		_ = i
		WriteData(c, "zty")
		ReadData(c)
	}
}
