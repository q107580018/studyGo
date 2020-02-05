package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a task")
		time.Sleep(time.Second)
	}
}
func main() {
	go newTask()
	// 要保证主程序在子程序运行中不能结束，主程序结束的话所有子程序也会结束
	for {
		fmt.Println("this is a main")
		time.Sleep(time.Second)
	}
}
