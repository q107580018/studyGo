package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// ReadMsg 读取服务器发送过来的内容
func ReadMsg(conn net.Conn) {
	buf := make([]byte, 512)
	for {
		n, err := conn.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("退出连接")
			return
		}
		fmt.Println("接收到数据:", string(buf[:n]))
	}
}

// WriteMsg 向服务器发送内容
func WriteMsg(conn net.Conn) {
	buf := make([]byte, 512)
	for {
		fmt.Print("请输入发送的内容:")
		n, err := os.Stdin.Read(buf) // 从键盘读取内容
		if err != nil {
			fmt.Println("os.Stdin err", err)
			break
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("conn.Write err", err)
		}
		time.Sleep(1e8)
	}
}
func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Dial err", err)
		return
	}
	// 调用完毕后自动关闭
	defer conn.Close()
	go WriteMsg(conn)
	ReadMsg(conn)
}
