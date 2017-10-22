package controllers

import (
	"github.com/astaxie/beego"
)

// MainController 网站根目录的控制器
type MainController struct {
	beego.Controller
}

// Get 网站根目录get方法
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

