package main

import "fmt"

// 声明变量
// var name string
// var age int
// var isOk bool

// 批量声明
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func main() {
	name = "李想"
	age = 18
	isOk = true
	fmt.Printf("name:%s, age:%d, isOk:%t\n", name, age, isOk)

	// 声明变量同时赋值
	var s1 string = "王思聪"
	fmt.Println(s1)
	// 类型推导（根据值判断该变量是什么类型）
	var s2 = 20
	fmt.Println(s2)
	// 简短变量声明, 只能在函数里使用
	s3 := "哈哈哈"
	fmt.Println(s3)

}
