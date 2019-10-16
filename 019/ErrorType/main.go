package main

import (
	"errors"
	"fmt"
)

func echo(request string) (string, error) {
	if request == "" {
		return "", errors.New("empty request")
	}
	return request, nil
}

func main() {
	requests := []string{"", "hello"}
	for _, r := range requests {
		if resp, err := echo(r); err != nil {
			continue
		} else {
			fmt.Printf("response : %s\n", resp)
		}
	}
}
