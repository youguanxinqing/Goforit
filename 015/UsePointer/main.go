package main

import (
	"fmt"
	"Goforit/015/DogPack"
)

func main() {
	dog := DogPack.New("aaaa")
	fmt.Println(dog)

	// p := DogPack.GetNameStr(dog)
	p := DogPack.GetNameStrA(&dog)
	fmt.Println(p)
	*p = "bbbb"
	fmt.Println(dog)
}