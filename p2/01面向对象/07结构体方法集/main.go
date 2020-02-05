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
	// p.name = name
	fmt.Println("SetNameValue")
}

// SetNamePointer 接收者为指针变量，引用传递
func (p *Person) SetNamePointer() {
	// p.name = name
	fmt.Println("SetNamePointer")

}
func main() {
	p1 := Person{"小球", 'm', 19}
	p1.SetNamePointer()
	p1.SetNameValue()
	p2 := &Person{"小安", 'f', 17}
	p2.SetNamePointer()
	p2.SetNameValue()
}
