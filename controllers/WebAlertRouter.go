package controllers

import (
	"PrometheusAlert/models"
	"encoding/json"
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

type AlertRouterJson struct {
	RouterId    string
	RouterName  string
	RouterTplId string
	RouterPurl  string
	RouterPat   string
	Rules       []LabelMap
}

type LabelMap struct {
	Name  string
	Value string
	Regex bool
}

func (c *MainController) AddRouter() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	WebAlertRouterJson := AlertRouterJson{}
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &WebAlertRouterJson)
	rules, err := json.Marshal(WebAlertRouterJson.Rules)
	if WebAlertRouterJson.RouterId == "" {
		tpl_id_int, _ := strconv.Atoi(WebAlertRouterJson.RouterTplId)
		err = models.AddAlertRouter(0, tpl_id_int, WebAlertRouterJson.RouterName, string(rules), WebAlertRouterJson.RouterPurl, WebAlertRouterJson.RouterPat)
	} else {
		id, _ := strconv.Atoi(WebAlertRouterJson.RouterId)
		tpl_id_int, _ := strconv.Atoi(WebAlertRouterJson.RouterTplId)
		err = models.UpdateAlertRouter(id, tpl_id_int, WebAlertRouterJson.RouterName, string(rules), WebAlertRouterJson.RouterPurl, WebAlertRouterJson.RouterPat)
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
