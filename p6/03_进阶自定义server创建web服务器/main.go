package main

import (
	"log"
	"net/http"
	"time"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world" + r.URL.String()))
}

type MyHandler struct{}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// time.Sleep(3 * time.Second) // 睡眠3秒，测试写入超时
	w.Write([]byte("bye bye"))
}
func main() {
	server := &http.Server{ // 自定义server
		Addr:         ":8080",
		WriteTimeout: 2 * time.Second, // 写入超时
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", SayHello)    // 方式一(如果为根路由，会匹配所有未匹配的路由)
	mux.Handle("/bye", &MyHandler{}) // 方式二
	server.Handler = mux             // 将mux集成到server中
	log.Println("starting server v2 ....")
	log.Fatal(server.ListenAndServe()) // 利用自定义server启动
}
