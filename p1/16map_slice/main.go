package main

import "fmt"

// map和slice组合

func main() {
	// 元素类型为map的切片
	var s1 = make([]map[int]string, 10)
	// 需要对内部的map做初始化
	s1[0] = make(map[int]string, 5)
	s1[0][10] = "邱文"
	fmt.Println(s1)

	// 值为切片类型的map
	var s2 = make(map[string][]int, 10)
	s2["邱文"] = make([]int, 10, 10)
	s2["琦君"] = []int{10, 20, 30}
	fmt.Println(s2)
}
