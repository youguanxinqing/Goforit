package main

import (
	"fmt"
	"math/rand"
)

func main() {
	intChans := [3]chan int {
		make(chan int, 3),
		make(chan int, 3),
		make(chan int, 3),
	}

	randNum := rand.Intn(3)
	fmt.Printf("randNum : %v\n", randNum)
	intChans[randNum] <- randNum 

	select {
	case <-intChans[0]:
		fmt.Printf("index = 0\n")
	case <-intChans[1]:
		fmt.Printf("index = 1\n")
	case elem := <-intChans[2]:
		fmt.Printf("index = 2 %v\n", elem)
	default:
		fmt.Printf("no elem\n")
	}
}