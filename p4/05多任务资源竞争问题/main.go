package main

import (
	"fmt"
	"time"
)

func printer(str string) {
	for _, v := range str {
		fmt.Printf("%c", v)
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Print("\n")
}

func main() {
	go printer("hello")
	go printer("world")
	for {
	}
}
