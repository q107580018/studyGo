package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

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
	temp := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &temp)
	if err == nil {
		fmt.Printf("%+v\n", temp)

	}
	fmt.Println(strings.Repeat("=", 100))
	for key, value := range temp {
		switch data := value.(type) { // 类型断言
		case string:
			fmt.Printf("temp[%s]的类型为string, 值为%v\n", key, data)
		case bool:
			fmt.Printf("temp[%s]的类型为bool, 值为%v\n", key, data)
		case float64:
			fmt.Printf("temp[%s]的类型为float64, 值为%v\n", key, data)
		case []interface{}: // 重点！此处不是[]string
			fmt.Printf("temp[%s]的类型为[]interface{}, 值为%v\n", key, data)

		}
	}
}
