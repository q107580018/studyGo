package main

import "fmt"

func a() {
	fmt.Println("aaaaaaaa")
}

func b(x int) {
	l := [3]int{1, 2, 3}
	fmt.Println(l[x])
}

func c() {
	fmt.Println("cccccccc")
}
func main() {
	a()
	b(3)
	c()
}
