package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

// Client 客户
type Client struct {
	C    chan string
	Name string
	Addr string
}

// Message 广播发送数据用
var Message = make(chan string)

// OnlineMap 用户数据表
var OnlineMap = make(map[string]*Client)

// MakeMsg 发送信息的格式
func MakeMsg(name, addr, str string) (msg string) {
	msg = fmt.Sprintf("[%s]%s:%s", addr, name, str)
	return
}

// SendMsgToClient 发送信息给客户端
func SendMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

// HandleConnect 处理接入的客户端
func HandleConnect(conn net.Conn) {
	defer conn.Close()
	isQuit := make(chan bool)  // 对方是否主动退出
	HasDate := make(chan bool) // HasDate 判断用户是否有发送数据，用于超时处理
	cliAddr := conn.RemoteAddr().String()
	fmt.Printf("%s:enter\n", cliAddr)
	cli := Client{make(chan string), cliAddr, cliAddr}
	OnlineMap[cliAddr] = &cli
	Message <- MakeMsg(cli.Name, cli.Addr, "login....")
	go SendMsgToClient(cli, conn) // SendMsgToClient 发送信息给客户端
	// 新建一个协程，接收客户发来的信息
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if n == 0 && err == io.EOF {
				isQuit <- true
				fmt.Println(MakeMsg(cli.Name, cliAddr, "login out"))
				return
			}
			msg := strings.TrimSpace(string(buf[:n]))
			switch {
			case msg == "who":
				msg1 := "user list:\n"
				for _, temp := range OnlineMap {
					msg1 += temp.Addr + ":" + temp.Name + "\n"
				}
				conn.Write([]byte(msg1))
			case strings.HasPrefix(msg, "rename|"):
				name := msg[7:]
				if name == "" {
					conn.Write([]byte("Name cannot be empty!"))
					break
				}
				cli.Name = name
				conn.Write([]byte("rename succesce!"))

			default:
				Message <- MakeMsg(cli.Name, cli.Addr, msg)
			}
			HasDate <- true
		}
	}()
	for {
		select {
		case <-isQuit: // 用户退出
			delete(OnlineMap, cliAddr) // 当前用户从map移除
			Message <- MakeMsg(cli.Name, cliAddr, "login out")
			return
		case <-HasDate:
			// 如果有数据，则会进入此处
		case <-time.After(time.Second * 60): // 如果超过时间还没有数据，则超时退出
			delete(OnlineMap, cliAddr) // 当前用户从map移除
			Message <- MakeMsg(cli.Name, cliAddr, "time out leave")
			fmt.Println(MakeMsg(cli.Name, cliAddr, "time out leave"))
			return
		}
	}
}

// Broadcast 只要有消息来了，遍历map，给每个成员广播消息
func Broadcast() {
	for {
		msg := <-Message // 没有消息前阻塞
		for _, cli := range OnlineMap {
			cli.C <- msg
		}
	}

}

func main() {
	fmt.Println("Server start...")
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer l.Close()
	go Broadcast() // Broadcast 只要有消息来了，遍历map，给每个成员广播消息
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("l.Accept err:", err)
			continue
		}
		go HandleConnect(conn)
	}
}
