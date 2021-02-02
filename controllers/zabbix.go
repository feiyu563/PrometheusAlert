package controllers

import (
	"PrometheusAlert/model"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ZabbixController struct {
	beego.Controller
}

type ZabbixMessage struct {
	ZabbixTarget  string `json:"zabbixtarget"`  //告警目标
	ZabbixMessage string `json:"zabbixmessage"` //告警消息
	ZabbixType    string `json:"zabbixtype"`    //告警类型
}

//zabbix告警消息入口
func (c *ZabbixController) ZabbixAlert() {
	alert := ZabbixMessage{}
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"] = SendMessageZabbix(alert, logsign)
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func SendMessageZabbix(message ZabbixMessage, logsign string) string {
	ret := ""
	model.AlertsFromCounter.WithLabelValues("zabbix", message.ZabbixMessage, "", "", "").Add(1)
	switch message.ZabbixType {
	//微信渠道
	case "wx":
		ret = PostToWeiXin(message.ZabbixMessage, message.ZabbixTarget, logsign)
	//钉钉渠道
	case "dd":
		ret = PostToDingDing("Zabbix告警消息", message.ZabbixMessage, message.ZabbixTarget, logsign)
	//飞书v1渠道
	case "fs":
		ret = PostToFS("Zabbix告警消息", message.ZabbixMessage, message.ZabbixTarget, logsign)
	//腾讯云短信
	case "txdx":
		ret = PostTXmessage(message.ZabbixMessage, message.ZabbixTarget, logsign)
	//华为云短信
	case "hwdx":
		ret = ret + PostHWmessage(message.ZabbixMessage, message.ZabbixTarget, logsign)
	//百度云短信
	case "bddx":
		ret = ret + PostBDYmessage(message.ZabbixMessage, message.ZabbixTarget, logsign)
	//阿里云短信
	case "alydx":
		ret = ret + PostALYmessage(message.ZabbixMessage, message.ZabbixTarget, logsign)
	//腾讯云电话
	case "txdh":
		ret = PostTXphonecall(message.ZabbixMessage, message.ZabbixTarget, logsign)
	//阿里云电话
	case "alydh":
		ret = ret + PostALYphonecall(message.ZabbixMessage, message.ZabbixTarget, logsign)
	//容联云电话
	case "rlydh":
		ret = ret + PostRLYphonecall(message.ZabbixMessage, message.ZabbixTarget, logsign)
	//7mo短信
	case "7moordx":
		ret = ret + Post7MOORmessage(message.ZabbixMessage, message.ZabbixTarget, logsign)
	//7mo电话
	case "7moordh":
		ret = ret + Post7MOORphonecall(message.ZabbixMessage, message.ZabbixTarget, logsign)
	//telegram
	case "tg":
		ret = ret + SendTG(message.ZabbixMessage, logsign)
	//workwechat
	case "workwechat":
		ret = ret + SendWorkWechat(beego.AppConfig.String("WorkWechat_ToUser"),beego.AppConfig.String("WorkWechat_ToParty"), beego.AppConfig.String("WorkWechat_ToTag"),message.ZabbixMessage, logsign)
	//百度Hi(如流)
	case "rl":
		ret = PostToRuLiu(beego.AppConfig.String("BDRL_ID"), message.ZabbixMessage, message.ZabbixTarget, logsign)
	//异常参数
	default:
		ret = "参数错误"
	}
	return ret
}
