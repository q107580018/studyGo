package main

import "fmt"

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

func main() {
	// 要求:f1(f2)
	ret := f3(f2, 100, 200) // 把原来需要传两个参数包装成一个不需要传参的函数
	f1(ret)
}
