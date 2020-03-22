package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct { // 定义一个接收参数的结构体
	UserName string `form:"username"` // 用到了反射，所以元素要以大写开头
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBind(&user)
		if err != nil {
			c.String(http.StatusBadRequest, "获取信息失败")
		} else {
			userName := user.UserName
			Password := user.Password
			c.JSON(http.StatusOK, gin.H{
				"user_name": userName,
				"password":  Password,
			})
		}
	})
	r.POST("/form", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBind(&user)
		if err != nil {
			c.String(http.StatusBadRequest, "获取信息失败")
		} else {
			userName := user.UserName
			Password := user.Password
			c.JSON(http.StatusOK, gin.H{
				"user_name": userName,
				"password":  Password,
			})
		}
	})
	r.Run(":8080")
}
