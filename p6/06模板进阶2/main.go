package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	t, err := template.ParseFiles("main.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, r)

}

func main() {
	http.HandleFunc("/", Login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
