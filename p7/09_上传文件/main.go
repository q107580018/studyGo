package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		// 从请求中读取文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {

			// 将读取到的文件保存到本地
			dstPath := "./" + file.Filename
			fmt.Println(dstPath)
			c.SaveUploadedFile(file, dstPath)
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}

	})
	r.Run(":8080")
}
