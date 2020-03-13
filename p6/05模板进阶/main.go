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

func SayHello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("go-web").Parse(`
Package name: {{.Name}}
Number of funcctions: {{.NumFuncs}}
Number of variables: {{.NumVars}}
	`)
	if err != nil {
		fmt.Println("temp.Parse failed,", err)
		return
	}
	tmpl.Execute(w, &Package{"go-web", 12, 30})
}

func main() {
	http.HandleFunc("/", SayHello)
	fmt.Println("server start ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
