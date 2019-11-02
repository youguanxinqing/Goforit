package main

import (
	"fmt"
	"net/http"
)

func main() {

	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Printf("err : %v\n", err)
	} else {
		fmt.Printf("status %v\n", resp.Status)
		for {
			buf := make([]byte, 1024)
			n, err := resp.Body.Read(buf)
			if err != nil {
				fmt.Printf("read err or over: %v\n", err)
				break
			} else {
				fmt.Printf("%v", string(buf[:n]))
			}
		}
		fmt.Printf("\n")
	}
}
