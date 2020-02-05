package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	// FormatBool bool转string
	str = strconv.FormatBool(false)
	fmt.Println(str)
	// FormatFloeat float转string
	str = strconv.FormatFloat(13.257, 'f', -1, 64)
	fmt.Println(str)
	// Itoa 整型转字符串
	str = strconv.Itoa(6666)
	fmt.Println(str)
	// FormatInt 返回i的base进制的字符串
	str = strconv.FormatInt(1234, 2) // 2 为二进制
	fmt.Println(str)

	// 字符串转其他类型
	flag, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(flag)
	} else {
		fmt.Println(err)
	}
	// 字符串转整型
	i, _ := strconv.Atoi("12311")
	fmt.Printf("type:%T, value:%d\n", i, i)

}
