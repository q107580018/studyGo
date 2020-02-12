package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个定时器，倒计时3秒，3秒后往timer.C通道(<-chan Time)写内容,
	// 内容为当前时间，只能使用一次
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("当前时间:", time.Now().Format("2006年1月2日15:04:05"))
	t := <-timer.C // 没数据前会阻塞
	fmt.Println("t:", t.Format("2006年1月2日15:04:05"))

}
