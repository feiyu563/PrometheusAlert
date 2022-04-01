package controllers

import (
	"PrometheusAlert/models"
	"github.com/astaxie/beego/logs"
	"strconv"
)

//router
func (c *MainController) AlertRouter() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsAlertRouter"] = true
	c.Data["IsAlertManageMenu"] = true
	c.TplName = "alertrouter.html"

	GlobalAlertRouter, _ = models.GetAllAlertRouter()
	c.Data["AlertRouter"] = GlobalAlertRouter

	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

//router add
func (c *MainController) RouterAdd() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	Template, err := models.GetPromtheusTpl()
	if err != nil {
		logs.Error(err)
	}
	c.Data["Template"] = Template
	c.Data["IsAlertRouter"] = true
	c.Data["IsAlertManageMenu"] = true
	c.TplName = "alertrouter_add.html"
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

func (c *MainController) AddRouter() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	//获取表单信息
	tid := c.Input().Get("id")
	name := c.Input().Get("name")
	tpl_id := c.Input().Get("tpl_id")
	rules := c.Input().Get("rules")
	purl := c.Input().Get("purl")
	pat := c.Input().Get("pat")
	var err error
	if len(tid) == 0 {
		id, _ := strconv.Atoi(tid)
		tpl_id_int, _ := strconv.Atoi(tpl_id)
		err = models.AddAlertRouter(id, tpl_id_int, name, rules, purl, pat)
	} else {
		id, _ := strconv.Atoi(tid)
		tpl_id_int, _ := strconv.Atoi(tpl_id)
		err = models.UpdateAlertRouter(id, tpl_id_int, name, rules, purl, pat)
	}
	var resp interface{}
	resp = err
	if err != nil {
		resp = err.Error()
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

//router edit
func (c *MainController) RouterEdit() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsAlertRouter"] = true
	c.Data["IsAlertManageMenu"] = true
	c.TplName = "alertrouter_edit.html"
	s_id, _ := strconv.Atoi(c.Input().Get("id"))
	AlertRouter, err := models.GetAlertRouter(s_id)
	if err != nil {
		logs.Error(err)
	}
	c.Data["AlertRouter"] = AlertRouter
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

func (c *MainController) RouterDel() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	s_id, _ := strconv.Atoi(c.Input().Get("id"))
	err := models.DelAlertRouter(s_id)
	if err != nil {
		logs.Error(err)
	}
	c.Redirect("/alertrouter", 302)
}
