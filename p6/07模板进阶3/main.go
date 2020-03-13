package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Package struct {
	Name     string
	NumFuncs int
	NumVars  int
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	t, err := template.ParseFiles("main.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, &Package{"go-web", 18, 1000})

}

func main() {
	http.HandleFunc("/", Login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
