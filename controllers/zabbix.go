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
	ZabbixTarget string `json:"zabbixtarget"`
	ZabbixMessage string `json:"zabbixmessage"`
	ZabbixType string `json:"zabbixtype"`
}

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
	case "wx":
		ret=PostToWeiXin(message.ZabbixMessage,message.ZabbixTarget,logsign)
	case "dd":
		ret=PostToDingDing("Zabbix告警消息",message.ZabbixMessage,message.ZabbixTarget,logsign)
	case "dx":
		ret=PostTXmessage(message.ZabbixMessage,message.ZabbixTarget,logsign)
		ret=ret+PostHWmessage(message.ZabbixMessage,message.ZabbixTarget,logsign)
		ret=ret+PostALYmessage(message.ZabbixMessage,message.ZabbixTarget,logsign)
	case "dh":
		ret=PostTXphonecall(message.ZabbixMessage,message.ZabbixTarget,logsign)
		ret=ret+PostALYphonecall(message.ZabbixMessage,message.ZabbixTarget,logsign)
	default:
		ret="参数错误"
	}
	return ret
}