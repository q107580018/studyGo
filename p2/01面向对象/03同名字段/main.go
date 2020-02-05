package main

import "fmt"

// Person 人
type Person struct {
	name string
	sex  byte
	age  int
}

// Student 学生
type Student struct {
	Person // 匿名字段
	id     int
	addr   string
	name   string //同名字段
}

func main() {
	var s Student
	s.name = "mike" // 操作的是Student的name
	s.id = 10010
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"
	fmt.Printf("%+v\n", s)
	// 显式调用
	s.Person.name = "bill" //Peroson的name
	fmt.Printf("%+v\n", s)

}
