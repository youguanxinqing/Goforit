[root path](../readme.md)
# 知识点

**库源码文件**是不能直接运行的源码文件。作用：存放程序实体(Go 语言中程序实体是变量，常量，函数，结构体和接口的统称)。

如果该目录下有一个命令源码文件，那么为了让同一目录下的文件都能通过编译，这些源码文件都应该声明属于 main 包。(code: mainbag)

在 mainbag 目录下执行 go build，得报错信息如下：
`can't load package: package gostudy/003/mainbag: found packages hello (hello.go) and main (main.go) in /Users/guan/Mine/Code/go/src/gostudy/003/mainbag`

**源码文件声明的包名(package xxx)可以与其所在目录名称不同**，但必须要保证同一目录下的包名一致。

**库源码文件包名与目录名不一致时**，执行 go build，pkg 目录下生成的是 dirname.a。**命令源码文件包名与目录名不一致时**，生成的可执行文件是 dirname。

1. 代码包的导入路径总与其所在目录的相对路径一致吗？(code: import_when_notrequal)

当包名与目录名不一致时，正确的导入办法--起别名：
```go
package main

import zhong "gostudy/003/import_when_notrequal/zhonghello"

func main() {
	zhong.Hello()
}
```

**源码文件对 src 目录的相对路径就是它的包导入路径，实际使用时需要限定符，这个符号默认是包名**。

2. 对于程序实体，还有其他的访问权限规则吗？

创建 internal 代码包可以让一些实体仅能被当前模块中的其他代码引用。叫做：模块级私有。
- internal 代码包中声明的**公开程序实体**只能被该代码包的**直接父级**及**其子包**中的代码引用。

更详细的说明(《Go 语言圣经》)：
>为了满足这些需求，Go语言的构建工具对包含internal名字的路径段的包导入路径做了特殊处理。这种包叫internal包，一个internal包只能被和internal目录有同一个父目录的包所导入。例如，net/http/internal/chunked内部包只能被net/http/httputil或net/http包导入，但是不能被net/url包导入。不过net/url包却可以导入net/http/httputil包。

```
net/http
net/http/internal/chunked
net/http/httputil
net/url
```

# 代码包最后一级同名 补充(code: samebagname)

当出现 dep/lib/flag 和 flag 包都需要导入时，会产生冲突。报错如下：
![](/003/png/QQ20191007-225043@2x.png)

正确的处理方式是取别名：
```go
import (
	"gostudy/003/samebagname/ting"
	zt "gostudy/003/samebagname/zhong/ting"
)
```
[root path](../readme.md)
