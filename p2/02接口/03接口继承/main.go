package main

import "fmt"

// Humaner 定义接口类型
type Humaner interface {
	// 方法，只有声明，没有实现
	SayHi()
}

// Personer 定义接口类型
type Personer interface {
	Humaner // 匿名字段，继承Humaner的方法
	Sing(string)
}

// Student ...
type Student struct{}

// SayHi ...
func (s *Student) SayHi() {
	fmt.Println("Student say hi!")
}

// Sing ...
func (s *Student) Sing(str string) {
	fmt.Println("Student is singing:", str)
}

func main() {
	s := &Student{}
	var i Personer
	i = s
	i.SayHi() // 继承了Humaner的方法
	i.Sing("小苹果")
}
