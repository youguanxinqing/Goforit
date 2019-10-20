package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Mutex{}

	m.Lock()
	count := 0
	for i := 0; i < 10; i++ {
		go func() {
			count++
		}()
	}
	m.Unlock()
	fmt.Println(count)
}
