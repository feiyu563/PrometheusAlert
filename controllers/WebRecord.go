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
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

func (c *MainController) RecordData() {
	if !CheckAccount(c.Ctx) {
		c.Data["json"] = map[string]interface{}{"error": "unauthorized"}
		c.ServeJSON()
		return
	}

	draw, _ := c.GetInt("draw")
	start, _ := c.GetInt("start")
	length, _ := c.GetInt("length")
	searchVal := c.GetString("search[value]")

	records, total, filtered, err := models.GetRecordPage(start, length, searchVal)
	if err != nil {
		logs.Error(err)
	}

	displayRecords := make([]models.DisplayRecord, 0)
	for _, r := range records {
		displayRecords = append(displayRecords, r.ToDisplay())
	}

	c.Data["json"] = map[string]interface{}{
		"draw":            draw,
		"recordsTotal":    total,
		"recordsFiltered": filtered,
		"data":            displayRecords,
	}
	c.ServeJSON()
}

func (c *MainController) RecordClean() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	models.RecordClean()
	c.Redirect("/record", 302)
}
