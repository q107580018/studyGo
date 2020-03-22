package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo --> 数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:root123@(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 创建表 自动迁移
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	//  u1 := UserInfo{1, "qiuwen", "男", "蛙泳"}
	//	db.Create(&u1)
	// 查询
	var u UserInfo
	db.First(&u) // 查询表中第一条数据
	fmt.Printf("%#v\n", u)
	// 更新
	db.Model(&u).Update("Hobby", "双色球")
	// 删除
	db.Delete(&u)
}
