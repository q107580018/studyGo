package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})
	r.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})
	r.PUT("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})
	r.DELETE("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("gin run err:", err)
	}
}
