package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("定时器2秒时间到")

	}()
	timer.Stop() // timer关闭， 子协程一直阻塞
	for {
	}
}
