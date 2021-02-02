package controllers

import (
	"PrometheusAlert/model"
	"encoding/json"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type GrafanaController struct {
	beego.Controller
}

var PhoneCallMessage = ""

// {"evalMatches":[],"message":"5分钟内申请云服务流量低于100","ruleId":6,"ruleName":"云服务任务成功数量过低","ruleUrl":"http://grafana.haimacloud.com/d/pH9lfnrmk/ias-p3-ji-gao-jing-xiang?fullscreen=true\u0026edit=true\u0026tab=alert\u0026panelId=28\u0026orgId=1","state":"ok","title":"[OK] 云服务任务成功数量过低"}
//{"evalMatches":[{"value":0,"metric":"Count","tags":{}}],"message":"5分钟内申请云服务流量低于100","ruleId":6,"ruleName":"云服务任务成功数量过低","ruleUrl":"http://grafana.haimacloud.com/d/pH9lfnrmk/ias-p3-ji-gao-jing-xiang?fullscreen=true\u0026edit=true\u0026tab=alert\u0026panelId=28\u0026orgId=1","state":"alerting","title":"[Alerting] 云服务任务成功数量过低"}

type Grafana struct {
	Message  string `json:"message"`
	RuleName string `json:"ruleName"`
	RuleUrl  string `json:"ruleUrl"`
	State    string `json:"state"`
}

func (c *GrafanaController) GrafanaEmail() {
	alert := Grafana{}
	email := c.GetString("email")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageGrafana(alert, 1, logsign, "", "", "", "", "", "", "", "", "", email, "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}
func (c *GrafanaController) GrafanaDingding() {
	alert := Grafana{}
	ddurl := c.GetString("ddurl")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageGrafana(alert, 2, logsign, ddurl, "", "", "", "", "", "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *GrafanaController) GrafanaWeixin() {
	alert := Grafana{}
	wxurl := c.GetString("wxurl")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageGrafana(alert, 3, logsign, "", wxurl, "", "", "", "", "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *GrafanaController) GrafanaTxdh() {
	alert := Grafana{}
	phone := c.GetString("phone")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageGrafana(alert, 4, logsign, "", "", "", "", phone, "", "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *GrafanaController) GrafanaTxdx() {
	alert := Grafana{}
	phone := c.GetString("phone")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageGrafana(alert, 5, logsign, "", "", "", phone, "", "", "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *GrafanaController) GrafanaHwdx() {
	alert := Grafana{}
	phone := c.GetString("phone")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageGrafana(alert, 6, logsign, "", "", "", "", "", phone, "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *GrafanaController) GrafanaALYdx() {
	alert := Grafana{}
	phone := c.GetString("phone")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageGrafana(alert, 7, logsign, "", "", "", "", "", "", "", phone, "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *GrafanaController) GrafanaALYdh() {
	alert := Grafana{}
	phone := c.GetString("phone")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageGrafana(alert, 8, logsign, "", "", "", "", "", "", "", "", phone, "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *GrafanaController) GrafanaRlydh() {
	alert := Grafana{}
	phone := c.GetString("phone")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageGrafana(alert, 9, logsign, "", "", "", "", "", "", phone, "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *GrafanaController) GrafanaFeishu() {
	alert := Grafana{}
	fsurl := c.GetString("fsurl")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageGrafana(alert, 10, logsign, "", "", fsurl, "", "", "", "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *GrafanaController) GrafanaTG() {
	alert := Grafana{}
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageGrafana(alert, 11, logsign, "", "", "", "", "", "", "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *GrafanaController) GrafanaWorkWechat() {
	alert := Grafana{}
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageGrafana(alert, 12, logsign, "", "", "", "", "", "", "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *GrafanaController) GrafanaBddx() {
	alert := Grafana{}
	phone := c.GetString("phone")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageGrafana(alert, 13, logsign, "", "", "", "", "", "", "", "", "", "", phone,"")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *GrafanaController) GrafanaRuLiu() {
	alert := Grafana{}
	groupid := c.GetString("groupid")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageGrafana(alert, 14, logsign, "", "", "", "", "", "", "", "", "", "", "",groupid)
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

//typeid 为0,触发电话告警和钉钉告警, typeid 为1 仅触发dingding告警
func SendMessageGrafana(message Grafana, typeid int, logsign, ddurl, wxurl, fsurl, txdx, txdh, hwdx, rlydh, alydx, alydh, email, bddx,groupid  string) string {
	Title := beego.AppConfig.String("title")
	Logourl := beego.AppConfig.String("logourl")
	Rlogourl := beego.AppConfig.String("rlogourl")
	var DDtext, RLtext,FStext, WXtext, EmailMessage, titleend string
	//告警级别定义 0 信息,1 警告,2 一般严重,3 严重,4 灾难
	AlertLevel := []string{"信息", "警告", "一般严重", "严重", "灾难"}
	if message.State == "ok" {
		titleend = "故障恢复信息"
		model.AlertsFromCounter.WithLabelValues("grafana", message.Message, "4", "", "resolved").Add(1)
		DDtext = "## [" + Title + "Grafana" + titleend + "](" + message.RuleUrl + ")\n\n#### " + message.RuleName + "\n\n###### 告警级别：" + AlertLevel[4] + "\n\n###### 开始时间：" + time.Now().Format("2006-01-02 15:04:05") + "\n\n##### " + message.Message + " 已经恢复正常\n\n" + "![" + Title + "](" + Rlogourl + ")"
		FStext = "[" + Title + "Grafana" + titleend + "](" + message.RuleUrl + ")\n\n" + message.RuleName + "\n\n告警级别：" + AlertLevel[4] + "\n\n开始时间：" + time.Now().Format("2006-01-02 15:04:05") + "\n\n" + message.Message + " 已经恢复正常\n\n" + "![" + Title + "](" + Rlogourl + ")"
		WXtext = "[" + Title + "Grafana" + titleend + "](" + message.RuleUrl + ")\n>**" + message.RuleName + "**\n>`告警级别:`" + AlertLevel[4] + "\n`开始时间:`" + time.Now().Format("2006-01-02 15:04:05") + "\n" + message.Message + " 已经恢复正常\n"
		PhoneCallMessage = message.Message + " 已经恢复正常"
		EmailMessage = `<h1><a href =` + message.RuleUrl + `>` + Title + "Grafana" + titleend + `</a></h1>
				<h2>` + message.RuleName + `</h2>
				<h5>告警级别：` + AlertLevel[4] + `</h5>
				<h5>开始时间：` + time.Now().Format("2006-01-02 15:04:05") + `</h5>
				<h3>` + message.Message + `</h3>
				<img src=` + Rlogourl + ` />`
	} else {
		titleend = "故障告警信息"
		model.AlertsFromCounter.WithLabelValues("grafana", message.Message, "4", "", "firing").Add(1)
		DDtext = "## [" + Title + "Grafana" + titleend + "](" + message.RuleUrl + ")\n\n" + "#### " + message.RuleName + "\n\n" + "###### 告警级别：" + AlertLevel[4] + "\n\n" + "###### 开始时间：" + time.Now().Format("2006-01-02 15:04:05") + "\n\n" + "##### " + message.Message + "\n\n" + "![" + Title + "](" + Logourl + ")"
		RLtext = "## [" + Title + "Grafana" + titleend + "](" + message.RuleUrl + ")\n\n" + "#### " + message.RuleName + "\n\n" + "###### 告警级别：" + AlertLevel[4] + "\n\n" + "###### 开始时间：" + time.Now().Format("2006-01-02 15:04:05") + "\n\n" + "##### " + message.Message + "\n\n" + "![" + Title + "](" + Logourl + ")"
		FStext = "[" + Title + "Grafana" + titleend + "](" + message.RuleUrl + ")\n\n" + "" + message.RuleName + "\n\n" + "告警级别：" + AlertLevel[4] + "\n\n" + "开始时间：" + time.Now().Format("2006-01-02 15:04:05") + "\n\n" + "" + message.Message + "\n\n" + "![" + Title + "](" + Logourl + ")"
		WXtext = "[" + Title + "Grafana" + titleend + "](" + message.RuleUrl + ")\n>**" + message.RuleName + "**\n>`告警级别:`" + AlertLevel[4] + "\n`开始时间:`" + time.Now().Format("2006-01-02 15:04:05") + "\n" + message.Message + "\n"
		PhoneCallMessage = message.Message
		EmailMessage = `<h1><a href =` + message.RuleUrl + `>` + Title + "Grafana" + titleend + `</a></h1>
				<h2>` + message.RuleName + `</h2>
				<h5>告警级别：` + AlertLevel[4] + `</h5>
				<h5>开始时间：` + time.Now().Format("2006-01-02 15:04:05") + `</h5>
				<h3>` + message.Message + `</h3>
				<img src=` + Logourl + ` />`
	}
	//触发email
	if typeid == 1 {
		if email == "" {
			email = beego.AppConfig.String("Default_emails")
		}
		SendEmail(EmailMessage, email, logsign)
	}
	//触发钉钉
	if typeid == 2 {
		if ddurl == "" {
			ddurl = beego.AppConfig.String("ddurl")
		}
		PostToDingDing(Title+titleend, DDtext, ddurl, logsign)
	}
	//触发微信
	if typeid == 3 {
		if wxurl == "" {
			wxurl = beego.AppConfig.String("wxurl")
		}
		PostToWeiXin(WXtext, wxurl, logsign)
	}

	//取到手机号

	//触发电话告警
	if typeid == 4 {
		if txdh == "" {
			txdh = GetUserPhone(1)
		}
		PostTXphonecall(PhoneCallMessage, txdh, logsign)
	}
	//触发腾讯云短信告警
	if typeid == 5 {
		if txdx == "" {
			txdx = GetUserPhone(1)
		}
		PostTXmessage(PhoneCallMessage, txdx, logsign)
	}
	//触发华为云短信告警
	if typeid == 6 {
		if hwdx == "" {
			hwdx = GetUserPhone(1)
		}
		PostHWmessage(PhoneCallMessage, hwdx, logsign)
	}
	//触发阿里云短信告警
	if typeid == 7 {
		if alydx == "" {
			alydx = GetUserPhone(1)
		}
		PostALYmessage(PhoneCallMessage, alydx, logsign)
	}
	//触发阿里云电话告警
	if typeid == 8 {
		if alydh == "" {
			alydh = GetUserPhone(1)
		}
		PostALYphonecall(PhoneCallMessage, alydh, logsign)
	}
	//触发容联云电话告警
	if typeid == 9 {
		if rlydh == "" {
			rlydh = GetUserPhone(1)
		}
		PostRLYphonecall(PhoneCallMessage, rlydh, logsign)
	}
	//触发飞书
	if typeid == 10 {
		if fsurl == "" {
			fsurl = beego.AppConfig.String("fsurl")
		}
		PostToFeiShu(Title+titleend, FStext, fsurl, logsign)
	}
	//触发TG
	if typeid == 11 {
		SendTG(PhoneCallMessage, logsign)
	}
	//触发企业微信消息
	if typeid == 12 {
		SendWorkWechat(beego.AppConfig.String("WorkWechat_ToUser"),beego.AppConfig.String("WorkWechat_ToParty"), beego.AppConfig.String("WorkWechat_ToTag"),WXtext, logsign)
	}
	//触发百度云短信告警
	if typeid == 13 {
		if bddx == "" {
			bddx = GetUserPhone(1)
		}
		PostBDYmessage(PhoneCallMessage, bddx, logsign)
	}
	//触发百度Hi(如流)
	if typeid == 14 {
		if groupid == "" {
			groupid = beego.AppConfig.String("BDRL_ID")
		}
		PostToRuLiu(groupid, RLtext, beego.AppConfig.String("BDRL_URL"), logsign)
	}
	return "告警消息发送完成."
}
