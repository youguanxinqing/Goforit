package main

import (
	"fmt"
	"net"
)

func ConnBaidu() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Printf("conn err: %v", err)
	}

	defer func() {
		conn.Close()
	}()

	// 请求百度
	str := "GET /search/error.html HTTP/1.1\r\nConnection: Close\r\n\r\n"

	conn.Write([]byte(str))

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil || n == 0 {
			fmt.Println("\n...")
			break
		} else {
			fmt.Printf("%s", buf[:n])
		}
	}
	fmt.Printf("\n")
}
