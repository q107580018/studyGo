package controllers

import (
	"math"
	"path"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/q107580018/studyGo/p6/11Beego入门/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) ShowLogin() {
	c.TplName = "login.tpl"
}
func (c *MainController) HandleLogin() {
	// 1.拿到数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	// 2.判断数据是否合法
	if userName == "" || pwd == "" {
		beego.Info("输入数据不合法")
		c.TplName = "login.tpl"
		return
	}
	// 3.查询账号密码是否正确
	o := orm.NewOrm()
	user := new(models.User)
	user.Name = userName
	err := o.Read(user, "name")
	if err != nil {
		beego.Info("账号不存在")
		c.Redirect("/login", 302)
		return
	}
	if user.Pwd == pwd {
		beego.Info("账号密码正确")
		c.Redirect("/index", 302)
	} else {
		beego.Info("输入的账号密码不正确")
		c.Ctx.WriteString("账号密码不正确，登录失败")
	}
}
func (c *MainController) ShowAdd() {
	var artiType []models.ArticleType
	o := orm.NewOrm()
	_, err := o.QueryTable("ArticleType").All(&artiType)
	if err != nil {
		beego.Info("获取artiType失败", err)
	}

	c.Data["articleType"] = artiType
	c.TplName = "add.tpl"
}
func (c *MainController) HandleAdd() {
	// 1.获取数据
	articleName := c.GetString("articleName")
	articleContent := c.GetString("content")
	articleType := c.GetString("select")
	beego.Info(articleName, articleContent)
	f, h, e := c.GetFile("uploadname")
	var fileTime string
	if e != nil {
		beego.Info("上传的图片为空", e)
	} else {
		defer f.Close()
		// 1.1限定文件格式
		fileExt := path.Ext(h.Filename)
		if fileExt != ".jpg" && fileExt != ".png" {
			beego.Info("上传文件格式错误")
			return
		}
		if h.Size > 1024*1024*10 { // 1.2限定文件大小
			beego.Info("上传文件过大")
			return
		}
		// 1.3文件重命名,防止重复
		fileTime = time.Now().Format("20060102150405")
		c.SaveToFile("uploadname", "./static/images/"+fileTime+h.Filename)

	}
	// 2.判断数据合法性
	if articleName == "" || articleContent == "" {
		beego.Info("添加文章错误，数据不能为空")
		return
	}
	// 3.插入数据
	o := orm.NewOrm()
	arti := new(models.Article)
	arti.Aname = articleName
	arti.Acontent = articleContent
	artiType := models.ArticleType{TypeName: articleType}
	err := o.Read(&artiType, "TypeName")
	if err != nil {
		beego.Info("获取类型错误", err)
		return
	}
	arti.ArtiType = &artiType
	if e == nil {
		arti.Aimg = "./static/images/" + fileTime + h.Filename
	}
	_, err = o.Insert(arti)
	if err != nil {
		beego.Info("插入数据库错误", err)
		return
	}
	// 4.返回界面
	c.Redirect("/index", 302)

}
func (c *MainController) Get() {
	c.TplName = "register.tpl"
}
func (c *MainController) Post() {
	// 1.拿到数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	beego.Info("userName:", userName, "pwd:", pwd)
	// 2.对数据进行校验
	if userName == "" || pwd == "" {
		beego.Info("数据不能为空")
		c.Redirect("/register", 302)
		return
	}
	// 3.插入数据库
	o := orm.NewOrm()
	user := &models.User{}
	user.Name = userName
	user.Pwd = pwd
	_, err := o.Insert(user)
	if err != nil {
		beego.Info("插入数据库失败", err)
		c.Redirect("/register", 302)
		return
	}
	// 4.返回登录界面
	c.Redirect("/login", 302)
}
func (c *MainController) ShowIndex() {
	// 获取数据库资料
	o := orm.NewOrm()
	var articles []models.Article
	qs := o.QueryTable("Article")
	count, _ := qs.Count() // 返回数据条目数
	c.Data["count"] = count
	pageSize := 2                                                   //每页显示的条数
	pageCount := int(math.Ceil(float64(count) / float64(pageSize))) // 统计一共多少页，向上取整
	pageIndex, err := c.GetInt("pageIndex")                         // 获取当前页码
	if err != nil {                                                 //错误代表不含pageIndex参数，说明是首页
		pageIndex = 1
	}
	qs.Limit(pageSize, pageSize*(pageIndex-1)).RelatedSel("ArtiType").All(&articles) // 显示的内容限制范围
	firstPage := false
	if pageIndex == 1 {
		firstPage = true
	}
	lastPage := false
	if pageIndex == pageCount {
		lastPage = true
	}
	var artiType []models.ArticleType
	_, err = o.QueryTable("ArticleType").All(&artiType)
	if err != nil {
		beego.Info("获取artiType失败", err)
	}

	c.Data["articleType"] = artiType
	c.Data["firstPage"] = firstPage
	c.Data["lastPage"] = lastPage
	c.Data["pageIndex"] = pageIndex
	c.Data["articles"] = articles
	c.Data["pageCount"] = pageCount
	c.TplName = "index.tpl"
}
func (c *MainController) ShowContent() {
	// 1.获取Id
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("获取文章Id错误", err)
		return
	}
	// 2.查询数据库获取数据
	o := orm.NewOrm()
	arti := &models.Article{Id: id}
	err = o.Read(arti)
	if err != nil {
		beego.Info("查询错误", err)
		return
	}
	// 3.传递数据给视图
	c.Data["article"] = arti
	c.TplName = "content.tpl"
}
func (c *MainController) ShowUpdate() {
	// 1.获取Id
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("获取文章Id错误", err)
		return
	}
	// 2.查询数据库获取数据
	o := orm.NewOrm()
	arti := &models.Article{Id: id}
	err = o.Read(arti)
	if err != nil {
		beego.Info("查询错误", err)
		return
	}
	// 3.传递数据给视图
	c.Data["article"] = arti
	c.TplName = "update.tpl"
}
func (c *MainController) HandleUpdate() {
	// 1.拿到数据
	Aname := c.GetString("articleName")
	Acontent := c.GetString("content")
	f, h, err := c.GetFile("uploadname")
	var fileTime string
	if err != nil {
		beego.Info("上传文件失败", err)
	} else {
		defer f.Close()
		// 1.1限定文件格式
		fileExt := path.Ext(h.Filename)
		if fileExt != ".jpg" && fileExt != ".png" {
			beego.Info("上传文件格式错误")
			return
		}
		if h.Size > 1024*1024*10 { // 1.2限定文件大小
			beego.Info("上传文件过大")
			return
		}
		// 1.3文件重命名,防止重复
		fileTime = time.Now().Format("20060102150405")
		c.SaveToFile("uploadname", "./static/images/"+fileTime+h.Filename)

	}
	if Aname == "" || Acontent == "" {
		beego.Info("更新数据失败")
		return
	}
	// 2.读取原数据
	o := orm.NewOrm()
	id, _ := c.GetInt("id")
	article := &models.Article{Id: id}
	err = o.Read(article, "id")
	if err != nil {
		beego.Info("读取article id失败")
		return
	}
	// 3.更改数据
	article.Aname = Aname
	article.Acontent = Acontent
	if fileTime != "" {
		article.Aimg = "./static/images/" + fileTime + h.Filename
	}
	// 4.更新数据
	_, err = o.Update(article, "Aname", "Acontent", "Aimg")
	if err != nil {
		beego.Info("更新数据失败", err)
		return
	}
	c.Redirect("/index", 302)
}
func (c *MainController) HandleDelete() {
	Id, err := c.GetInt("id")
	if err != nil {
		beego.Info("获取Id为空", err)
		return
	}
	o := orm.NewOrm()
	arti := &models.Article{Id: Id}
	_, err = o.Delete(arti)
	if err != nil {
		beego.Info("删除失败", err)
		return
	}
	c.Redirect("/index", 302)
}
func (c *MainController) ShowAddType() {
	// 1.读取类型表
	o := orm.NewOrm()
	var artType []*models.ArticleType
	// 2.查询表数据
	_, err := o.QueryTable("ArticleType").All(&artType)
	if err != nil {
		beego.Info("读取article type错误", err)
	}
	c.Data["types"] = artType
	c.TplName = "addType.tpl"
}
func (c *MainController) HandleAddType() {

	o := orm.NewOrm()
	typeName := c.GetString("typeName")
	artiType := models.ArticleType{TypeName: typeName}
	_, err := o.Insert(&artiType)
	if err != nil {
		beego.Info("添加新类型失败", err)
		return
	}
	c.Redirect("/addType", 302)
}
