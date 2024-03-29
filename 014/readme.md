# 知识点

接口类型与其他数据类型不同，不能实例化。

一个接口的方法集合就是它的全部特征。对于任何数据类型，只要它的方法集合完全包含一个接口的全部特征，那么它就这个接口的实现类型。(这种无情侵入式的接口实现方法又叫做鸭子类型 Duck Typing)

1. 如何判断一个数据类型的某个方法实现的就是某个接口类型中的某个方法呢？

要求：
- 两个方法的签名需要完全一致。
- 两个方法的名称要一摸一样。

2. 动态类型与动态值

```go

// 前提：假定结构体类型 Dog，实现了接口类型 Pet

dog := Dog{"little dog"}
var p Pet = &dog
```
- p Pet 类型的变量(Pet 是 p 的静态类型)。
- p 的动态类型 Dog，动态值是 &dog。
- 对 p 来说，它的静态类型永远不会改变，但它的动态类型会随着赋值给它的动态值改变而改变。
- **接口类型的变量被赋予实际值之前，它的动态类型是不存在的。**

3. 为一个接口变量赋值时会发生什么？

```go
dog := Dog{"little dog"}
var p Pet = dog  // Dog 的值类型是接口 Pet 的实现
dog.setName("monitor")
```
- p 变量中的 name 字段依然是 “little pig”。
- 当一个变量给另一个变量赋值，那么真正赋给后者的，并不是前者，而是前者的一个副本。

接口本身不能被值化，所以在实际赋值之前，接口类型的变量一定是 nil，这是它的零值。

**当给一个接口变量赋值时，该变量的动态类型和动态值一起被存储到一个专用的数据结构中。**严格地说，这样一个变量的值是专用数据结构的一个实例，而不是赋值给该变量的实际值。可以认为 p 的值中包含了 dog 值的副本。

专用数据结构叫做 ifrace，其实例会包含两个指针，一个指向**类型信息**，另一个指向动态值。这里的类型信息由另一个专用数据结构的实例承载，包含了动态值的类型，实现的接口方法以及调用途径，等等。

4. 接口变量的值在什么情况下才真正为 nil?(code: InterfaceVar)

**只要把一个有类型的 nil 赋值给接口变量，那么这个变量的值就一定不会是真正的 nil。** p == nil 会返回 false。

当接口变量声明但不初始化时，或者直接赋值 nil，就是真正的 nil。

5. 如何实现接口之间的组合？

接口之间的嵌入比结构体之间的嵌入更简单，因为它**不会设计方法间的“屏蔽”**。只要组合的接口之间有同名的方法就会产生冲突，从而无法通过编译，即使同名方法的签名不同。

# 补充(code: NilInterfaceVar)

把一个值为 nil 的某个实现类型的变量赋给接口变量，这个接口变量可以调用该接口的方法。但是方法内不是使用实现类型内的变量。方法的接收者必须是指针。