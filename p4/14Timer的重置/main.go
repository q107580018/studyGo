package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	timer := time.NewTimer(3 * time.Second)
	ok := timer.Reset(time.Second)
	fmt.Println("ok=", ok)
	<-timer.C
	fmt.Println("时间到,合计用时：", time.Since(start))

}
