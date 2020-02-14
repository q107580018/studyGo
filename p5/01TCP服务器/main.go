package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	fmt.Println(string(buf[:n]))
}
