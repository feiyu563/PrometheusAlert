package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ZabbixController struct {
	beego.Controller
}

type ZabbixMessage struct{
	ZabbixTarget string `json:"zabbixtarget"` //告警目标
	ZabbixMessage string `json:"zabbixmessage"`  //告警消息
	ZabbixType string `json:"zabbixtype"`     //告警类型
}

//zabbix告警消息入口
func (c *ZabbixController) ZabbixAlert() {
	alert:=ZabbixMessage{}
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageZabbix(alert,logsign)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}

func SendMessageZabbix(message ZabbixMessage,logsign string) (string){
	ret:=""
	switch message.ZabbixType {
	//微信渠道
	case "wx":
		ret=PostToWeiXin(message.ZabbixMessage,message.ZabbixTarget,logsign)
	//钉钉渠道
	case "dd":
		ret=PostToDingDing("Zabbix告警消息",message.ZabbixMessage,message.ZabbixTarget,logsign)
	//飞书渠道
	case "fs":
		ret=PostToFeiShu("Zabbix告警消息",message.ZabbixMessage,message.ZabbixTarget,logsign)
	//短信渠道
	case "dx":
		//腾讯云短信
		ret=PostTXmessage(message.ZabbixMessage,message.ZabbixTarget,logsign)
		//华为云短信
		ret=ret+PostHWmessage(message.ZabbixMessage,message.ZabbixTarget,logsign)
		//阿里云短信
		ret=ret+PostALYmessage(message.ZabbixMessage,message.ZabbixTarget,logsign)
	//电话渠道
	case "dh":
		//腾讯云电话
		ret=PostTXphonecall(message.ZabbixMessage,message.ZabbixTarget,logsign)
		//阿里云电话
		ret=ret+PostALYphonecall(message.ZabbixMessage,message.ZabbixTarget,logsign)
		//容联云电话
		ret=ret+PostRLYphonecall(message.ZabbixMessage,message.ZabbixTarget,logsign)
	//异常参数
	default:
		ret="参数错误"
	}
	return ret
}