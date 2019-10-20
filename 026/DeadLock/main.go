package main

import (
	"fmt"
	"sync"
)

func main() {
	lock := sync.Mutex{}

	fmt.Println("start lock")
	lock.Lock()
	lock.Lock()

	sync.RWMutex{}
}
