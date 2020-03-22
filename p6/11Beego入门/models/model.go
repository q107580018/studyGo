package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string `orm:"unique"` // unique唯一
	Pwd  string
	Articles []*Article `orm:"rel(m2m)"`
}

// 文章结构体
type Article struct {
	Id       int       `orm:"pk;auto"`                     // pk主键，auto自增
	Aname    string    `orm:"size(20)"`                    // 限制长度
	Atime    time.Time `orm:"auto_now_add;type(datetime)"` // auto_now_add保存时自动更新时间，type设置时间格式
	Acount   int       `orm:"default(0);null"`             // default默认值
	Acontent string
	Aimg     string `orm:"null"` // null允许为空
	ArtiType *ArticleType `orm:"rel(fk)"`
	Users    []*User      `orm:"reverse(many)"`
}

// 文章类型结构体
type ArticleType struct {
	Id       int
	TypeName string `orm:"size(20)"`
	Articles []*Article `orm:"reverse(many)"`
}

func init() {
	// 设置数据库基本信息
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	// 映射model数据
	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	// 生成表
	orm.RunSyncdb("default", false, true)
}
