package main

import (
	"log"
	"net/http"
)

// SayHello ...
func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world" + r.URL.String()))
}

// MyHandler ...
type MyHandler struct{}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", SayHello)    // 方式一(如果为根路由，会匹配所有未匹配的路由)
	mux.Handle("/bye", &MyHandler{}) // 方式二
	log.Println("starting server v1 ....")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
