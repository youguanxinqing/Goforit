# 知识点

1. range 子句的使用细节？

```go
func main() {
	arr := [...]int{1, 2, 3, 4}
	for i, e := range arr {
		if i == len(arr) - 1 {
			arr[0] += e
		} else {
			arr[i + 1] += e
		}
	}
	fmt.Println(arr)
}
```
- 输出：[5 3 5 7]。
- 原因：range 迭代对象的副本，而不是原值。(code: ForRange)

range 的注意点：
- range 表达式只会在 for 语句**开始执行时被求值一次**，无论后边会有多少次迭代；
- range 表达式的**求值结果会被复制**，被迭代的对象是 range 表达式结果值的**副本而不是原值**。

2. swicth 与 case 之间的注意事项。(code: UseSwitch)

```go
// xxx1 是 switch 表达式
// xxx2 是 case 表达式
switch xxx1 {
    case: xxx2
}
```

```go
arr := [...]int8{1, 2, 3}
switch 1 + 2 {  // 1 + 2 属于无类型常量
case arr[0]:
	fmt.Println(arr[0])
case arr[1]:
	fmt.Println(arr[1])
case arr[2]:
	fmt.Println(arr[2])
}
// 不能正常被编
```
- **如果 switch 表达式的结果值是无类型的常量**，该常量会被自动转换这种常量的默认类型，如 4 -> int；3.14 -> float64。
- int 与 uint8 不是一种类型，不能判等，因而编译报错。
- **如果 case 表达式的结果是无类型的常量**，switch 表达式是什么类型，case 表达式就会自动转换成该类型，然后做判等。如果类型转换失败，编译报错。

```go
arr := [...]int8{1, 2, 3}
switch arr[1] {
case 1:
    fmt.Println(1)	
case 2:
    fmt.Println(2)	
case 3:
    fmt.Println(3)	
}
```

3. switch 语句对 case 的约束。

switch 语句在 case 子句的选择上具有**唯一性**。

switch 语句**不允许** case 表达式中的子表达式结果值存在相等的情况，不论这些结果值相等的子表达式，是否存在不同的 case 表达式中。(**只针对常量的子表达式**)
```go
// 不允许操作
switch num {
case 1, 2:
    fmt.Println("1 or 2")
case 2, 3:
    fmt.Println("2 or 3")
}

// 允许操作
arr := [...]int{1, 2, 3}
num := 2
arr := [...]int{1, 2, 3}
switch num {
case arr[0], arr[1]:
    fmt.Println("index 0 or 1")
case arr[1], arr[2]:
    fmt.Println("index 1 or 2")
}
```
- 最上边的 case 子句中的子表达式总是被最先求值，判等顺序也是这样。因此，如果某些子表达式的结果值有重复并且它们与 switch 表达式的结果值相等，那么位置靠上的 case 子句总会被选中。

对类型判断的 switch 语句无效，因为类型判断时，case 表达式中必须使用类型的字面量。

```go
switch t := interface{}(num).(type) {
case float32:
    fmt.Println("float32")
case int:
    fmt.Println("int")
default:
    fmt.Printf("do not know %T\n", t)
}
```

**由于 byte 是 uint8的别名类型，因而也不允许同时出现**：
```go
// 错误示例
switch t := interface{}(num).(type) {
case byte:
    ...
case uint8:
    ...
}
```