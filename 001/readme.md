[root path](../readme.md)

# 知识点
GOROOT：Go 语言安装根目录的路径，也就是 GO 语言的安装路径。
GOPATH：若干工作区目录的路径。是我们自己定义的工作空间。
GOBIN：GO 程序生成的可执行文件（executable file）的路径。

1. go 源码组织方式
代码包的名称一般会与源码文件所在的目录同名。如果不同名，那么在构建、安装的过程中会以代码包名称为准。

2. 源码安装结果
归档文件的相对目录与 pkg 目录之间还有一级目录，叫做**平台相关目录**。平台相关目录的名称是由 build（也称“构建”）的**目标操作系统**、**下划线**和**目标计算架构**的代号组成的。
- eg: `darwin_amd64`
- 格式：![结构](/001/png/2fdfb5620e072d864907870e61ae5f3c.png)

3. 构建与安装
go build 构建，go install 安装。构建和安装代码包**都会执行编译、打包等操作**，并且，这些操作生成的任何文件都会**先保存到某个临时的目录中**。
- 如果**构建的库源码**文件，操作后产生的结果文件**只会存在临时目录中**(pkg)。这里的构建的主要意义在于检查和验证。
- 如果**构建的是命令源码**文件，操作后的结果文件**会被搬运到源码文件所在的目录中**($GOPATH/bin, $GOBIN/bin)。


# go build 补充
go build 默认不会编译目标代码包依赖的其他代码包，除非依赖的代码包的归档文件不存在，或者源码文件有了变化(? 这里的源码文件是说目标呢还是依赖呢)。
- 经过验证，当依赖包中的源码变动时，依赖代码包会被重新编译。
- 如果目标代码包出现变动，依赖包没有变动，不会重新编译依赖代码包。
- `-a` 实现强制编译，即使依赖的是标准库中的代码包也是如此。

go build -n 表示会有哪些操作，但不会执行。
go build -x 可见看到有哪些具体操作被执行(go build -n 的执行版本)。
go build -v 能够显示编译代码包的名称。


# go get 补充
命令 go get 从一些主流公用代码仓库（比如 GitHub）下载目标代码包，并把它们安装到环境变量GOPATH包含的第 1 工作区的相应目录中。**如果存在环境变量GOBIN，那么仅包含命令源码文件的代码包会被安装到GOBIN指向的目录。**
- -u 下载并安装代码包，不论工作区中是已经否存在。
- -d 只下载代码包，不安装。
- -fix 在下载代码包后先运行一个用于根据当前 Go 语言版本修正代码的工具，然后再安装代码包。
- -t 同时下载测试所需的代码包。
- -insecure 允许通过非安全的网络协议下载和安装代码包。HTTP 就是这样的协议。


# 其他
为避免存储源码的代码仓库或者代码包的相对路径改变，影响到其他代码包的导入，**可自定义导入路径**。
```go
package semaphore // import "golang.org/x/sync/semaphore"
```
- 注释后面的路径并不真实存在，但可以使用。如：go get "golang.org/x/..."，之后会去真实路径 github 上下载，但还需要在 golang.org 的后端做些设置。


[root path](../readme.md)