package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/q107580018/studyGo/p7/15_small_list/controller"
	"github.com/q107580018/studyGo/p7/15_small_list/model"
)

func main() {
	// 连接数据库
	err := model.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer model.DB.Close()
	// 模型绑定
	model.DB.AutoMigrate(&model.Todo{})

	r := gin.Default()
	// 静态资源路径
	r.Static("/static", "./statics")
	// 模板文件路径
	r.LoadHTMLGlob("./templates/*")
	r.GET("/index", controller.IndexHandler)
	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateATodo)
		// 查看所有事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	r.Run()
}
