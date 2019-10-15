package main

import "fmt"

func main() {
	count := 0
	chans := make(chan struct{}, 11)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			chans <- struct{}{}
		}()
	}

	for count < 10 {
		<-chans
		count++
	}
}
