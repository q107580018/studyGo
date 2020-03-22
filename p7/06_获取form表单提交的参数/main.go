package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		// 方法一:
		// userName := c.PostForm("user")
		// pwd := c.PostForm("pwd")
		// 方法二:
		// userName := c.DefaultPostForm("user", "somebody")
		// pwd := c.DefaultPostForm("pwd", "****")
		// 方法三:
		userName, ok := c.GetPostForm("user")
		if !ok {
			return
		}
		pwd, ok := c.GetPostForm("pwd")
		if !ok {
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name": userName,
			"Pwd":  pwd,
		})
	})
	r.Run(":8080")
}
