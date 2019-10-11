package main

import "fmt"

func RetriveData(c chan int) {
	for {

		if v, ok := <-c; ok {
			fmt.Printf("cur value is %v\n", v)
		} else {
			fmt.Printf("chan is already closed.\n")
			goto end
		}
	}

end:
	fmt.Printf("over\n")
}

func AddData(c chan int) {
	for _, i := range []int{1, 2, 3} {
		if i > 2 {
			close(c)
		} else {
			c <- i
			c <- i + 1
		}
	}
}

func main() {
	c := make(chan int, 10)
	go AddData(c)
	go RetriveData(c)
	for {
	}
}
