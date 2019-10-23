package main

import (
	"fmt"
	"sync"
)

func main() {
	name := "zhong"

	var RWLock sync.RWMutex

	sendCond := sync.NewCond(&RWLock)
	recvCond := sync.NewCond(RWLock.RLocker())

	go func() {
		RWLock.Lock()
		for name == "zhong" {
			// wait 是用来阻塞自己的（当前 goroutine）
			fmt.Println("write lock block itself.")
			sendCond.Wait()
		}
		fmt.Println("write lock continue doing.")
		name = "zhong"
		fmt.Println("write lock change name value.")
		RWLock.Unlock()
		fmt.Println("write lock notify read lock.")
		recvCond.Signal()
	}()

	go func() {
		RWLock.RLock()

		/*
		** 向下执行的两个条件：1. name=zhong 2. recv notify
		 */
		for name != "zhong" {
			fmt.Println("read lock block itself.")
			recvCond.Wait()
		}
		fmt.Println("read lock continue doing.")
		name = "ding"
		fmt.Println("read lock change name value.")
		RWLock.RUnlock()
		fmt.Println("read lock notify write lock.")
		sendCond.Signal()
	}()

	for {

	}
}
