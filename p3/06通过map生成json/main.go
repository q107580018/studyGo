package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 创建一个map
	m := make(map[string]interface{})
	m["company"] = "itcast"
	m["subject"] = []string{"Go", "Python", "C++"}
	m["isok"] = true
	m["price"] = 666.666
	// result, err := json.Marshal(m)
	result, err := json.MarshalIndent(m, "", " ")
	if err == nil {
		fmt.Println(string(result))
	} else {
		fmt.Println(err)
	}
}
