package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 静态文件 html页面上用到的样式文件 .css .js 图片
	r.Static("/xxx", "./statics")  // 当解析模板中出现xxx路径时，则指向statics文件目录
	r.SetFuncMap(template.FuncMap{ // gin框架给模板添加自定义函数
		"safe": func(str string) template.HTML {
			return template.HTML(str) // template.HTML不转义字符串
		},
	})
	// r.LoadHTMLFiles("templates/index.tmpl") // 模板解析
	r.LoadHTMLGlob("templates/**/index.tmpl") //正则匹配模板解析(**表示文件夹，*表示文件)
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ // 模板渲染
			"title": "小王子",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{ // 模板渲染
			"title": " <a href='https://www.baidu.com'>百度一下</a> ", // 一般情况下a标签会被转义成字符串
		})
	})
	r.Run(":8080") // 启动server
}
