package main

import "fmt"

type MyFunc func(int) // MyFunc函数实现了print方法,兼容MyInterface

func (m MyFunc) Print(a int) {
	m(a)
}

func aFun(a int) { // 原函数没有实现print方法，不兼容MyInterface
	fmt.Println(a)
}

type MyInterface interface {
	Print(int)
}

func bFun(m MyInterface, a int) {
	m.Print(a)
}

func main() {
	bFun(MyFunc(aFun), 5)
}
