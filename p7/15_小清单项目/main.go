package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:root123@(127.0.0.1:3306)/bubble?charset=utf8"
	DB, err = gorm.Open("mysql", dsn)
	return
}
func main() {
	// 连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	// 模型绑定
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	r.Static("/static", "./statics")
	r.LoadHTMLGlob("./templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// 从请求中获取数据
			var todo Todo
			c.BindJSON(&todo)
			// 存入数据库
			if err := DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"err": err.Error})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 查看所有事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			if err := DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		// 查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, _ := strconv.Atoi(c.Param("id"))
			todo := &Todo{}
			DB.First(todo, id)
			c.BindJSON(todo)
			DB.Save(todo)

		})
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {

		})
	}
	r.Run()
}
