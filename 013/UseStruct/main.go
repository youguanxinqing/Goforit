package main

import "fmt"

// Colors red, green, blud
type Colors struct {
	red   string
	green string
	blue  string
}

func (c Colors) String() string {
	return fmt.Sprintf("%s %s %s", c.red, c.green, c.blue)
}

func main() {
	cls := Colors{green: "green", blue: "blue"}
	fmt.Println(cls)
}
