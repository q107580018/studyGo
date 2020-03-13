package main

import (
	"fmt"
	"time"
)

// 开启一个writeData协程向管道写入50个整数，一个readData协程，从管道读取数据
// 主线程需要等两个协程都完成工作后才能退出。
var ch = make(chan int)
var isOk = make(chan bool)

func writeData() {
	for i := 0; i < 50; i++ {
		ch <- i
		time.Sleep(1e8)
	}
	close(ch)
	isOk <- true
}

func readData() {
	for num := range ch {
		fmt.Println(num)
	}
	isOk <- true
}
func main() {
	go writeData()
	go readData()
	<-isOk
	<-isOk
}
