package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
)

type Graylog3Controller struct {
	beego.Controller
}
//graylog3告警部分
type Graylog3 struct {
	title string `json:"event_definition_title"`
	description string `json:"event_definition_description"`
	event alertEvent `json:"event"`
	//backlog alertBacklog `json:"backlog"`
}
type alertEvent struct {
	Timestamp string `json:"timestamp"` //开始时间
	Timestamp_processing string `json:"timestamp_processing"` //开始时间
	Message string `json:"message"`
	Source string `json:"source`
}
//for graylog3 alert
func (c *Graylog3Controller) GraylogDingding() {
	alert:=Graylog3{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG3(alert,2)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog3Controller) GraylogWeixin() {
	alert:=Graylog3{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG3(alert,3)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog3Controller) GraylogTxdx() {
	alert:=Graylog3{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG3(alert,5)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog3Controller) GraylogHwdx() {
	alert:=Graylog3{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG3(alert,6)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog3Controller) GraylogPhone() {
	alert:=Graylog3{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG3(alert,4)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}
func SendMessageG3(message Graylog3,typeid int)(string)  {
	Title:=beego.AppConfig.String("title")
	Logourl:=beego.AppConfig.String("logourl")
	//返回的内容
	returnMessage:=""
	DDtext:="## ["+Title+"Graylog3告警信息]("+message.event.Source+")\n\n"+"#### "+message.description+"\n\n"+"###### 告警名称："+message.title+"\n\n"+"###### 开始时间："+message.event.Timestamp+" UTC\n\n"+"###### 持续时间："+message.event.Timestamp_processing+"\n\n"+"!["+Title+"]("+Logourl+")"
	WXtext:="["+Title+"Graylog3告警信息]("+message.event.Source+")\n>**"+message.description+"**\n>`告警名称:`"+message.title+"\n`开始时间:`"+message.event.Timestamp+" UTC\n`持续时间:`"+message.event.Timestamp_processing+"\n"
	PhoneCallMessage=message.description
	//触发钉钉
	if typeid==2 {
		ddurl:=beego.AppConfig.String("ddurl")
		returnMessage=returnMessage+"PostToDingDing:"+PostToDingDing(Title+"告警信息", DDtext, ddurl)+"\n"
	}
	//触发微信
	if typeid==3 {
		wxurl:=beego.AppConfig.String("wxurl")
		returnMessage=returnMessage+"PostToWeiXin:"+PostToWeiXin(WXtext, wxurl)+"\n"
	}
	//取到手机号
	phone:=GetUserPhone(1)
	//触发电话告警
	if typeid==4 {
		returnMessage=returnMessage+"PostTXphonecall:"+PostTXphonecall(PhoneCallMessage,phone)+"\n"
	}
	//触发腾讯云短信告警
	if typeid==5 {
		returnMessage=returnMessage+"PostTXmessage:"+PostTXmessage(PhoneCallMessage,phone)+"\n"
	}
	//触发华为云短信告警
	if typeid==6 {
		returnMessage=returnMessage+"PostHWmessage:"+PostHWmessage(PhoneCallMessage,phone)+"\n"
	}
	return returnMessage
}