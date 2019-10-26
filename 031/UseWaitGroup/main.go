package main

import (
	"fmt"
	"sync"
)

func main() {
	flagCount := 0
	// 注意：这里不能是空指针
	var wg sync.WaitGroup

	var lock sync.Mutex

	wg.Add(10)
	for flagCount < 10 {
		go func(wg *sync.WaitGroup, lock *sync.Mutex) {
			lock.Lock()
			flagCount++
			fmt.Printf("waiting for %d year(s)\n", flagCount)
			lock.Unlock()
			wg.Done()
		}(&wg, &lock)
	}
	wg.Wait()
}
