package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func printer(str string) {
	for _, v := range str {
		fmt.Printf("%c", v)
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Print("\n")
}

func person1() {
	printer("hello")
	ch <- 666 // 给通道写数据
}

func person2() {
	<-ch // 等待接收数据，如果通道没有数据就会阻塞
	printer("world")
}

func main() {
	go person1()
	go person2()
	for {
	}
}
