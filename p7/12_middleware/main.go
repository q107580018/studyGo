package main

// 中间件
import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index handler")
	c.JSON(http.StatusOK, gin.H{"msg": "index"})
}

// 定义一个中间件m1:统计耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	// 计时
	start := time.Now()
	c.Next() //调用后续函数(indexHandler)
	cost := time.Since(start)
	fmt.Printf("耗时:%v\n", cost)
	fmt.Println("m1 out...")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Set("name", "qiuwen")
	// c.Abort() //阻止调用后续的处理函数
	fmt.Println("m2 out...")
}

func main() {
	r := gin.Default() //默认使用了logger和recover两个中间件
	// 添加中间件方法一：
	//r.GET("/index", m1, indexHandler)
	//r.GET("/shop",m1, func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{
	//		"msg":"shop",
	//	})
	//})
	//r.GET("/user", func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{
	//		"msg":"user",
	//	})
	//})
	// 添加中间件方法二:
	r.Use(m1, m2) // 全局注册中间件函数m1
	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		name, ok := c.Get("name")
		if !ok {
			name = "匿名用户"
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": name,
		})
	})
	_ = r.Run(":8080")
}
