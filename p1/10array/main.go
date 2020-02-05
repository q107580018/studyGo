package main

import "fmt"

// 数组
// 存放元素的容器

func main() {
	var a1 [3]int
	fmt.Printf("%T\n", a1)

	// 数组的初始化
	// 不初始化，默认零值
	fmt.Println(a1) // [0 0 0]
	// 初始化方法1
	a1 = [3]int{3, 4, 5}
	fmt.Println(a1)
	// 初始化方式2
	// 根据数组长度自动推断
	a2 := [...]int{1, 2, 3, 5, 6, 8, 10}
	fmt.Println(a2)
	// 初始化方式3
	// 根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)
	// 数组的遍历
	// 1.根据索引遍历
	cities := [...]string{"北京", "上海", "广州", "深圳"}
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	// 2.for range遍历
	for _, v := range cities {
		fmt.Println(v)
	}

	// 多维数组
	a4 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a4)
	// 多维数组的遍历
	for _, v1 := range a4 {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	// 练习题，求数组所有元素的和
	a5 := [...]int{1, 3, 5, 7, 8}
	var sum int
	for _, v := range a5 {
		sum += v
	}
	fmt.Println(sum)
	// 练习题，找出数组中相加等于8的两个元素的下标
	for i := 0; i < len(a5); i++ {
		for j := i + 1; j < len(a5); j++ {
			if a5[i]+a5[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
