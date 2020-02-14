package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

// HandleConn ...
func HandleConn(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	fmt.Printf("%s 已连接服务器\n", addr)
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		switch {
		case err == io.EOF:
			fmt.Println(addr, "退出连接")
			return
		case err != nil:
			fmt.Println("connect read错误:", err)
			break
		default:
			StrTrimSpace := strings.TrimSpace(string(buf[:n]))
			if "exit" == StrTrimSpace {
				fmt.Println(addr, "退出连接")
				return
			}
			str := strings.ToUpper(StrTrimSpace)
			fmt.Printf("接收到数据来自[%s]==>:%s\n", addr, StrTrimSpace)
			_, err := conn.Write([]byte(str))
			if err != nil {
				fmt.Println("conn.Write err", err)
			}
		}
	}
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listener错误:", err)
		return
	}
	defer l.Close()
	fmt.Println("start server...")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("connect accept错误:", err)
			return
		}
		go HandleConn(conn)
	}
}
