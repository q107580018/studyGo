package main

import "fmt"

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1
	var a3 = []int{0, 0, 0}

	copy(a3, a1)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	a1[0] = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

}
