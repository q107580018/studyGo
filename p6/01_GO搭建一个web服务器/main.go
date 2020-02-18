package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认不解析
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
