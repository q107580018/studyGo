package routers

import (
	"github.com/astaxie/beego"
	"github.com/q107580018/studyGo/p6/11Beego入门/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:ShowLogin")
	beego.Router("/register", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/index", &controllers.MainController{}, "get:ShowIndex")
	beego.Router("/add", &controllers.MainController{}, "get:ShowAdd;post:HandleAdd")
	beego.Router("/content", &controllers.MainController{}, "get:ShowContent")
	beego.Router("/update", &controllers.MainController{}, "get:ShowUpdate;post:HandleUpdate")
	beego.Router("/delete", &controllers.MainController{}, "get:HandleDelete")
	beego.Router("/addType", &controllers.MainController{}, "get:ShowAddType;post:HandleAddType")
}
