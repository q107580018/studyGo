package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		fmt.Println("http.get err:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("status:", resp.Status)
	fmt.Println("StatusCode:", resp.StatusCode)
	fmt.Println("Header:", resp.Header)
	buf := make([]byte, 2048*2)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.body.read err:", err)
			break
		}
		fmt.Println(string(buf[:n]))
	}

}
