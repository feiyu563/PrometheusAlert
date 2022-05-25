package controllers

import (
	"PrometheusAlert/models"
	"github.com/astaxie/beego/logs"
	"strconv"
)

//template page
func (c *MainController) Template() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsTemplate"] = true
	c.Data["IsTemplateMenu"] = true
	c.TplName = "template.html"
	GlobalPrometheusAlertTpl, err := models.GetAllTpl()
	if err != nil {
		logs.Error(err)
	}
	c.Data["Template"] = GlobalPrometheusAlertTpl
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

//template add
func (c *MainController) TemplateAdd() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsTemplate"] = true
	c.Data["IsTemplateMenu"] = true
	c.TplName = "template_add.html"
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}
func (c *MainController) AddTpl() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	//获取表单信息
	tid := c.Input().Get("id")
	name := c.Input().Get("name")
	t_tpye := c.Input().Get("type")
	t_use := c.Input().Get("use")
	content := c.Input().Get("content")
	var err error
	if len(tid) == 0 {
		id, _ := strconv.Atoi(tid)
		err = models.AddTpl(id, name, t_tpye, t_use, content)
	} else {
		id, _ := strconv.Atoi(tid)
		err = models.UpdateTpl(id, name, t_tpye, t_use, content)
	}
	var resp interface{}
	if err != nil {
		resp = err.Error()
	} else {
		resp = err
		GlobalPrometheusAlertTpl, _ = models.GetAllTpl()
	}
	GlobalAlertRouter, _ = models.GetAllAlertRouter()
	c.Data["json"] = resp
	c.ServeJSON()
}
func (c *MainController) TemplateEdit() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsTemplate"] = true
	c.Data["IsTemplateMenu"] = true
	c.TplName = "template_edit.html"
	s_id, _ := strconv.Atoi(c.Input().Get("id"))
	Template, err := models.GetTpl(s_id)
	if err != nil {
		logs.Error(err)
	} else {
		GlobalPrometheusAlertTpl, _ = models.GetAllTpl()
	}
	c.Data["Template"] = Template
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

func (c *MainController) TemplateDel() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	s_id, _ := strconv.Atoi(c.Input().Get("id"))
	err := models.DelTpl(s_id)
	if err != nil {
		logs.Error(err)
	} else {
		GlobalPrometheusAlertTpl, _ = models.GetAllTpl()
	}
	c.Redirect("/template", 302)
}
