package main

import "fmt"

// Biology 生物
type Biology struct {
	Animal
	People
}

// People 人类
type People struct {
	male   string
	female string
}

func (p People) String() string {
	return fmt.Sprintf("{ male: %s, female: %s}", p.male, p.female)
}

func (p People) Walk(n int) {
	fmt.Println("people walk...")
}

type Animal struct {
}

func (a Animal) String() string {
	return "no animal"
}

func (a Animal) Walk() {
	fmt.Println("animal walk...")
}

func main() {
	bio1 := Biology{Animal{}, People{"jay", "ting"}}
	bio2 := Biology{People: People{male: "jay", female: "ting"}}

	fmt.Println(bio1)
	fmt.Println(bio2)
	// bio1.Walk() // ambiguous selector bio1.Walk
	bio1.People.Walk(1)
}
