package main

import (
	"encoding/json"
	"fmt"
)

// IT 变量必须首字母大写才能被json包读取到
type IT struct {
	Company string   `json:"company"` // 二次编码转为首字母小写
	Subject []string `json:"-"`       // "-"代表此字段不会输出到屏幕
	IsOk    bool     `json:",string"` // 转换为字符串
	Price   float64
}

func main() {
	s := IT{"itcast", []string{"Go", "Python", "C++"}, true, 66.66}
	// buf, err := json.Marshal(s) // 转换为json字符串
	buf, err := json.MarshalIndent(s, "", " ") // 转换为json字符串，格式化输出
	if err == nil {
		fmt.Println(string(buf))
	}
}
