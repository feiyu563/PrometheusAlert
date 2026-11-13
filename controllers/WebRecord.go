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

func (c *MainController) RecordFilters() {
	if !CheckAccount(c.Ctx) {
		c.Data["json"] = map[string]interface{}{"error": "unauthorized"}
		c.ServeJSON()
		return
	}
	sources, channels, statuses, _ := models.GetRecordFilters()

	type FilterOption struct {
		Value string `json:"value"`
		Text  string `json:"text"`
	}

	formatOpts := func(items []string, isChannel, isStatus bool) []FilterOption {
		opts := make([]FilterOption, 0)
		for _, item := range items {
			text := item
			if isChannel {
				text = models.GetChannelName(item)
			} else if isStatus {
				switch item {
				case "success": text = "成功"
				case "failed":  text = "失败"
				case "firing":  text = "告警"
				case "resolved": text = "恢复"
				}
			}
			opts = append(opts, FilterOption{Value: item, Text: text})
		}
		return opts
	}

	c.Data["json"] = map[string]interface{}{
		"sources":  formatOpts(sources, false, false),
		"channels": formatOpts(channels, true, false),
		"statuses": formatOpts(statuses, false, true),
	}
	c.ServeJSON()
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
	source := c.GetString("source")
	channel := c.GetString("channel")
	status := c.GetString("status")

	records, total, filtered, err := models.GetRecordPage(start, length, searchVal, source, channel, status)
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
