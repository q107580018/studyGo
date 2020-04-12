package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("user", "qw", 600, "/", "127.0.0.1", false, true)
		c.String(200, "Login successful")

	})
	r.Use(CheckCookie)
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "home",
		})
	})
	_ = r.Run()
}

func CheckCookie(c *gin.Context) {
	_, err := c.Cookie("user")
	if err != nil {
		c.JSON(200, gin.H{
			"error": "StatusUnauthoriezd",
		})
		c.Abort()
	} else {
		c.Next()
	}

}
