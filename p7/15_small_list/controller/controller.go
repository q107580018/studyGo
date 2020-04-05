package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/q107580018/studyGo/p7/15_small_list/model"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func CreateATodo(c *gin.Context) {
	// 从请求中获取数据
	var todo model.Todo
	c.BindJSON(&todo)
	// 存入数据库
	if err := model.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
func GetTodoList(c *gin.Context) {
	var todoList []model.Todo
	if err := model.DB.Find(&todoList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id := c.Param("id")
	var todo model.Todo
	model.DB.First(&todo, id)
	c.BindJSON(&todo)
	if err := model.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func DeleteATodo(c *gin.Context) {
	var todo model.Todo
	id := c.Param("id")
	model.DB.First(&todo, id)
	if err := model.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "删除成功"})
	}
}
