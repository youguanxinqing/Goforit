# 知识点

Go 语言的链表实现在标准库 container/list 中。其中有两个公开实体：
- List 实现了一个双向链表：
```go
type List struct {
	root Element 
	len  int 
}
```
- Element 链表中的元素结构，含四个字段：
```go
type Element struct {
	next, prev *Element
	list *List
	Value interface{}
}
```

List 类型拥有的方法：
```go
func (l *List) MoveBefore(e, mark *Element)
func (l *List) MoveAfter(e, mark *Element)

func (l *List) MoveToFront(e *Element)
func (l *List) MoveToBack(e *Element)


func (l *List) Front() *Element
func (l *List) Back() *Element

func (l *List) InsertBefore(v interface{}, mark *Element) *Element
func (l *List) InsertAfter(v interface{}, mark *Element) *Element

func (l *List) PushFront(v interface{}) *Element
func (l *List) PushBack(v interface{}) *Element
```
- 链表不允许将自己生成的 Element 插入到一个 List 对象中。
- List 是一个**双向**且**环**的链表。(**私以为环的处理是为了提高追加元素的效率。**)

1. list 与 ring 之间的区别在哪儿？
- Ring 类型的数据结构，仅由自身就可以代表，而 List 类型需要由自己和 Element 类型联合表示。
```go
type Ring struct {
	next, prev *Ring
	Value      interface{} // for use by client; untouched by this library
}
```
- 一个 Ring 类型的值只代表了循环链表中的一个元素，一个 List 类型的值代表了完整的链表。
- 创建并初始化一个 Ring 时，需要指定包含的元素个数，但 List 不行。
- 通过 `var r ring.Ring` 声明的 r 会有一个长度，List 类型的零值长度为 0。
- Ring 的 Len 方法时间复杂度 O(n)；List 的是 O(1)。