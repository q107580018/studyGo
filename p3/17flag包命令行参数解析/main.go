package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		user string
		pwd  string
		port int
		host string
	)
	flag.StringVar(&user, "u", "", "用户名")
	flag.StringVar(&pwd, "pwd", "", "密码")
	flag.IntVar(&port, "p", 8080, "端口号")
	flag.StringVar(&host, "h", "localhost", "主机名")
	flag.Parse() // 重要！转换必须调用改函数
	fmt.Printf("user:%s pwd:%s host:%s port:%d\n", user, pwd, host, port)

}
