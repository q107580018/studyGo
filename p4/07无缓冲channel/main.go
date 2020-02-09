package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("子协程:i= %d\n", i)
			ch <- i // 往通道里写内容，没被读取前会阻塞
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-ch // 读管道中的内容，没有内容前，阻塞
		fmt.Println("num =", num)
	}
}
