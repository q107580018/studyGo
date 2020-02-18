package main

import "fmt"

type MyFunc func(int)

func (m MyFunc) Print(a int) {
	m(a)
}

func aFun(a int) {
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
