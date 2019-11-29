package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
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
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,2)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogWeixin() {
	alert:=Graylog2{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,3)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogTxdx() {
	alert:=Graylog2{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,5)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogHwdx() {
	alert:=Graylog2{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,6)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogPhone() {
	alert:=Graylog2{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,4)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}
func SendMessageG(message Graylog2,typeid int)(string)  {
	Title:=beego.AppConfig.String("title")
	Alerturl:=beego.AppConfig.String("GraylogAlerturl")
	Logourl:=beego.AppConfig.String("logourl")
	//返回的内容
	returnMessage:=""
	DDtext:="## ["+Title+"Graylog2告警信息]("+Alerturl+")\n\n"+"#### "+message.Check_result.Result_description+"\n\n"+"###### 告警名称："+message.Check_result.Triggered_condition.Title+"\n\n"+"###### 告警类型："+message.Check_result.Triggered_condition.Type+"\n\n"+"###### 开始时间："+message.Check_result.Triggered_at+" UTC\n\n"+"###### 持续时间："+strconv.Itoa(message.Check_result.Triggered_condition.Parameters.Time)+"\n\n"+"!["+Title+"]("+Logourl+")"
	WXtext:="["+Title+"Graylog2告警信息]("+Alerturl+")\n>**"+message.Check_result.Result_description+"**\n>`告警名称:`"+message.Check_result.Triggered_condition.Title+"\n`告警类型:`"+message.Check_result.Triggered_condition.Type+"\n`开始时间:`"+message.Check_result.Triggered_at+" UTC\n`持续时间:`"+strconv.Itoa(message.Check_result.Triggered_condition.Parameters.Time)+"\n"
	PhoneCallMessage=message.Check_result.Result_description
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