package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var p person
	p.name = "周林"
	p.age = 28
	p.gender = "男"
	p.hobby = []string{"篮球", "足球", "台球"}
	fmt.Println(p)
	fmt.Printf("%#v\n", p)
	fmt.Println(p.name)

	// 匿名结构体:多用于临时场景
	var s struct {
		name string
		age  int
	}
	s.name = "球"
	s.age = 32
	fmt.Println(s)
	fmt.Printf("type:%T value:%v\n", s, s)
}
