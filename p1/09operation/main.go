package main

import "fmt"

// 运算符
func main() {
	var (
		a = 5
		b = 2
	)
	// 算数运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句 ==> a = a + 1
	b-- //单独的语句 ==> b = b - 1
	fmt.Println(a)
	fmt.Println(b)

	// 关系运算符
	fmt.Println(a == b) // Go语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a < b)

	// 逻辑运算符
	age := 20
	if age > 18 && age < 60 { // 如果年龄大于18岁 并且 小于60岁
		fmt.Println("苦逼上班")
	} else {
		fmt.Println("不用上班")
	}

	if age < 18 || age > 60 { // 如果年龄小于18岁 或者 大于60岁
		fmt.Println("不用上班")
	} else {
		fmt.Println("苦逼上班")
	}

	// !取反
	flag := false
	fmt.Println(flag)
	fmt.Println(!flag)

	// 位运算：针对二进制数
	// 5的二进制是：101
	// 2的二进制是： 10
	// &:按位与（二者都为1才为1）
	fmt.Println(5 & 2) // 000
	// |:按位或 (二者有一个为1就为1)
	fmt.Println(5 | 2) // 111
	// ^:按位异或(二者不一样则为1)
	fmt.Println(5 ^ 2) // 111
	// <<:二进制左移位数
	fmt.Println(5 << 1)  //1010(二)==>10(十)
	fmt.Println(1 << 10) //10000000000(二) ==>1024(十)
	// >>:二进制右移位数
	fmt.Println(5 >> 1) // 10(二)==>2(十)

}
