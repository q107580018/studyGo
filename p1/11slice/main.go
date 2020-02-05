package main

import "fmt"

// 切片slice
func main() {
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)

	s1 = []int{1, 3, 5, 7, 9}
	s2 = []string{"北京", "上海", "广州", "深圳"}
	fmt.Println(s1)
	fmt.Println(s2)
	// 长度和容量
	fmt.Printf("len(s1):%d, cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d, cap(s2):%d\n", len(s2), cap(s2))
	// 由数组得到的切片
	s3 := [...]int{1, 5, 7, 4, 13, 53}
	s4 := s3[0:4]
	fmt.Printf("%T, %d\n", s4, s4)
	fmt.Printf("%p\n", s4)

}
