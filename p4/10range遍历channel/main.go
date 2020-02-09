package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 500)
		}
		close(ch) // 关闭channel后无法再发送数据
	}()

	for num := range ch { // 如果channel关闭，自动跳出循环
		fmt.Println("num=", num)
	}
}
