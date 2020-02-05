package main

import "fmt"

// Mystr ...
type Mystr string

// Person 人
type Person struct {
	name string
	sex  byte
	age  int
}

// Student 学生
type Student struct {
	Person // 匿名字段
	int
	Mystr
}

func main() {
	s1 := Student{Person{"qiu", 'm', 18}, 10086, "Hello"}
	fmt.Printf("%+v]\n", s1)
	s1.int = 10010
	s1.Mystr = "World!"
	fmt.Printf("%+v]\n", s1)

}
