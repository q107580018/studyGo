package main

import (
	"fmt"
	// "os"
)

func main() {
	// os.Stdout.Close() // 标准输出设备，关闭后无法输出
	fmt.Println("are u ok?")

	// os.Stdin.Close() // 标准输入设备，关闭后无法输入
	var a int
	fmt.Print("请输入:")
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
