package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

func downloadFile(conn net.Conn, path string) {
	_, fileName := filepath.Split(path)
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer f.Close()
	_, err = conn.Write([]byte(path))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
	buf := make([]byte, 1024*1024)
	for {
		// 读取下载文件
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件下载完毕")
			} else {
				fmt.Println("conn.Read err:", err)
			}
			return
		}
		// 写入文件
		f.Write(buf[:n])
	}
}

func main() {
	list := os.Args
	if len(list) != 3 {
		fmt.Println("usage: xxx 127.0.0.1:8080 1.jpg")
		return
	}
	ip := list[1]
	path := list[2]
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	downloadFile(conn, path)
}
