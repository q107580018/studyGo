package main

import "fmt"

// Humaner 定义接口类型
type Humaner interface {
	// 方法，只有声明，没有实现
	SayHi()
}

type Teacher struct{}

func (t *Teacher) SayHi() {
	fmt.Println("Teacher say hi!")
}

type Student struct{}

func (s *Student) SayHi() {
	fmt.Println("Student say hi!")
}

func main() {
	var i Humaner
	i = &Teacher{}
	i.SayHi()
	i = &Student{}
	i.SayHi()
}
