package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
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

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.gtpl")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm() // 解析参数，默认不解析
		// fmt.Println("username:", r.Form["username"]) // 不推荐
		fmt.Println("username:", r.Form.Get("username:")) // 如果多个值，此方法只能获取第一个值
		fmt.Println("password:", r.Form.Get("password"))
	}
}

func main() {
	http.HandleFunc("/", SayHello)
	http.HandleFunc("/login", Login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
