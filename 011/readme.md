# 知识点

单向通道，即只能发不能收，或者只能收不能发。初始化一个单向通道：
```go
var uselessChan = make(chan<- int, 1)
```

1. 单向通道有什么应用价值？

**单向通道主要是用来约束其他代码行为的。**(code: SingleSideChan)

`for range` 用法：
```go
intChan2 := getIntChan()
for elem := range intChan2 {
	fmt.Printf("The element in intChan2: %v\n", elem)
}
```
- for 语句会不断地尝试从 intChan2 中取出元素值，即使 intChan2 被关闭，它也会在取出所有剩余的元素值之后再结束执行。
- 当 intChan2 中没有元素值时，它会被阻塞在有 for 关键字的那一行，直到有新的元素值可取。
- 当 intChan2 值为 nil，那么它会被永远阻塞在有 for 关键字的那一行。

2. select 语句与通道怎样联用，应该注意什么？

**select 只能与通道联用，每个 case 都只能包含操作通道的表达式。**
```go
select {
case <-intChans[0]:
	fmt.Printf("index = 0\n")
case <-intChans[1]:
	fmt.Printf("index = 1\n")
case elem := <-intChans[2]:  // 允许短声明赋值
	fmt.Printf("index = 2 %v\n", elem)
default:
	fmt.Printf("no elem\n")
}
```

关于 select 的注意事项：
- 当存在 default case 时，不论涉及通道操作的表达式是否有阻塞，select 语句都不会被阻塞，会走默认分支。
- 如果不存在 default case，当所有 case 都不满足时，select 语句被阻塞。
- 需要对每个 case 对应的通道判断是已经关闭，并采取对应的处理措施。
```go
select {
case _, ok := <- intChan:
	...
}
```
- select 语句只会对其中的每个 case 表达式各求值一次，所以连续操作需要用 for 循环。**注意**，select case 中的 break 只能跳出 select， 不能跳出外层的 for。

3. select case 规则有哪些？(code: SelectDoRuler)

- 每个 case 都必须有一个发送或接受的通道操作。接收表达式有短变量声明时，允许赋值符号左边存在一个或者两个表达式(这类表达式允许被赋值)，当这样的 case 被求值时，它包含的多个表达式总会从左到右的顺序被求值。
- 一个 select 语句中的多个 case 分支，求值时存在先上后下的顺序。
- 当 case 对应的通道操作处于阻塞状态时，就认为该 case 不满足选择条件。
- **仅当 select 语句中所有 case 表达式都被求值完毕后，它才会开始选择候选分支。** 如果没有满足的 case，也没有 default case，select 进入阻塞状态。直到至少有一个 case 满足条件，select 所在 goroutine 被唤醒，进入该 case 执行分支语句。
- 如果 select 语句发现多个 case 都满足选择条件，那么它会用一种伪随机算法在这些分支中选择一个并执行。
- 一条 select 语句只能有一个 default case。
- select 语句的每次执行，包括 case 表达式求值和分支选择，都是独立的。

4. 如何永久避开一个分支？(code: AvoidOneCase)

将其置为 nil：
```go
select {
	case <-c1:
		...
	case _, ok := <-c2:
		if !ok {
			c2 = nil 
			...
		}
}
```