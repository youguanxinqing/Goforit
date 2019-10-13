package main

import "fmt"

type Animal struct {
	People
}

func (p Animal) Walk() {
	fmt.Printf("Animal walked \n")
}

type People struct {
}

func (p People) Walk(speed int) {
	fmt.Printf("people walked by %d m/s\n", speed)
}

func main() {
	a := Animal{}
	a.Walk()
	a.Walk(10) // error: too many arguments in call to a.Walk
}
