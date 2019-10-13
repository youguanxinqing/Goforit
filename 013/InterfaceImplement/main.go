package main

import "fmt"

type Walker interface {
	walk()
}

type Flyer interface {
	fly()
}

type People struct {
}

func (p *People) walk() {
	fmt.Println("people walk...")
}

func (p People) fly() {
	fmt.Println("people fly...")
}

// func (p People) walk() {
// 	fmt.Println("people walk...")
// }

func WalkAct(w Walker) {
	w.walk()
}

func FlyAct(f Flyer) {
	f.fly()
}

func main() {

	var p *People

	// 空指针允许调用指针方法，不能调用值方法
	p.walk()
	p.fly() // 编译不会报错，运行报 panic

	// 空值既可以调用值方法，也能调用指针方法
	var v People
	v.walk()
	v.fly()

	// 只实现了指针类型的 walk 方法
	WalkAct(p)  // 指针类型才是 Walker 接口的实现
	WalkAct(&v) // WalkAct(v) 编译异常，值类型不是 Walker 的实现

	// 只实现了值类型的 fly 方法
	FlyAct(p) // 运行时报错，但指针类型是 Flyer 的实现
	FlyAct(v) // 值类型是 Flyer 的实现

}
