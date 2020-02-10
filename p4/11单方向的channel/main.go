package main

import "fmt"

func producer(out chan<- int) { // 此通道只能写，不能读
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func customer(in <-chan int) { // 此通道只能读，不能写
	for i := range in {
		fmt.Println("num=", i)
	}
}

func main() {
	// 创建一个双向通道
	ch := make(chan int)
	go producer(ch) // channel是引用传递
	customer(ch)
}
