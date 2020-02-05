package main

import (
	"encoding/json"
	"fmt"
)

// IT ...
type IT struct {
	Company string   `json:"company"` // 二次编码转为首字母小写
	Subject []string `json:"subject"`
	IsOk    bool     `json:"isok"`
	Price   float64  `json:"price"`
}

func main() {
	jsonStr := `
	{
		"company": "itcast",
		"isok": true,
		"price": 666.666,
		"subject": [
		 "Go",
		 "Python",
		 "C++"
		]
	   }
	`
	var temp IT
	err := json.Unmarshal([]byte(jsonStr), &temp) // 第二个参数需要传地址
	if err == nil {
		fmt.Printf("%+v\n", temp)
	}

}
