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
}

func main() {
	s1 := Student{Person{"mike", 'm', 18}, 10010, "bj"}
	s1.name = "qiuqiu"
	s1.sex = 'f'
	fmt.Println(s1.name, s1.sex, s1.id, s1.addr)
}
