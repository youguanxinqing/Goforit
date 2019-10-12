package main

import "fmt"

func changeData(s [2][]string) [2][]string {
	s[0][1] = "g"
	return s
}

func main() {
	sliceInArr := [2][]string{
		[]string{"z", "t", "y"},
		[]string{"d", "k", "h"},
	}
	fmt.Printf("s %v\n", sliceInArr)
	sChanged := changeData(sliceInArr)
	fmt.Printf("sChanged %v\n", sChanged)
	fmt.Printf("s %v\n", sliceInArr)

	/* output */
	// 	s [[z t y] [d k h]]
	// sChanged [[z g y] [d k h]]
	// s [[z g y] [d k h]]  发生改变
}
