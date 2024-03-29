# 知识点

函数是一等公民。函数不但可以用于封装代码、分割功能、解耦逻辑，还可以化身为普通的值，在其他函数间传递、赋予变量、做类型判断和转换等，就像切片和字典的值那样。

Go 允许使用者定义一个函数类型：
```go
type Printer func(s string) (n int, err error)

// 作为 Printer 类型的实现
func PrinteToStd(str string) (byteNum int, err error) {
    ...
}
```
只要函数签名一致，就可以认为 PrinterToStd 是 Printer 的实现。

>函数签名是函数的参数列表和结果列表的统称，用来鉴别不同的函数的特征。同时定义了使用者与函数交互的方式。**各个参数与结果名不能算作函数签名的一部分，严格说，函数的名称也不能算作函数签名的一部分。**

只要两个函数的参数列表和结果列表中的元素顺序以及类型一致，就可以说它们是一样的函数，或者说是实现了同一个函数类型的函数。(code: FuncType)

1. 如何编写高阶函数？

告诫函数有以下两个条件：
- 接受其他函数作为参数传入。
- 将其他函数作为结果返回。
只要满足上述任意一点，就可以说这个函数是一个高阶函数。(函数类型属于引用类型，它的值可以为 nil)

编写一个简单的高阶函数。(code: HighLevelFunc)

>卫述语句用来检查关键的先决条件的合法性，并在检查未通过的情况下立即终止当前代码块执行的语句。

2. 如何实现一个闭包？

在一个函数中，存在对外来标识符的引用。(**外来标识符**：既不代表当前函数的任何参数或结果，也不代表函数内部声明的，它是直接从外边拿来的。也叫做自由变量)

闭包体现的是由“不确定”变为“确定”的一个过程。
- 闭包函数因为引用了自由变量，呈现出一种“不确定”的状态，也叫“开放”状态。即：闭包函数内部逻辑并不完整，有一部分逻辑需要自由变量参与完成，但自由变量代表了什么在闭包函数被定义的时候是未知的。
- 直到程序运行到闭包时，自由变量有了确定的值，函数状态才由不确定变为确定。

闭包函数示例：
```go
type operate func(int, int) int
type calculateFunc func(int, int) (int, error)

func genCalculator(op operate) calculateFunc {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("op is nil")
		}
		return op(x, y), nil
	}
}
```
- 使用闭包可以延迟实现一部分程序逻辑或功能。更准确的说：我们在动态地生成那部分程序逻辑。借此可以在程序运行过程中，根据需要生成功能不同的函数。

3. 传入函数的那些数值后来都怎么样了？(code: ChangePassValue)

**所有传给函数的参数值都会被复制**，函数在其内部使用的并不是参数值的原值，而是它的副本。
- 数组是值类型，所以每次作为参数传递都会拷贝它的全部内容。
- 对于引用类型(切片，字典，通道)，作为参数传递也会拷贝它们本身，但不会拷贝它们引用的底层数据。(浅拷贝而不是深拷贝)

# 关于函数传参的注意事项

不要把程序的细节暴露给外界，也尽量不要让外界的变动影响到你的程序。