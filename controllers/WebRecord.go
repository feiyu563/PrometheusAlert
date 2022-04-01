package controllers

import (
	"PrometheusAlert/models"
	"github.com/astaxie/beego/logs"
)

//Record page
func (c *MainController) Record() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsRecord"] = true
	c.Data["IsAlertManageMenu"] = true
	c.TplName = "record.html"
	Record, err := models.GetAllRecord()
	if err != nil {
		logs.Error(err)
	}
	c.Data["Record"] = Record
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

func (c *MainController) RecordClean() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	models.RecordClean()
	c.Redirect("/record", 302)
}
