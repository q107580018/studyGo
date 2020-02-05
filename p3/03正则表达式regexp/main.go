package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 例1===================================
	buf := "abc azc a7c aac 888 a99 tac"

	// 解释规则
	reg := regexp.MustCompile(`a.c`)
	// 根据规则提取
	s := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println(s)

	// 例二===================================
	// 找出数字
	buf = "41.36 1.2 fjaekaw 3.83 36.1 5. 9.99 kerjekw .1231"
	reg = regexp.MustCompile(`\d+\.\d+`)
	s2 := reg.FindAllString(buf, -1)
	fmt.Println(s2)
}
