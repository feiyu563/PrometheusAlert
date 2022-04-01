package controllers

import (
	"bytes"
	"encoding/json"
	"text/template"
)

func (c *MainController) SetupWeixin() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	if c.Ctx.Request.Method == "GET" {
		c.Data["IsWeixin"] = true
		c.Data["IsSetupMenu"] = true
		c.TplName = "setup_weixin.html"
		c.Data["IsLogin"] = CheckAccount(c.Ctx)
	} else {
		var p_json interface{}
		var resp string
		JsonContent := c.Input().Get("jsoncontent")
		TplContent := c.Input().Get("tplcontent")
		json.Unmarshal([]byte(JsonContent), &p_json)

		funcMap := template.FuncMap{
			"GetCSTtime": GetCSTtime,
			"TimeFormat": TimeFormat,
			"GetTime":    GetTime,
		}
		buf := new(bytes.Buffer)
		tpl, err := template.New("").Funcs(funcMap).Parse(TplContent)
		if err != nil {
			resp = err.Error()
		} else {
			err = tpl.Execute(buf, p_json)
			if err != nil {
				resp = err.Error()
			} else {
				resp = buf.String()
			}
		}
		c.Data["json"] = resp
		c.ServeJSON()
	}

}
