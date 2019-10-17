package main

import (
	"fmt"
	"errors"
)

func main() {
	defer func() {
		if p:=recover(); p!=nil {
			fmt.Println(p)
		}
	}()

	panic(errors.New("where is zhong?"))

}