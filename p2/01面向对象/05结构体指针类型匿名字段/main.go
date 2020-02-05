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
	*Person // 指针类型
	id      int
	addr    string
}

func main() {
	s1 := Student{&Person{"mike", 'm', 18}, 10086, "bj"}
	fmt.Printf("%+v\n", s1)
	fmt.Println(*s1.Person)
}
