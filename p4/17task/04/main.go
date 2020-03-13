package main

import (
	"fmt"
)

// 统计1-200000的数字中，哪些是素数？使用4个协程

var NumChan = make(chan int, 200000)
var PrimeChan = make(chan int, 200000)
var IsQuit = make(chan bool, 4)

func GetPrimeNum() {
	for num := range NumChan {
		flag := true
		for i := 2; i <= num-1; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			PrimeChan <- num
		}
	}
	IsQuit <- true
}
func PutNum() {
	for i := 2; i <= 200000; i++ {
		NumChan <- i
	}
	close(NumChan)
}
func main() {
	ProcsNum := 4
	go PutNum()
	for i := 0; i < ProcsNum; i++ {
		go GetPrimeNum()
	}
	go func() {
		for i := 0; i < ProcsNum; i++ {
			<-IsQuit
		}
		close(PrimeChan)
	}()
	for num := range PrimeChan {
		fmt.Println(num)
	}
}
