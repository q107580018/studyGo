package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	score := r.FormValue("score")
	num, _ := strconv.Atoi(score)
	t, err := template.ParseFiles("main.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, num)
}
func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("login.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, r)
}
func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/login", Login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
