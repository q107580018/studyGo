package model

import "github.com/jinzhu/gorm"

var DB *gorm.DB

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func InitMySQL() (err error) {
	dsn := "root:root123@(127.0.0.1:3306)/bubble?charset=utf8"
	DB, err = gorm.Open("mysql", dsn)
	return
}
