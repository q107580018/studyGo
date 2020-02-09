package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3) // 设置3个空间
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i // 往通道里写内容，通道满了后会阻塞
			fmt.Printf("子协程:i= %d\n", i)
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		num := <-ch // 读管道中的内容，没有内容时，阻塞
		fmt.Println("num =", num)
	}
}
