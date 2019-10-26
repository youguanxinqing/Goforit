package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var lock sync.RWMutex
	rlock := sync.NewCond(lock.RLocker())

	var wg sync.WaitGroup

	go func(wg *sync.WaitGroup) {
		lock.RLock()
		rlock.Wait()
		fmt.Println("add")
		wg.Add(1)
		lock.RUnlock()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		lock.RLock()
		rlock.Wait()
		fmt.Println("wait")
		wg.Wait()
		lock.RUnlock()
	}(&wg)

	time.Sleep(time.Second * 5)
	rlock.Broadcast()
	time.Sleep(time.Second * 5)
}
