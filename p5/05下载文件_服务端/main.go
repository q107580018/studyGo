package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

// uploadFile 向客户端发送文件
func uploadFile(conn net.Conn, addr, path string) {
	defer conn.Close()
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("未找到客户需要的文件")
		return
	}
	defer f.Close()
	buf := make([]byte, 1024*1024)
	// 读取内容
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err:", err)
			}
			return
		}
		// 发送内容
		conn.Write(buf[:n])
	}

}
func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.listen err:", err)
		return
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("L.accept err:", err)
			return
		}
		defer conn.Close()
		addr := conn.RemoteAddr().String()
		fmt.Println(addr, "已连接本服务器....")
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		path := string(buf[:n])
		go uploadFile(conn, addr, path)
	}
}
