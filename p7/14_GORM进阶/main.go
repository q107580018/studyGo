package main

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 例子:定义类型
type User struct {
	gorm.Model   // 内嵌gorm.Model
	Name         string
	Age          sql.NullInt64 // 零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号 唯一且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置num为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

func main() {
	// 连接MySQL数据库
	db, err := gorm.Open("mysql", "root:root123@(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true) /// 禁用改名为复数
	db.AutoMigrate(&User{})
	db.Table("qiuwen").AutoMigrate(&User{}) // 使用User结构体创建名叫qiuwen的表
}
