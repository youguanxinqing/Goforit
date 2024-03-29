# 知识点

sync.Pool 类型被称为**临时对象池**。属于结构体类型，它的值被真正使用之后，就不应该再被复制。
- 临时对象是指：**不需要持久使用的某一类值**。它们的创建与销毁可以在任何时候发生，完全不会影响程序功能。它们无需被区分，其中的任何一个值都可以代替另一个。
- Put 在池中存放临时对象；Get 从池中获取临时对象。
- Get 可能会从当前池中删除任何一个值，然后把这个值作为结果返回。如果池中没有任何值，会使用 New 字段新建一个返回(这个结果值不会被存入当前的临时对象池中，而是直接返回给 Get 方法的调用方)。
- New 字段的实际值在初始化临时对象池的时候给定。

1. 为什么说临时对象池中的值会被及时清理掉？

Go 语言运行时系统中的垃圾回收器，在每次开始执行之前，都会对所有已创建的临时对象池中的值进行全面清除。(?)

sync 包被初始化的时候，会向 Go 语言运行时系统注册一个函数。该函数用于清除所有已创建的临时对象池中的值。
- 当函数注册后，Go 语言运行时系统每次即将执行垃圾回收时，都会先执行这个函数。

sync 包中有一个包级私有的全局变量，这个变量是所有临时对象池的汇总。类型为 *sync.Pool 的切片，可以称为**池汇总列表**。
- 每个临时对象池的 Put 方法或 Get 方法第一次被调用时，该池就会被添加到池汇总列表中。
- 清理函数遍历池汇总列表，获得所有临时对象池。清理函数**先将池中所有的私有临时对象和共享临时对象列表**置为 nil。然后把所有本地池列表都销毁。最后，将池汇总列表重置为空切片。
- 如果临时对象池以外的代码对临时对象无引用，那么在垃圾回收过程中，临时对象会被销毁。

2. 临时对象池使用怎样的数据结构？

临时对象池的数据结构顶层，称为**本地池列表**。其实就是数组，长度总是与 Go 语言调度器中的 P 的数量相同。
- P 存在的意义：分散并发程序的执行能力。

本地池列表中的每个本地池包含三个字段：存储私有临时对象的字段 private，共享临时对象列表的字段 shared，以及 sync.Mutex 嵌入字段。

3. 临时对象池怎样利用内部数据结构存取值的？

Put() 总试图把新的临时对象，存储到对应的本地池的 private。但 private 已经有某个值时，会尝试存 shared。

shared 字段可以被任何 goroutine 中的代码访问(不论 goroutine 关联的是哪个 P)。本地池的 private 字段只能被与之对应的 P 所关联的 goroutine 中的代码访问。

由于 shared 字段是共享的，因此需要互斥锁保护。Put() 在互斥锁保护下，将新的临时对象追加到共享临时对象列表的末尾。Get() 会在互斥锁的保护下，试图从共享临时对象列表中取出最后一个元素。

**临时对象池存储的临时对象，都应该拥有较长生命周期，且不应该被某个 goroutine 长期持有和使用。**