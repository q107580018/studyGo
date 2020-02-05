package main

import "fmt"

// 结构体是值类型
type person struct {
	name   string
	age    int
	gender string
}

func f1(p person) {
	p.gender = "女" // 修改的是副本
}

func f2(p *person) {
	p.gender = "女"
}

func main() {
	var p person
	p.name = "周林"
	p.age = 18
	p.gender = "男"
	f1(p)
	fmt.Println(p.gender) // 男
	f2(&p)
	fmt.Println(p.gender) // 女

}
