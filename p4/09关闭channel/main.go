package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // 关闭channel后无法再发送数据
	}()

	for {
		if num, ok := <-ch; ok == true {
			fmt.Println("num=", num)
		} else { // channel关闭
			break
		}
	}
}
