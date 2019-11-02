package main

import (
	"fmt"
	"net"
)

// MySever is server
func MySever(over func()) {
	defer func() {
		over()
	}()

	// 监听
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("[server] listen err(%v)\n", err)
		return
	}

	defer func() {
		l.Close()
	}()

	for {
		// 等待连接
		c, err := l.Accept()
		if err != nil {
			fmt.Println("[server] accept err")
			continue
		}

		for {
			buf := make([]byte, 1024)
			// 接收数据
			n, err := c.Read(buf)
			if err != nil {
				fmt.Println("[server] read err")
				goto end
			} else {
				fmt.Printf("[server] content: %s\n", buf[:n])
			}
		}
	}

end:
	fmt.Println("[server] client leave")
}

// MyClient is client
func MyClient(over func()) {
	defer func() {
		over()
	}()

	// 连接服务器
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Printf("[client] dial err(%v)\n", err)
		return
	}

	defer func() {
		conn.Close()
	}()

	// 发送数据
	n, err := conn.Write([]byte("ding kai hui"))
	if err != nil {
		fmt.Println("[client] write err")
	} else {
		fmt.Printf("[client] send n(%d) byte\n", n)
	}
}
