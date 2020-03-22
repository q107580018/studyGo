package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Any("/index", func(c *gin.Context) { // Any 匹配所有请求
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
		}
	})
	// NoRoute处理没有路由的页面
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Not found"})
	})

	// 路由组

	// 视频的首页和详情页
	// 把公用的前缀提取出来，创建一个路由组
	routeGroup := r.Group("/video")
	// 以下都为路由的子项，写后续的路径就可以
	{ // 加个代码块，更加清晰
		routeGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/index", "status": "OK"})
		})
		routeGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/xx", "status": "OK"})
		})
		routeGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/oo", "status": "OK"})
		})
	}
	r.Run(":8080")
}
