package main

import "errors"

type operate func(int, int) int
type calculateFunc func(int, int) (int, error)

func genCalculator(op operate) calculateFunc {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("op is nil")
		}
		return op(x, y), nil
	}
}

func main() {

}
