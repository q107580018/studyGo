package main

import "fmt"

func main() {
	a := make(map[string]int, 10)
	a["qiuwen"] = 29
	a["qijun"] = 28
	fmt.Println(a)
	fmt.Println(a["qiuwen"])
	for k, v := range a {
		fmt.Println(k, v)
	}
}
