package controllers

import (
	"PrometheusAlert/model"
	"encoding/json"
	"sort"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type PrometheusController struct {
	beego.Controller
}

type Labels struct {
	Alertname string `json:"alertname"`
	Instance  string `json:"instance"`
	Level     string `json:"level"` //2019年11月20日 16:03:10更改告警级别定义位置,适配prometheus alertmanager rule
}
type Annotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
	//Level string `json:"level"`  //2019年11月20日 16:04:04 删除Annotations level,改用label中的level
	Mobile string `json:"mobile"` //2019年2月25日 19:09:23 增加手机号支持
	Ddurl  string `json:"ddurl"`  //2019年3月12日 20:33:38 增加多个钉钉告警支持
	Wxurl  string `json:"wxurl"`  //2019年3月12日 20:33:38 增加多个钉钉告警支持
	Fsurl  string `json:"fsurl"`  //2020年4月25日 17:33:38 增加多个飞书告警支持
	Email  string `json:"email"`  //2020年7月4日 10:15:20 增加多个email告警支持
	Groupid  string `json:"groupid"`  //2021年2月2日 17:28:23 增加多个如流告警支持
}
type Alerts struct {
	Status       string
	Labels       Labels      `json:"labels"`
	Annotations  Annotations `json:"annotations"`
	StartsAt     string      `json:"startsAt"`
	EndsAt       string      `json:"endsAt"`
	GeneratorUrl string      `json:"generatorURL"` //prometheus 告警返回地址
}
type Prometheus struct {
	Status      string
	Alerts      []Alerts
	Externalurl string `json:"externalURL"` //alertmanage 返回地址
}

// 按照 Alert.Level 从大到小排序
type AlerMessages []Alerts

func (a AlerMessages) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a AlerMessages) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a AlerMessages) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Labels.Level < a[i].Labels.Level
}

//for prometheus alert
//关于告警级别level共有5个级别,0-4,0 信息,1 警告,2 一般严重,3 严重,4 灾难
func (c *PrometheusController) PrometheusAlert() {
	alert := Prometheus{}
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageR(alert, "", "", "", "", "","", logsign)
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func (c *PrometheusController) PrometheusRouter() {
	wxurl := c.GetString("wxurl")
	ddurl := c.GetString("ddurl")
	fsurl := c.GetString("fsurl")
	phone := c.GetString("phone")
	email := c.GetString("email")
	groupid := c.GetString("groupid")
	logsign := "[" + LogsSign() + "]"
	alert := Prometheus{}
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageR(alert, wxurl, ddurl, fsurl, phone, email,groupid,logsign)
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func SendMessageR(message Prometheus, rwxurl, rddurl, rfsurl, rphone, remail,rgroupid,logsign string) string {
	//增加日志标志  方便查询日志

	Title := beego.AppConfig.String("title")
	Logourl := beego.AppConfig.String("logourl")
	Rlogourl := beego.AppConfig.String("rlogourl")
	Messagelevel, _ := beego.AppConfig.Int("messagelevel")
	PhoneCalllevel, _ := beego.AppConfig.Int("phonecalllevel")
	PhoneCallResolved, _ := beego.AppConfig.Int("phonecallresolved")
	Silent, _ := beego.AppConfig.Int("silent")
	PCstTime, _ := beego.AppConfig.Int("prometheus_cst_time")
	var ddtext, wxtext, fstext, MobileMessage, PhoneCallMessage, EmailMessage, titleend,rltext string
	//对分组消息进行排序
	AlerMessage := message.Alerts
	sort.Sort(AlerMessages(AlerMessage))
	//告警级别定义 0 信息,1 警告,2 一般严重,3 严重,4 灾难
	AlertLevel := []string{"信息", "警告", "一般严重", "严重", "灾难"}
	//遍历消息
	for _, RMessage := range AlerMessage {
		nLevel, _ := strconv.Atoi(RMessage.Labels.Level)
		At := RMessage.StartsAt
		Et := RMessage.EndsAt
		if PCstTime == 1 {
			At = GetCSTtime(RMessage.StartsAt)
			Et = GetCSTtime(RMessage.EndsAt)
		}
		if RMessage.Status == "resolved" {
			titleend = "故障恢复信息"
			model.AlertsFromCounter.WithLabelValues("prometheus", RMessage.Annotations.Description, RMessage.Labels.Level, RMessage.Labels.Instance, "resolved").Add(1)
			ddtext = "## [" + Title + "Prometheus" + titleend + "](" + RMessage.GeneratorUrl + ")\n\n" + "#### [" + RMessage.Labels.Alertname + "](" + message.Externalurl + ")\n\n" + "###### 告警级别：" + AlertLevel[nLevel] + "\n\n" + "###### 开始时间：" + At + "\n\n" + "###### 结束时间：" + Et + "\n\n" + "###### 故障主机IP：" + RMessage.Labels.Instance + "\n\n" + "##### " + RMessage.Annotations.Description + "\n\n" + "![" + Title + "](" + Rlogourl + ")"
			rltext = "## [" + Title + "Prometheus" + titleend + "](" + RMessage.GeneratorUrl + ")\n\n" + "#### [" + RMessage.Labels.Alertname + "](" + message.Externalurl + ")\n\n" + "###### 告警级别：" + AlertLevel[nLevel] + "\n\n" + "###### 开始时间：" + At + "\n\n" + "###### 结束时间：" + Et + "\n\n" + "###### 故障主机IP：" + RMessage.Labels.Instance + "\n\n" + "##### " + RMessage.Annotations.Description + "\n\n" + "![" + Title + "](" + Rlogourl + ")"
			fstext = "## [" + Title + "Prometheus" + titleend + "](" + RMessage.GeneratorUrl + ")\n\n" + "#### [" + RMessage.Labels.Alertname + "](" + message.Externalurl + ")\n\n" + "###### 告警级别：" + AlertLevel[nLevel] + "\n\n" + "###### 开始时间：" + At + "\n\n" + "###### 结束时间：" + Et + "\n\n" + "###### 故障主机IP：" + RMessage.Labels.Instance + "\n\n" + "##### " + RMessage.Annotations.Description + "\n\n" + "![" + Title + "](" + Rlogourl + ")"
			wxtext = "[" + Title + "Prometheus" + titleend + "](" + RMessage.GeneratorUrl + ")\n>**[" + RMessage.Labels.Alertname + "](" + message.Externalurl + ")**\n>`告警级别:`" + AlertLevel[nLevel] + "\n`开始时间:`" + At + "\n`结束时间:`" + Et + "\n`故障主机IP:`" + RMessage.Labels.Instance + "\n**" + RMessage.Annotations.Description + "**"
			MobileMessage = "\n[" + Title + "Prometheus" + titleend + "]\n" + RMessage.Labels.Alertname + "\n" + "告警级别：" + AlertLevel[nLevel] + "\n" + "故障主机IP：" + RMessage.Labels.Instance + "\n" + RMessage.Annotations.Description
			PhoneCallMessage = "故障主机IP " + RMessage.Labels.Instance + RMessage.Annotations.Description + "已经恢复"
			EmailMessage = `<h1><a href =` + RMessage.GeneratorUrl + `>` + Title + "Prometheus" + titleend + `</a></h1>
				<h2><a href ` + message.Externalurl + `>` + RMessage.Labels.Alertname + `</a></h2>
				<h5>告警级别：` + AlertLevel[nLevel] + `</h5>
				<h5>开始时间：` + At + `</h5>
				<h5>结束时间：` + Et + `</h5>
				<h5>故障主机IP：` + RMessage.Labels.Instance + `</h5>
				<h3>` + RMessage.Annotations.Description + `</h3>
				<img src=` + Rlogourl + ` />`
		} else {
			titleend = "故障告警信息"
			model.AlertsFromCounter.WithLabelValues("prometheus", RMessage.Annotations.Description, RMessage.Labels.Level, RMessage.Labels.Instance, "firing").Add(1)
			ddtext = "## [" + Title + "Prometheus" + titleend + "](" + RMessage.GeneratorUrl + ")\n\n" + "#### [" + RMessage.Labels.Alertname + "](" + message.Externalurl + ")\n\n" + "###### 告警级别：" + AlertLevel[nLevel] + "\n\n" + "###### 开始时间：" + At + "\n\n" + "###### 结束时间：" + Et + "\n\n" + "###### 故障主机IP：" + RMessage.Labels.Instance + "\n\n" + "##### " + RMessage.Annotations.Description + "\n\n" + "![" + Title + "](" + Logourl + ")"
			rltext = "## [" + Title + "Prometheus" + titleend + "](" + RMessage.GeneratorUrl + ")\n\n" + "#### [" + RMessage.Labels.Alertname + "](" + message.Externalurl + ")\n\n" + "###### 告警级别：" + AlertLevel[nLevel] + "\n\n" + "###### 开始时间：" + At + "\n\n" + "###### 结束时间：" + Et + "\n\n" + "###### 故障主机IP：" + RMessage.Labels.Instance + "\n\n" + "##### " + RMessage.Annotations.Description + "\n\n" + "![" + Title + "](" + Logourl + ")"
			fstext = "## [" + Title + "Prometheus" + titleend + "](" + RMessage.GeneratorUrl + ")\n\n" + "#### [" + RMessage.Labels.Alertname + "](" + message.Externalurl + ")\n\n" + "###### 告警级别：" + AlertLevel[nLevel] + "\n\n" + "###### 开始时间：" + At + "\n\n" + "###### 结束时间：" + Et + "\n\n" + "###### 故障主机IP：" + RMessage.Labels.Instance + "\n\n" + "##### " + RMessage.Annotations.Description + "\n\n" + "![" + Title + "](" + Logourl + ")"
			wxtext = "[" + Title + "Prometheus" + titleend + "](" + RMessage.GeneratorUrl + ")\n>**[" + RMessage.Labels.Alertname + "](" + message.Externalurl + ")**\n>`告警级别:`" + AlertLevel[nLevel] + "\n`开始时间:`" + At + "\n`结束时间:`" + Et + "\n`故障主机IP:`" + RMessage.Labels.Instance + "\n**" + RMessage.Annotations.Description + "**"
			MobileMessage = "\n[" + Title + "Prometheus" + titleend + "]\n" + RMessage.Labels.Alertname + "\n" + "告警级别：" + AlertLevel[nLevel] + "\n" + "故障主机IP：" + RMessage.Labels.Instance + "\n" + RMessage.Annotations.Description
			PhoneCallMessage = "故障主机IP " + RMessage.Labels.Instance + RMessage.Annotations.Description
			EmailMessage = `<h1><a href =` + RMessage.GeneratorUrl + `>` + Title + "Prometheus" + titleend + `</a></h1>
				<h2><a href ` + message.Externalurl + `>` + RMessage.Labels.Alertname + `</a></h2>
				<h5>告警级别：` + AlertLevel[nLevel] + `</h5>
				<h5>开始时间：` + At + `</h5>
				<h5>结束时间：` + Et + `</h5>
				<h5>故障主机IP：` + RMessage.Labels.Instance + `</h5>
				<h3>` + RMessage.Annotations.Description + `</h3>
				<img src=` + Logourl + ` />`
		}
		//发送消息到钉钉
		if rddurl == "" && RMessage.Annotations.Ddurl == "" {
			url := beego.AppConfig.String("ddurl")
			PostToDingDing(Title+titleend, ddtext, url, logsign)
		} else {
			if rddurl != "" {
				Ddurl := strings.Split(rddurl, ",")
				for _, url := range Ddurl {
					PostToDingDing(Title+titleend, ddtext, url, logsign)
				}
			}
			if RMessage.Annotations.Ddurl != "" {
				Ddurl := strings.Split(RMessage.Annotations.Ddurl, ",")
				for _, url := range Ddurl {
					PostToDingDing(Title+titleend, ddtext, url, logsign)
				}
			}
		}

		//发送消息到如流
		if rgroupid == "" && RMessage.Annotations.Groupid == "" {
			gid := beego.AppConfig.String("BDRL_ID")
			PostToRuLiu(gid, rltext, beego.AppConfig.String("BDRL_URL"), logsign)
		} else {
			if rgroupid != "" {
				PostToRuLiu(rgroupid, rltext, beego.AppConfig.String("BDRL_URL"), logsign)
			}
			if RMessage.Annotations.Groupid != "" {
				PostToRuLiu(RMessage.Annotations.Groupid, rltext, beego.AppConfig.String("BDRL_URL"), logsign)
			}
		}

		//发送消息到微信
		if rwxurl == "" && RMessage.Annotations.Wxurl == "" {
			url := beego.AppConfig.String("wxurl")
			PostToWeiXin(wxtext, url, logsign)
		} else {
			if rwxurl != "" {
				Wxurl := strings.Split(rwxurl, ",")
				for _, url := range Wxurl {
					PostToWeiXin(wxtext, url, logsign)
				}
			}
			if RMessage.Annotations.Wxurl != "" {
				Wxurl := strings.Split(RMessage.Annotations.Wxurl, ",")
				for _, url := range Wxurl {
					PostToWeiXin(wxtext, url, logsign)
				}
			}
		}
		//发送消息到飞书
		if rfsurl == "" && RMessage.Annotations.Fsurl == "" {
			url := beego.AppConfig.String("fsurl")
			PostToFS(Title+titleend, fstext, url, logsign)
		} else {
			if rfsurl != "" {
				Fsurl := strings.Split(rfsurl, ",")
				for _, url := range Fsurl {
					PostToFS(Title+titleend, fstext, url, logsign)
				}
			}
			if RMessage.Annotations.Fsurl != "" {
				Fsurl := strings.Split(RMessage.Annotations.Fsurl, ",")
				for _, url := range Fsurl {
					PostToFS(Title+titleend, fstext, url, logsign)
				}
			}
		}
		//发送消息到Email
		if remail == "" && RMessage.Annotations.Email == "" {
			Emails := beego.AppConfig.String("Default_emails")
			SendEmail(EmailMessage, Emails, logsign)
		} else {
			if remail != "" {
				SendEmail(EmailMessage, remail, logsign)
			}
			if RMessage.Annotations.Email != "" {
				Emails := RMessage.Annotations.Email
				SendEmail(EmailMessage, Emails, logsign)
			}
		}
		//发送消息到短信
		if nLevel == Messagelevel {
			if rphone == "" && RMessage.Annotations.Mobile == "" {
				phone := GetUserPhone(1)
				PostTXmessage(MobileMessage, phone, logsign)
				PostHWmessage(MobileMessage, phone, logsign)
				PostALYmessage(MobileMessage, phone, logsign)
				Post7MOORmessage(MobileMessage, phone, logsign)
				PostBDYmessage(MobileMessage, phone, logsign)
			} else {
				if rphone != "" {
					PostTXmessage(MobileMessage, rphone, logsign)
					PostHWmessage(MobileMessage, rphone, logsign)
					PostALYmessage(MobileMessage, rphone, logsign)
					Post7MOORmessage(MobileMessage, rphone, logsign)
					PostBDYmessage(MobileMessage, rphone, logsign)
				}
				if RMessage.Annotations.Mobile != "" {
					PostTXmessage(MobileMessage, RMessage.Annotations.Mobile, logsign)
					PostHWmessage(MobileMessage, RMessage.Annotations.Mobile, logsign)
					PostALYmessage(MobileMessage, RMessage.Annotations.Mobile, logsign)
					Post7MOORmessage(MobileMessage, RMessage.Annotations.Mobile, logsign)
					PostBDYmessage(MobileMessage, RMessage.Annotations.Mobile, logsign)
				}
			}
		}
		//发送消息到语音
		if nLevel == PhoneCalllevel {
			//判断如果是恢复信息且PhoneCallResolved
			if RMessage.Status == "resolved" && PhoneCallResolved != 1 {
				logs.Info(logsign, "告警恢复消息已经关闭")
			} else {
				if rphone == "" && RMessage.Annotations.Mobile == "" {
					phone := GetUserPhone(1)
					PostTXphonecall(PhoneCallMessage, phone, logsign)
					PostALYphonecall(PhoneCallMessage, phone, logsign)
					PostRLYphonecall(PhoneCallMessage, phone, logsign)
					Post7MOORphonecall(PhoneCallMessage, phone, logsign)
				} else {
					if rphone != "" {
						PostTXphonecall(PhoneCallMessage, rphone, logsign)
						PostALYphonecall(PhoneCallMessage, rphone, logsign)
						PostRLYphonecall(PhoneCallMessage, rphone, logsign)
						Post7MOORphonecall(PhoneCallMessage, rphone, logsign)
					}
					if RMessage.Annotations.Mobile != "" {
						PostTXphonecall(PhoneCallMessage, RMessage.Annotations.Mobile, logsign)
						PostALYphonecall(PhoneCallMessage, RMessage.Annotations.Mobile, logsign)
						PostRLYphonecall(PhoneCallMessage, RMessage.Annotations.Mobile, logsign)
						Post7MOORphonecall(PhoneCallMessage, RMessage.Annotations.Mobile, logsign)
					}
				}
			}
		}
		// 发送消息到Telegram
		SendTG(PhoneCallMessage, logsign)
		// 推送消息到企业微信
		SendWorkWechat(beego.AppConfig.String("WorkWechat_ToUser"),beego.AppConfig.String("WorkWechat_ToParty"), beego.AppConfig.String("WorkWechat_ToTag"),wxtext, logsign)

		//告警抑制开启就直接跳出循环
		if Silent == 1 {
			break
		}
	}
	return "告警消息发送完成."
}
