package main

import "fmt"

// Humaner 定义接口类型
type Humaner interface {
	// 方法，只有声明，没有实现
	SayHi()
}

// Teacher ...
type Teacher struct{}

// SayHi ...
func (t *Teacher) SayHi() {
	fmt.Println("Teacher say hi!")
}

// Student ...
type Student struct{}

// SayHi ...
func (s *Student) SayHi() {
	fmt.Println("Student say hi!")
}

// Mystr ...
type Mystr string

// SayHi ...
func (s *Mystr) SayHi() {
	fmt.Println("Mystr say hi!")
}

// WhoSayHi ...
func WhoSayHi(h Humaner) {
	h.SayHi()
}

func main() {
	s := &Student{}
	t := &Teacher{}
	str := Mystr("我")
	// 定义一个普通函数，参数为接口类型,一个函数，不同表现，多态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	// 创建切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str
	for _, v := range x {
		WhoSayHi(v)
	}
}
