package main

import (
	"fmt"
	"container/list"
)

func PrintList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// use *List method of pushback
func UsePushBack() {
	nums := list.New()  // *list.List
	fmt.Printf("%T\n", nums)

	nums.PushBack(1)
	nums.PushBack(2)
	nums.PushBack(3)

	PrintList(nums)	
}

// use *List method of pushfront
func UsePushFront() {
	nums := list.New()

	nums.PushBack(1)
	nums.PushBack(2)
	nums.PushFront(3)  // 链表最前插入

	PrintList(nums)
}


func main() {
	// UsePushBack()
	UsePushFront()
}
