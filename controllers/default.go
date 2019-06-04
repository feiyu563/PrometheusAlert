package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
//main page
func (c *MainController) Get() {
	c.Data["Email"] = "244217140@qq.com"
	c.TplName = "index.tpl"
}