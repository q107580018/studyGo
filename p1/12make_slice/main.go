package main

import "fmt"

func main() {
	a1 := make([]int, 1, 10)
	fmt.Printf("a1 = %v, len(a1) = %d, cap(a1) = %d\n", a1, len(a1), cap(a1))

}
