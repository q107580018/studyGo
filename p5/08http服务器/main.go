package main

import (
	"fmt"
	"net/http"
)

// HelloServer hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello world!\n"))
}
func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
