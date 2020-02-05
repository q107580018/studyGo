// 打印出100-999中所有的“水仙花数”

package main

import (
	"fmt"
)

func cube(a int) int {
	return a * a * a
}
func main() {
	var Narcissistic = make([]int, 0, 1000)
	for i := 100; i < 1000; i++ {
		a1 := i % 10
		a2 := i / 10 % 10
		a3 := i / 100 % 10
		if cube(a1)+cube(a2)+cube(a3) == i {
			Narcissistic = append(Narcissistic, i)
		}
	}
	fmt.Println(Narcissistic)

}
