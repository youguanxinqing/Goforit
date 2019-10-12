package main

import "fmt"

// Printer define func type
type Printer func(s string) (n int, err error)

// PrinterToStd impletement print func
func PrinterToStd(str string) (bytesNum int, err error) {
	return fmt.Println(str)
}

func main() {
	var p Printer
	p = PrinterToStd
	p("zhong")
}
