package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains 查看是否包含
	fmt.Println(strings.Contains("hellomynameisqiu", "name"))
	fmt.Println("--------------------------------")
	// Join 组合
	s := strings.Join([]string{"abc", "def", "ghi", "jkl"}, ",")
	fmt.Println(s)
	fmt.Println("--------------------------------")

	// Index 返回索引位置
	fmt.Println(strings.Index("abcdehello", "hello"))
	fmt.Println(strings.Index("abcdehello", "bye")) // 找不到返回-1
	fmt.Println("--------------------------------")

	// Repeat 重复
	buf := strings.Repeat("go", 3)
	fmt.Println(buf)
	fmt.Println("--------------------------------")

	// Split 分割
	s2 := strings.Split("jfekfjsfjklaesjfeknrsnfrewsfjeafsfjkeneqw", "s")
	fmt.Println(s2)
	fmt.Println("--------------------------------")

	// Trim 去掉两头的字符
	s3 := strings.Trim("    are u ok?      ", " ")
	fmt.Println(s3)
	fmt.Println("--------------------------------")

	// Fields 以空格分割字符
	s4 := strings.Fields("    are u ok?      ")
	for i, v := range s4 {
		fmt.Printf("%d=%v\n", i, v)
	}
	fmt.Println("--------------------------------")

}
