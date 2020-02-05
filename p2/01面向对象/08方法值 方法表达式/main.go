package main

import "fmt"

// Person 人
type Person struct {
	name string
	sex  byte
	age  int
}

// SetNameValue 接收者为值变量，为拷贝
func (p Person) SetNameValue() {
	fmt.Println("SetNameValue")
}

// SetNamePointer 接收者为指针变量，引用传递
func (p *Person) SetNamePointer() {
	fmt.Println("SetNamePointer")

}
func main() {
	p := Person{"qiu", 'f', 18}
	// 传统调用
	p.SetNameValue()
	p.SetNamePointer()
	// 方法值
	f1 := p.SetNameValue // 无需再传递接受者，隐藏了接收者
	f1()
	f2 := p.SetNamePointer
	f2()
	// 方法表达式
	f3 := (*Person).SetNamePointer
	f3(&p) // 显式传递接受者
	f4 := (Person).SetNameValue
	f4(p)

}
