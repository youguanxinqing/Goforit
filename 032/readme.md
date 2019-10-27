# 知识点

1. 怎样使用 context 包中的程序实体，实现一对多的 goroutine 协作流程？(code: UseContext)

Context 类型提供一类代表上下文的值。此类值并发安全，它可以被传播给多个 goroutine。Context 类型是接口类型，而 context 包中实现该接口的所有私有类型，都是基于某个数据类型的指针类型，所以传播并不会影响该类型值的功能和安全。

一个 Context 值产生出任意个子值。这些子值可以携带其父值的属性和数据，也可以响应我们通过其父值传达的信号。

所有的 Context值共同构成了一颗代表了上下文全貌的树形结构。这棵树的树根（或者称上下文根节点）是一个已经在context包中预定义好的Context值，它是全局唯一的。通过调用 context.Background 函数，我们就可以获取到它。(上下文根节点仅仅是一个最基本的支点，不提供任何额外功能。也就是说，它既不可以被撤销（cancel），也不能携带任何数据。)

繁衍 Context 值的函数：
- WithCancel
- WithDeadline
- WithTimeout
- WithValue
(Deadline 与 Timeout 都会产生定时取消的 parent 子值; Value 产生携带额外数据的子值)

2. “可撤销的”在 context 包中代表着什么？“撤销”一个 Context 值又意味着什么？

对于一个未包含任何元素的值的通道来说，它的关闭会使任何针对它的操作立即结束。

3. 撤销信号如何在上下文树中传播？

WithCancel() 会产生两个结果值，第一个是可撤销的 Context 值，第二个是触发撤销信号的函数。调用撤销函数之后，对应的 Context 值会先关闭它内部的接收通道(Done对应的通道)，然后向它的所有子值传达撤销信号，这些子值的行为又与父节点一致。最后，**Context 值会断开它与父值之间的关联**。

WaitTimeout() 和 WaitDeadline() 生成的 Context 值**可手动撤销，也可以定时撤销**。

WaitValue() 生成的 Context 值不可以撤销。撤销信号被传播时，会跳过它们。
![取消信号的传播](/032/png/a801f8f2b5e89017ec2857bc1815fc9e.png)

4. 怎样通过 Context 值携带数据，怎样从中获取数据？(code: UseWithValue)

WithValue() 产生新的 Context 需要三个参数：父值，键，值。**键的类型必须可判等**。

调用 Value() 会去当前 Context 中找，如果没找到，往父值中找。**只有 WithValue() 返回允许携带数据**，其他的不可以，因此往上配对键的时候，遇到其他 Context，直接略过。最后都没找到返回 nil。(由于 Context 的父值以嵌入的方式进入子值，因此如果当前 Context 值没有 Value 方法，那么调用的就是父值或者祖辈值的 Value)

Context 不提供改变数据的方法，只能通过上下文树添加节点存储新的数据，或者撤销该值的父值丢掉相应数据。**如果存储的数据可以从外部改变，则要自行保证安全。**

# 传递信号是广度优先还是深度优先 补充

**深度寻找**。也就是说，不会去找与父值同级的节点。(code: SearchByWay)