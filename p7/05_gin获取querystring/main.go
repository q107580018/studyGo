package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器发来的请求携带的query string参数
		query := c.Query("query")
		age := c.Query("age")
		// query := c.DefaultQuery("query", "somebody") // 获取query的值，如果取不到就使用默认值
		// query, ok := c.GetQuery("query") // 获取query的值，如果取不到ok返回false
		// if ok != true {
		// query = "somebody"
		// }
		c.JSON(http.StatusOK, gin.H{
			"query": query,
			"age":   age,
		})
	})
	r.Run(":8080")
}
