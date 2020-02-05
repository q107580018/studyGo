package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("aaaaaa")
	runtime.Goexit() // 终止所在协程
	fmt.Println("bbbbbbbb")
}

func main() {
	// 创建新的协程
	go func() {
		fmt.Println("cccccccc")
		// 调用别的函数
		test()
		fmt.Println("ddddddddd")
	}()
	// 死循环，不让主协程结束
	for {
	}
}
