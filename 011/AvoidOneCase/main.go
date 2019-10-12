package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)
	for _, i := range [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		c <- i
	}

	time.AfterFunc(time.Second * 1, func(){
		close(c)
	})

	for {
		select {
		case v, _ := <-c:
			if v == 3 {
				c = nil
				fmt.Printf("c = nil \n")
			}
			fmt.Printf("come into %v\n", v)
		default:
			fmt.Printf("come into default\n")
			goto end
		}
	}
end:
	fmt.Printf("over \n")
}