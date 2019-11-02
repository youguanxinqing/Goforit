package main

import (
	"sync"
	"time"
)

func main() {
	ConnBaidu()

	return
	var wait sync.WaitGroup

	wait.Add(2)
	go MySever(wait.Done)
	time.Sleep(time.Second * 2)
	go MyClient(wait.Done)
	wait.Wait()
}
