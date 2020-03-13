package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("go-web").Parse(`<html>
<head>
	<title>首页</title>
</head>

<body>
	<a href="/login">登陆</a>
</body>
</html>
	`)
	if err != nil {
		fmt.Println("temp.Parse failed,", err)
		return
	}
	tmpl.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm() // 解析参数，默认不解析
		// fmt.Println("username:", r.Form["username"]) // 不推荐
		fmt.Println("username:", r.Form.Get("username")) // 如果多个值，此方法只能获取第一个值
		fmt.Println("password:", r.Form.Get("password"))
		t, _ := template.New("logi").Parse(r.Form.Get("username") + "你好")
		t.Execute(w, nil)
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
