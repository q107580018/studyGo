package main

import "fmt"

func main() {
	// 10进制
	i1 := 107
	fmt.Printf("i1:%d\n", i1)
	fmt.Printf("二进制:%b\n", i1)
	fmt.Printf("八进制:%o\n", i1)
	fmt.Printf("十六进制:%x\n", i1)
	// 八进制
	i2 := 066
	fmt.Printf("i2:%d\n", i2)
	// 十六进制
	i3 := 0x123
	fmt.Printf("i3:%d\n", i3)
	// 查看变量类型
	fmt.Printf("%T\n", i3)
	// 声明int8类型
	// var i4 int8 = 9
	i4 := int8(9) // 明确指定int8类型，否则就是默认的int类型
	fmt.Printf("%T\n", i4)
}
