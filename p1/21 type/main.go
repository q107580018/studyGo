package main

import "fmt"

type newInt int  // 自定义类型
type myInt = int // 类型别名

func main() {
	var m newInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)
}
