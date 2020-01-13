package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type Graylog2Controller struct {
	beego.Controller
}
//graylog2告警部分
type Graylog2 struct {
	Check_result Check_result `json:"check_result"`
}
type Check_result struct {
	Result_description string `json:"result_description"`
	Triggered_condition Triggered_condition `json:"triggered_condition"`
	Triggered_at string `json:"triggered_at"`
}
type Triggered_condition struct {
	Type string `json:"type"`
	Title string `json:"title"`
	Parameters Parameters `json:"parameters"`
}
type Parameters struct {
	Time int `json:"time"`
}
//for graylog alert
func (c *Graylog2Controller) GraylogDingding() {
	alert:=Graylog2{}
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,2,logsign)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogWeixin() {
	alert:=Graylog2{}
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,3,logsign)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogTxdx() {
	alert:=Graylog2{}
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,5,logsign)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogHwdx() {
	alert:=Graylog2{}
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,6,logsign)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogTxdh() {
	alert:=Graylog2{}
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,4,logsign)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogALYdx() {
	alert:=Graylog2{}
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,7,logsign)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogALYdh() {
	alert:=Graylog2{}
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,8,logsign)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func SendMessageG(message Graylog2,typeid int,logsign string)(string)  {
	Title:=beego.AppConfig.String("title")
	Alerturl:=beego.AppConfig.String("GraylogAlerturl")
	Logourl:=beego.AppConfig.String("logourl")
	DDtext:="## ["+Title+"Graylog2告警信息]("+Alerturl+")\n\n"+"#### "+message.Check_result.Result_description+"\n\n"+"###### 告警名称："+message.Check_result.Triggered_condition.Title+"\n\n"+"###### 告警类型："+message.Check_result.Triggered_condition.Type+"\n\n"+"###### 开始时间："+message.Check_result.Triggered_at+" \n\n"+"###### 持续时间："+strconv.Itoa(message.Check_result.Triggered_condition.Parameters.Time)+"\n\n"+"!["+Title+"]("+Logourl+")"
	WXtext:="["+Title+"Graylog2告警信息]("+Alerturl+")\n>**"+message.Check_result.Result_description+"**\n>`告警名称:`"+message.Check_result.Triggered_condition.Title+"\n`告警类型:`"+message.Check_result.Triggered_condition.Type+"\n`开始时间:`"+message.Check_result.Triggered_at+" \n`持续时间:`"+strconv.Itoa(message.Check_result.Triggered_condition.Parameters.Time)+"\n"
	PhoneCallMessage=message.Check_result.Result_description
	//触发钉钉
	if typeid==2 {
		ddurl:=beego.AppConfig.String("ddurl")
		PostToDingDing(Title+"告警信息", DDtext, ddurl,logsign)
	}
	//触发微信
	if typeid==3 {
		wxurl:=beego.AppConfig.String("wxurl")
		PostToWeiXin(WXtext, wxurl,logsign)
	}
	//取到手机号
	phone:=GetUserPhone(1)
	//触发电话告警
	if typeid==4 {
		PostTXphonecall(PhoneCallMessage,phone,logsign)
	}
	//触发腾讯云短信告警
	if typeid==5 {
		PostTXmessage(PhoneCallMessage,phone,logsign)
	}
	//触发华为云短信告警
	if typeid==6 {
		PostHWmessage(PhoneCallMessage,phone,logsign)
	}
	//触发阿里云短信告警
	if typeid==7 {
		PostALYmessage(PhoneCallMessage,phone,logsign)
	}
	//触发阿里云电话告警
	if typeid==8 {
		PostALYphonecall(PhoneCallMessage,phone,logsign)
	}
	return "告警消息发送完成."
}