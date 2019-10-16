package main

import (
	"errors"
	"net"
	"os"
	"os/exec"
)

func main() {
	errors.New()

	// 类型已知
	os.PathError
	os.LinkError
	os.SyscallError
	exec.Error

	// 类型相同
	os.ErrClosed
	os.ErrInvalid
	os.ErrPermission

	// 没有相应变量，且类型未知，只能用错误信息的字符串表示形式来做判断
	os.IsExist()      // 判断错误是否因为存在
	os.IsNotExist()   // 判断错误是否因为不存在
	os.IsPermission() // 判断错误是否属于权限问题

	net.Error
	net.DNSConfigError
	net.Conn
}
