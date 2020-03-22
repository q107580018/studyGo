package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		// 方法一:使用map
		// data := gin.H{
		// 	"name":    "小王子",
		// 	"message": "hello world",
		// 	"age":     18,
		// }
		// 方法二:使用结构体
		var data struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		data.Name = "小王子"
		data.Message = "hello world"
		data.Age = 19
		c.JSON(http.StatusOK, data)

	})
	r.Run(":8080")
}
