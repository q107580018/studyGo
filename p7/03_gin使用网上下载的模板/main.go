package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.Static("/css", "./statics/css")
	r.Static("/img", "./statics/img")
	r.Static("/js", "./statics/js")
	r.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/index-carousel.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index-carousel.html", nil)
	})
	r.GET("/ui-elements.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ui-elements.html", nil)
	})
	r.Run(":8080")
}
