package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/base.tmpl", "./template/index.tmpl")
	if err != nil {
		log.Fatal("template parse failed, err:", err)
	}
	msg := "小王子"
	t.ExecuteTemplate(w, "index.tmpl", msg)

}
func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/base.tmpl", "./template/home.tmpl")
	if err != nil {
		log.Fatal("template parse failed, err:", err)
	}
	msg := "小王子"
	t.ExecuteTemplate(w, "home.tmpl", msg)

}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}
