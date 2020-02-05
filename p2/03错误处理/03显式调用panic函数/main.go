package main

import "fmt"

func a() {
	fmt.Println("aaaaaaaa")
}

func b() {
	fmt.Println("bbbbbbbb")
	panic("this is a panic test")
}

func c() {
	fmt.Println("cccccccc")
}
func main() {
	a()
	b()
	c()
}
