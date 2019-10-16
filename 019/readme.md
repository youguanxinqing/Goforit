# 知识点

errors.New() 返回 *errorString 类型。对于 error 类型的值来说，Error() 就相当于其他类型的 String。

```go
// errors source code
package errors

func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```
- error 接口类型要求实现 Error() 即可。

当需要通过模版化的方式生成错误信息，并得到错误值时，使用 `fmt.Errorf()`。

1. 对于具体错误的判断，Go 语言中有哪些惯用用法(怎样判断一个错误值具体代表的是哪一类错误)。

- 对类型已知范围内的一系列错误值，一般使用类型断言表达式或类型 switch 语句判断。

这一类错误字符串描述可能存在变动，但 Err 信息在官方库中已经声明。这一类错误的类型不同，因而可以根据类型区分错误类型。
```go
// Errors
const (
	E2BIG           = Errno(0x7)
	EACCES          = Errno(0xd)
	EADDRINUSE      = Errno(0x30)
	EADDRNOTAVAIL   = Errno(0x31)
	EAFNOSUPPORT    = Errno(0x2f)
	EAGAIN          = Errno(0x23)
	EALREADY        = Errno(0x25)
	EAUTH           = Errno(0x50)
    EBADARCH        = Errno(0x56)
    ...
```
- 对于已有相应变量且类型相同的一些列错误值，一般使用判等操作来判断。

这一类错误在官方库已经存在，没有独立类型。通过函数的层层封装，最后都会 errors.New()，错误信息各不相同，所以可以通过判等进行区分。
```go
var (
	ErrInvalid    = errors.New("invalid argument")
	ErrPermission = errors.New("permission denied")
	ErrExist      = errors.New("file already exists")
	ErrNotExist   = errors.New("file does not exist")
	ErrClosed     = errors.New("file already closed")
)
```
- 对于没有相应变量且类型未知的一系列错误值，只能使用错误信息的字符串表示形式来做判断。
(暂时不太明白是什么意思)

2. 如何根据实际情况，给予恰当的错误值？

构建错误的基本方式：
- 创建立体的错误类型体系
- 创建扁平的错误值列表

error 是接口类型，所以 errors.New() 生成的错误值只能赋给变量，不能赋给常量。

Errno 类型代表了系统调用时可能发生的底层错误，这个错误类型是 error 接口的实现类型，同时也是 uintptr 的再定义类型。