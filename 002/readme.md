[root path](../readme.md)
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

对 flag.Usage 重新赋值(flag.Usage 是一个函数)。
```go
func main() {
    flag.Usage = func() {
        ...
        flag.PrintDefaults()  // 输出参数信息
    }
    
    flag.Parse()
}
```

3. 更灵活的自定义 commandline(code: commandline, yourselfcmd)

os.Args 负责接收命令行的中的参数。索引为 0 的位置是可执行文件的全路径。
```go
fmt.Println(os.Args[1:])

// 运行 go run main.go -name=zhong -gender=girl
// 打印
[-name=zhong -gender=girl]
```

在调用 flag.StringVar(), flag.Parse() 时，其实是在调用 flag.CommandLine 中对应的方法。

因此自定义还可以这样写：
```go
flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
flag.CommandLine.Usage = func() {
    ...
}
```
- flag.NewFlagSet() 中的第二个参数是常量。**flag.ExitOnError** 表示：当命令跟 --help 或者参数不正确时，打印 Usage 信息后以状态码 2 结束当前程序。**flag.PanicOnError** 表示：最后抛出运行时恐慌。
![Exit VS Panic](/two/png/108B0E92-40DD-4CAA-9402-B7D6AE8D1C6A.png)

上述处理方式会改变 flag.CommandLine 的默认行为，更推荐的方式是创建一个自己的 CommandLine：
```go
var cmdline = flag.NewFlagSet("", flag.ExitOnError)
cmdline.Usage = func() {
    ...
}
cmdline.Parse(os.Args[1:])
```

相应的 flag.xxx 都应该修改为 cmdline.xxx。

# go doc 与 godoc 的补充
[go doc与godoc](https://github.com/hyper0x/go_command_tutorial/blob/master/0.5.md)

本地启动帮助文档：
```go
godoc -http=:8080
```

终端包查询：
```go
go doc flag
```
- -c 大小写敏感
- -u 同时输出不可导出函数

# 使用自定义的数据结构 补充(code: yourds)

自定义的数据结构需要实现接口：
```go
type Value interface {
	String() string
	Set(string) error
}
```

解析命令行参数：
```go
func MyStrVar(m *MyStr, name string, val string, usage string) {
	flag.CommandLine.Var(m, name, usage)
}
```
[root path](../readme.md)
