package main

import "fmt"

// 常量
// 定义了常量后不得修改
const pi = 3.1415926

// 批量声明常量
const (
	statusOk = 200
	notFound = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认和上一行一致
const (
	n1 = 100
	n2
	n3
)

// iota
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

const (
	b1 = iota
	b2
	_
	b3
)

// 插队的
const (
	c1 = iota
	c2 = 100
	c3 = iota
	c4
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	// fmt.Println("n1:", n1)
	// fmt.Println("n2:", n2)
	// fmt.Println("n3:", n3)
	// fmt.Println("a1:", a1)
	// fmt.Println("a2:", a2)
	// fmt.Println("a3:", a3)
	// fmt.Println("b1:", b1)
	// fmt.Println("b2:", b2)
	// fmt.Println("b3:", b3)
	// fmt.Println("c1:", c1)
	// fmt.Println("c2:", c2)
	// fmt.Println("c3:", c3)
	// fmt.Println("c4:", c4)
	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	fmt.Println("TB:", TB)
	fmt.Println("PB:", PB)

}
