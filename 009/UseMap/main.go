package main

import "fmt"

func main() {
	var myMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	k := "two"
	v, ok := myMap[k]
	if ok {
		fmt.Printf("k is %s and v is %d\n", k, v)
	} else {
		fmt.Printf("k(%s) not exists", k)
	}

	k2 := "four"
	if v2, ok := myMap[k2]; ok {
		fmt.Printf("k is %s and v is %d\n", k2, v2)
	} else {
		fmt.Printf("k(%s) not exists\n", k2)
	}
}
