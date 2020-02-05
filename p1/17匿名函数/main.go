package main

import "fmt"

func main() {

	v := func(x int) int {
		return x * x
	}
	fmt.Println(v(10))
}
