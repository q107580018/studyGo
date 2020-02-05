package main

import "fmt"

// 构造函数

type person struct {
	name string
	age  int
}

// 构造函数:约定成俗用new开头
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("qiu", 20)
	fmt.Println(p1)
}
