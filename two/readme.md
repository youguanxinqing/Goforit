# 知识点
源码文件分为：命令源码文件，库源码文件，测试源码文件。

**命令源码文件**是程序的运行入口，**每个可独立运行的程序必须拥有**。我们可以通过构建或安装，生成与其对应的可执行文件，后者一般会与该命令源码文件的直接父目录同名。

**如果一个源码文件声明属于main包，且包含一个无参数声明且无结果声明的 main 函数，那么它就是命令源码文件。**
```go
package main

func main() {
    ...
}
```
> 对于一个独立的程序来说，命令源码文件永远只会也只能有一个。


1. 命令源码文件如何接收参数(code: extractargs)

Go 语言中用于接收与解析参数的包：flag。

命令行参数解析方式：
- flag.StringVar()
```go
var name string
flag.StringVar(&name, "name", "everyone", "person object")
```
- flag.String()
```go
var name string = flag.String("name", "everyone", "person object")
```

第一个参数存放命令行传参，第二个参数是命令行参数的名字，第三个参数是默认值，第四个参数起说明作用。

命令行传参方式：
```bash
go run main.go -h  # 查看帮助信息

go run main.go -name=zhong
# zhong hello world!
```


2. 自定义 help 信息(code: changehelpinfo)

对 flag.Usage 重新赋值。