package controllers

import (
	"PrometheusAlert/model"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type Graylog2Controller struct {
	beego.Controller
}

//graylog2告警部分
type Graylog2 struct {
	Check_result Check_result `json:"check_result"`
}
type Check_result struct {
	MatchingMessages   []MatchingMessage `json:"matching_messages"`
	Result_description string            `json:"result_description"`
}
type MatchingMessage struct {
	Index     string  `json:"index"`
	Message   string  `json:"message"`
	Fields    G2Field `json:"fields"`
	Timestamp string  `json:"timestamp"`
}
type G2Field struct {
	Gl2RemoteIp   string `json:"gl2_remote_ip"`
	Gl2RemotePort int    `json:"gl_2_remote_port"`
}

//for graylog alert
func (c *Graylog2Controller) GraylogEmail() {
	alert := Graylog2{}
	email := c.GetString("email")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageG(alert, 1, logsign, "", "", "", "", "", "", "", "", "", email, "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogDingding() {
	alert := Graylog2{}
	ddurl := c.GetString("ddurl")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageG(alert, 2, logsign, ddurl, "", "", "", "", "", "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogWeixin() {
	alert := Graylog2{}
	wxurl := c.GetString("wxurl")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageG(alert, 3, logsign, "", wxurl, "", "", "", "", "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogTxdh() {
	alert := Graylog2{}
	phone := c.GetString("phone")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageG(alert, 4, logsign, "", "", "", "", phone, "", "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogTxdx() {
	alert := Graylog2{}
	phone := c.GetString("phone")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageG(alert, 5, logsign, "", "", "", phone, "", "", "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogHwdx() {
	alert := Graylog2{}
	phone := c.GetString("phone")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageG(alert, 6, logsign, "", "", "", "", "", phone, "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogALYdx() {
	alert := Graylog2{}
	phone := c.GetString("phone")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageG(alert, 7, logsign, "", "", "", "", "", "", "", phone, "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogALYdh() {
	alert := Graylog2{}
	phone := c.GetString("phone")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageG(alert, 8, logsign, "", "", "", "", "", "", "", "", phone, "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogRLYdh() {
	alert := Graylog2{}
	phone := c.GetString("phone")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageG(alert, 9, logsign, "", "", "", "", "", "", phone, "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *Graylog2Controller) GraylogFeishu() {
	alert := Graylog2{}
	fsurl := c.GetString("fsurl")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageG(alert, 10, logsign, "", "", fsurl, "", "", "", "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *Graylog2Controller) GraylogTG() {
	alert := Graylog2{}
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageG(alert, 11, logsign, "", "", "", "", "", "", "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *Graylog2Controller) GraylogWorkWechat() {
	alert := Graylog2{}
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageG(alert, 12, logsign, "", "", "", "", "", "", "", "", "", "", "","")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogBddx() {
	alert := Graylog2{}
	phone := c.GetString("phone")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageG(alert, 13, logsign, "", "", "", "", "", "", "", "", "", "", phone,"")
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogRuLiu() {
	alert := Graylog2{}
	groupid := c.GetString("groupid")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageG(alert, 14, logsign, "", "", "", "", "", "", "", "", "", "", "",groupid)
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func SendMessageG(message Graylog2, typeid int, logsign, ddurl, wxurl, fsurl, txdx, txdh, hwdx, rlydh, alydx, alydh, email, bddx,groupid string) string {
	Title := beego.AppConfig.String("title")
	Alerturl := beego.AppConfig.String("GraylogAlerturl")
	Logourl := beego.AppConfig.String("logourl")
	PCstTime, _ := beego.AppConfig.Int("prometheus_cst_time")
	if len(message.Check_result.MatchingMessages) == 0 {
		model.AlertsFromCounter.WithLabelValues("graylog2", message.Check_result.Result_description, "4", "", "").Add(1)
		if ddurl == "" {
			ddurl = beego.AppConfig.String("ddurl")
		}
		PostToDingDing(Title+"告警信息", "## ["+Title+"Graylog2告警信息]("+Alerturl+")\n\n"+"#### "+message.Check_result.Result_description+"\n\n"+"!["+Title+"]("+Logourl+")", ddurl, logsign)
		if fsurl == "" {
			fsurl = beego.AppConfig.String("fsurl")
		}
		PostToFeiShu(Title+"告警信息", "["+Title+"Graylog2告警信息]("+Alerturl+")\n\n"+""+message.Check_result.Result_description+"\n\n"+"!["+Title+"]("+Logourl+")", fsurl, logsign)
		if wxurl == "" {
			wxurl = beego.AppConfig.String("wxurl")
		}
		PostToWeiXin("["+Title+"Graylog2告警信息]("+Alerturl+")\n>**"+message.Check_result.Result_description+"**", wxurl, logsign)
		if email == "" {
			email = beego.AppConfig.String("Default_emails")
		}
		EmailMessage := `<h1><a href =` + Alerturl + `>` + Title + "Graylog2告警信息" + `</a></h1>
				<h2>` + message.Check_result.Result_description + `</h2>
				<img src=` + Logourl + ` />`
		SendEmail(EmailMessage, email, logsign)
		return "告警消息发送完成."
	}
	for _, m := range message.Check_result.MatchingMessages {
		GraylogTime := m.Timestamp
		if PCstTime == 1 {
			GraylogTime = GetCSTtime(m.Timestamp)
		}
		model.AlertsFromCounter.WithLabelValues("graylog2", message.Check_result.Result_description, "4", m.Fields.Gl2RemoteIp, "firing").Add(1)
		DDtext := "## [" + Title + "Graylog2告警信息](" + Alerturl + ")\n\n" + "#### " + message.Check_result.Result_description + "\n\n" + "###### 告警索引：" + m.Index + "\n\n" + "###### 开始时间：" + GraylogTime + " \n\n" + "###### 告警主机：" + m.Fields.Gl2RemoteIp + ":" + strconv.Itoa(m.Fields.Gl2RemotePort) + "\n\n" + "##### " + m.Message + "\n\n" + "![" + Title + "](" + Logourl + ")"
		RLtext := "## [" + Title + "Graylog2告警信息](" + Alerturl + ")\n\n" + "#### " + message.Check_result.Result_description + "\n\n" + "###### 告警索引：" + m.Index + "\n\n" + "###### 开始时间：" + GraylogTime + " \n\n" + "###### 告警主机：" + m.Fields.Gl2RemoteIp + ":" + strconv.Itoa(m.Fields.Gl2RemotePort) + "\n\n" + "##### " + m.Message + "\n\n" + "![" + Title + "](" + Logourl + ")"
		FStext := "[" + Title + "Graylog2告警信息](" + Alerturl + ")\n\n" + "" + message.Check_result.Result_description + "\n\n" + "告警索引：" + m.Index + "\n\n" + "开始时间：" + GraylogTime + " \n\n" + "告警主机：" + m.Fields.Gl2RemoteIp + ":" + strconv.Itoa(m.Fields.Gl2RemotePort) + "\n\n" + "" + m.Message + "\n\n" + "![" + Title + "](" + Logourl + ")"
		WXtext := "[" + Title + "Graylog2告警信息](" + Alerturl + ")\n>**" + message.Check_result.Result_description + "**\n>`告警索引:`" + m.Index + "\n`开始时间:`" + GraylogTime + " \n`告警主机:`" + m.Fields.Gl2RemoteIp + ":" + strconv.Itoa(m.Fields.Gl2RemotePort) + "\n**" + m.Message + "**"
		PhoneCallMessage = "告警主机 " + m.Fields.Gl2RemoteIp + "端口 " + strconv.Itoa(m.Fields.Gl2RemotePort) + "告警消息 " + m.Message
		EmailMessage := `<h1><a href =` + Alerturl + `>` + Title + "Graylog2告警信息" + `</a></h1>
				<h2>` + message.Check_result.Result_description + `</h2>
				<h5>告警索引：` + m.Index + `</h5>
				<h5>开始时间：` + GraylogTime + `</h5>
				<h5>告警主机：` + m.Fields.Gl2RemoteIp + `:` + strconv.Itoa(m.Fields.Gl2RemotePort) + `</h5>
				<h3>` + m.Message + `</h3>
				<img src=` + Logourl + ` />`
		//触发邮件
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
			PostToDingDing(Title+"告警信息", DDtext, ddurl, logsign)
		}
		//触发微信
		if typeid == 3 {
			if wxurl == "" {
				wxurl = beego.AppConfig.String("wxurl")
			}
			PostToWeiXin(WXtext, wxurl, logsign)
		}
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
			PostToFeiShu(Title+"告警信息", FStext, fsurl, logsign)
		}
		//触发TG
		if typeid == 11 {
			SendTG(PhoneCallMessage, logsign)
		}
		//触发企业微信应用消息
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
	}
	return "告警消息发送完成."
}
