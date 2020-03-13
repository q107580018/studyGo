package main

import (
	"fmt"
	"sync"
)

// 需求：计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中，最后显示出来
// 要求用goruntime完成。
var isQuit = make(chan bool)
var Lock = new(sync.Mutex)

func main() {
	m := make(map[uint64]uint64)
	for i := uint64(1); i <= 200; i++ {
		go Factorial(i, m)
	}
	for i := uint64(1); i <= 200; i++ {
		<-isQuit
	}
	for i := uint64(1); i <= 200; i++ {
		fmt.Printf("%v:%v\n", i, m[i])
	}
}

func Factorial(num uint64, m map[uint64]uint64) {
	sum := uint64(1)
	for i := uint64(1); i <= num; i++ {
		sum *= i
	}
	Lock.Lock()
	m[num] = sum
	Lock.Unlock()
	isQuit <- true
}
