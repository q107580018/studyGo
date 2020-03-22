package main

import (
	_ "github.com/q107580018/studyGo/p6/11Beego入门/models"

	"github.com/astaxie/beego"
	_ "github.com/q107580018/studyGo/p6/11Beego入门/routers"
)

func main() {
	beego.AddFuncMap("PrePage", PrePage)
	beego.AddFuncMap("NextPage", NextPage)
	beego.Run()
}
func PrePage(in int) (out int) {
	if in <= 1 {
		return 1
	}
	return in - 1
}
func NextPage(in, pCount int) (out int) {
	if in >= pCount {
		return in
	} else {

		return in + 1
	}
}
