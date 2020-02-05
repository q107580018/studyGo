package main

import "fmt"

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret := adder(100)
	fmt.Println(ret(200))
}
