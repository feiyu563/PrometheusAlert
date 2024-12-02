package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

func (c *MainController) AlertTest() {
	MessageData := c.Input().Get("mtype")
	logsign := "[" + LogsSign() + "]"
	switch MessageData {
	case "wx":
		wxtext := "[PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n>**测试告警**\n>`告警级别:`测试\n**PrometheusAlert**"
		ret := PostToWeiXin(wxtext, beego.AppConfig.String("wxurl"), "jikun.zhang", logsign)
		c.Data["json"] = ret
	case "dd":
		ddtext := "## [PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n\n" + "#### 测试告警\n\n" + "###### 告警级别：测试\n\n##### PrometheusAlert\n\n" + "![PrometheusAlert](" + beego.AppConfig.String("logourl") + ")"
		ret := PostToDingDing("PrometheusAlert", ddtext, beego.AppConfig.String("ddurl"), "15888888888", logsign)
		c.Data["json"] = ret
	case "fs":
		fstext := "[PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n\n" + "测试告警\n\n" + "告警级别：测试\n\nPrometheusAlert\n\n" + "![PrometheusAlert](" + beego.AppConfig.String("logourl") + ")"
		//飞书API要求@ 邮箱地址必须有填充
		ret := PostToFS("PrometheusAlert", fstext, beego.AppConfig.String("fsurl"), "244217140@qq.com", logsign)
		c.Data["json"] = ret
	case "txdx":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := PostTXmessage(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "txdh":
		ret := PostTXphonecall("PrometheusAlertCenter测试告警", beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "hwdx":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := PostHWmessage(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "alydx":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := PostALYmessage(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "alydh":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := PostALYphonecall(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "rlydh":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := PostRLYphonecall(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "email":
		TestEmailMessage := `
            <h3>PrometheusAlert邮件告警测试</h3>
			欢迎使用<a href ="https://feiyu563.gitee.io">PrometheusAlert</a><br>
			`
		ret := SendEmail(TestEmailMessage, beego.AppConfig.String("Default_emails"), beego.AppConfig.String("Email_title"), logsign)
		c.Data["json"] = ret
	case "7moordx":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := Post7MOORmessage(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "7moordh":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := Post7MOORphonecall(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "tg":
		TgMessage := "PrometheusAlertCenter测试告警"
		ret := SendTG(TgMessage, logsign)
		c.Data["json"] = ret
	case "workwechat":
		WorkwechatMessage := "[PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n" + "测试告警\n" + "告警级别：测试\nPrometheusAlert\n" + "![PrometheusAlert](" + beego.AppConfig.String("logourl") + ")"
		ret := SendWorkWechat(beego.AppConfig.String("WorkWechat_ToUser"), beego.AppConfig.String("WorkWechat_ToParty"), beego.AppConfig.String("WorkWechat_ToTag"), WorkwechatMessage, logsign)
		c.Data["json"] = ret
	case "bddx":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := PostBDYmessage(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "bdrl":
		RLMessage := "## [PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n\n" + "#### 测试告警\n\n" + "###### 告警级别：测试\n\n##### PrometheusAlert\n\n" + "![PrometheusAlert](" + beego.AppConfig.String("logourl") + ")"
		ret := PostToRuLiu(beego.AppConfig.String("BDRL_ID"), RLMessage, beego.AppConfig.String("BDRL_URL"), logsign)
		c.Data["json"] = ret
	case "bark":
		TgMessage := "PrometheusAlertCenter测试告警"
		ret := SendBark(TgMessage, logsign)
		c.Data["json"] = ret
	case "voice":
		vMessage := "Prometheus Alert Center 测试告警"
		ret := SendVoice(vMessage, logsign)
		c.Data["json"] = ret
	case "fsapp":
		fstext := "[PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n\n" + "测试告警\n\n" + "告警级别：测试\n\nPrometheusAlert\n\n" + "![PrometheusAlert](" + beego.AppConfig.String("logourl") + ")"
		ret := PostToFeiShuApp("PrometheusAlert", fstext, beego.AppConfig.String("AT_USER_ID"), logsign)
		c.Data["json"] = ret
	default:
		c.Data["json"] = "hahaha!"
	}
	c.ServeJSON()
}

// markdown test
func (c *MainController) MarkdownTest() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	if c.Ctx.Request.Method == "GET" {
		c.Data["IsMarkDownTest"] = true
		c.Data["IsTemplateMenu"] = true
		c.TplName = "markdown_test.html"
		c.Data["IsLogin"] = CheckAccount(c.Ctx)
	} else {
		var p_json interface{}
		var resp string
		JsonContent := c.Input().Get("jsoncontent")
		TplContent := c.Input().Get("tplcontent")
		json.Unmarshal([]byte(JsonContent), &p_json)

		err, tpl := TransformAlertMessage(p_json, TplContent)
		if err != nil {
			resp = err.Error()
		} else {
			resp = tpl
		}
		c.Data["json"] = resp
		c.ServeJSON()
	}

}

// test page
func (c *MainController) Test() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsTest"] = true
	c.Data["IsAlertManageMenu"] = true
	c.TplName = "test.html"
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}
