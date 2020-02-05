package main

import "fmt"

// Person 人
type Person struct {
	name string
	sex  byte
	age  int
}

// SetNameValue 接收者为值变量，为拷贝
func (p Person) SetNameValue(name string) {
	p.name = name
}

// SetNamePointer 接收者为指针变量，引用传递
func (p *Person) SetNamePointer(name string) {
	p.name = name
}
func main() {
	p1 := Person{"小球", 'm', 19}
	fmt.Println("没改前名字为：", p1.name)
	p1.SetNameValue("小猫")
	fmt.Println("使用值接收者后名字为：", p1.name)
	p1.SetNamePointer("小狗")
	fmt.Println("使用指针接收者后名字为：", p1.name)

}
