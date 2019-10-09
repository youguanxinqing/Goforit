package main

import "fmt"

func main() {
	var num int16 = 255
	fmt.Println(int8(num)) // -1

	var numUnderZero int16 = -255
	fmt.Println(int8(numUnderZero)) // 1

	var numUTOuint uint16 = 255
	fmt.Println(uint8(numUTOuint)) // 255

	var numUTOint uint16 = 255
	fmt.Println(int8(numUTOint)) // -1
}