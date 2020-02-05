package main

import "fmt"

func main() {
	var i interface{}
	i = 1
	fmt.Println(i)
	i = "Hello"
	fmt.Println(i)
	i = 'm'
	fmt.Println(i)

}
