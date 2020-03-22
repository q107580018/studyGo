package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "https://www.baidu.com")
	})

	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b" // 将请求路径改为/b
		r.HandleContext(c)        // 继续后续的处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})
	_ = r.Run(":8080")
}
