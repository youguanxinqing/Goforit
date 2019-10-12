package main

import "fmt"

type Person interface{
	Name()
}

type T struct {
	name int
}

func (t *T) Name(){
	fmt.Println("132")
}

func ffff() Person {
	var t *T
	return t 
}

func defer1(n int)(t int){
	t = n
	defer func(t int){
		t += 3
	}(t)
    return 3
}

func main(){
	fmt.Println(defer1(2))
	var t *T
	t.Name()
	fmt.Println(t.name)
}