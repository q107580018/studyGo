package main

import "fmt"

func main() {
	// 基本格式
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 变种1
	// i := 0
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 变种2
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 无限循环
	// for {
	// 	fmt.Println("hello")
	// }

	// for range循环
	// s := "hello你好"
	// for i, v := range s {
	// 	fmt.Printf("%d: %c\n", i, v)
	// }

	// 九九乘法表
	for i := 1; i < 10; i++ {
		for n := 1; n <= i; n++ {
			fmt.Printf("%d * %d = %d\t", n, i, n*i)
		}
		fmt.Println()
	}
}
