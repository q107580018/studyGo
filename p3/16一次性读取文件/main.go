package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content, err := ioutil.ReadFile("demo.txt")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("%v\n", string(content))
}
