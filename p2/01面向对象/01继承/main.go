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

func (p *Person) run() {
	fmt.Println(p.name, "跑起来了")
}

func main() {
	s1 := Student{Person{"mike", 'm', 18}, 10010, "bj"}
	fmt.Printf("%+v\n", s1)
	s2 := Student{Person: Person{age: 18}, id: 10010}
	fmt.Printf("%+v\n", s2)
	s1.run()
}
