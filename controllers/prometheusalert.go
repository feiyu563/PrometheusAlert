package controllers

import (
	"PrometheusAlert/model"
	"PrometheusAlert/models"
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type PrometheusAlertController struct {
	beego.Controller
}

func (c *PrometheusAlertController) PrometheusAlert() {
	logsign := "[" + LogsSign() + "]"
	var p_json interface{}
	logs.Debug(logsign, strings.Replace(string(c.Ctx.Input.RequestBody), "\n", "", -1))
	json.Unmarshal(c.Ctx.Input.RequestBody, &p_json)
	P_type := c.Input().Get("type")
	P_tpl := c.Input().Get("tpl")
	P_ddurl := c.Input().Get("ddurl")
	P_wxurl := c.Input().Get("wxurl")
	P_fsurl := c.Input().Get("fsurl")
	P_phone := c.Input().Get("phone")
	if P_phone == "" {
		P_phone = GetUserPhone(1)
	}
	P_email := c.Input().Get("email")
	//get tpl
	message := ""
	if P_tpl != "" && P_type != "" {
		tpltext, err := models.GetTplOne(P_tpl)
		if err != nil {
			logs.Error(logsign, err)
		}
		buf := new(bytes.Buffer)
		tpl, err := template.New("").Funcs(template.FuncMap{"GetCSTtime": GetCSTtime}).Parse(tpltext.Tpl)
		if err != nil {
			logs.Error(logsign, err.Error())
			message = err.Error()
		} else {
			tpl.Execute(buf, p_json)
			message = SendMessagePrometheusAlert(buf.String(), P_type, P_ddurl, P_wxurl, P_fsurl, P_phone, P_email, logsign)
		}
	} else {
		message = "接口参数缺失！"
		logs.Error(logsign, message)
	}
	c.Data["json"] = message
	c.ServeJSON()
}

func SendMessagePrometheusAlert(message, ptype, pddurl, pwxurl, pfsurl, pphone, email, logsign string) string {
	Title := beego.AppConfig.String("title")
	ret := ""
	model.AlertsFromCounter.WithLabelValues("PrometheusAlert", message, "", "", "").Add(1)
	switch ptype {
	//微信渠道
	case "wx":
		Wxurl := strings.Split(pwxurl, ",")
		for _, url := range Wxurl {
			ret += PostToWeiXin(message, url, logsign)
		}

	//钉钉渠道
	case "dd":
		Ddurl := strings.Split(pddurl, ",")
		for _, url := range Ddurl {
			ret += PostToDingDing(Title+"告警消息", message, url, logsign)
		}

	//飞书渠道
	case "fs":
		Fsurl := strings.Split(pfsurl, ",")
		for _, url := range Fsurl {
			ret += PostToFS(Title+"告警消息", message, url, logsign)
		}

	//腾讯云短信
	case "txdx":
		ret = PostTXmessage(message, pphone, logsign)
	//华为云短信
	case "hwdx":
		ret = ret + PostHWmessage(message, pphone, logsign)
	//百度云短信
	case "bddx":
		ret = ret + PostBDYmessage(message, pphone, logsign)
	//阿里云短信
	case "alydx":
		ret = ret + PostALYmessage(message, pphone, logsign)
	//腾讯云电话
	case "txdh":
		ret = PostTXphonecall(message, pphone, logsign)
	//阿里云电话
	case "alydh":
		ret = ret + PostALYphonecall(message, pphone, logsign)
	//容联云电话
	case "rlydh":
		ret = ret + PostRLYphonecall(message, pphone, logsign)
	//邮件
	case "email":
		ret = ret + SendEmail(message, email, logsign)
	// Telegram
	case "tg":
		ret = ret + SendTG(message, logsign)
	// Workwechat
	case "workwechat":
		ret = ret + SendWorkWechat(message, logsign)
	//异常参数
	default:
		ret = "参数错误"
	}
	return ret
}
