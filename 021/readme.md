# 知识点

大多数操作系统中，退出状态码不是 0，基本意味着程序运行非正常结束。

1. 从 panic 被引发到程序终止运行，过程中发生了什么？

程序引发 panic，panic 详情会被建立，程序控制权由当前代码行转至所属函数(调用栈中的上一级)，控制权继续往上一级转移，直到最外层函数。

error 与 panic 的意义不同。返回错误值时，函数调用方有权选择不处理，并且不处理的后果往往是不致命的。

2. 如何让一个 panic 包含一个值，以及应该包含什么样的值？

```go
// go source code
func panic(v interface{})
```

panic() 允许传入任何类型的值，但一般是传入 error 类型的错误值。

3. 怎样施加对 panic 的保护措施，避免程序崩溃？

使用内建函数 recover()。
```go
func main() {
	defer func() {
		if p:=recover(); p!=nil {
			fmt.Println(p)
		}
	}()

	panic(errors.New("where is zhong?"))
}
```

defer 中不允许 Go 语言内建函数的调用表达式，以及针对 unsafe 包中的函数的调用表达式。

尽量把 defer 语句写在函数体的开始处。

4. 如果函数中有多条 defer 语句，那么 defer 的调用顺序如何？

从下到上。

```go
func main() {
    defer func() {
        ... (1)
    }()

    panic(errors.New(...))

    defer func() {
        ... (2)
    }()
}
```
- 上述情况下，(2) 处的代码不会被执行。