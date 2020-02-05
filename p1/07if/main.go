package main

import "fmt"

func main() {
	// age := 18
	if age := 18; age > 35 {
		fmt.Println("中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("少年")
	}
}
