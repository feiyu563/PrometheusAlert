package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type Graylog3Controller struct {
	beego.Controller
}
//graylog3告警部分
type Graylog3 struct {
	Description string `json:"event_definition_description"`
	Event AlertEvent `json:"event"`
	Backlogs []Backlog `json:"backlog"`
}
type AlertEvent struct {
	Source string `json:"source`
}
type Backlog struct {
	Timestamp string `json:"timestamp"` //开始时间
	Index string `json:"index"` //索引
	Message string `json:"message"`
	Fields G3Field `json:"fields`
}
type G3Field struct {
	Gl2RemoteIp string `json:"gl2_remote_ip"`
	Gl2RemotePort int `json:"gl_2_remote_port"`
}


//for graylog3 alert
func (c *Graylog3Controller) GraylogDingding() {
	alert:=Graylog3{}
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG3(alert,2,logsign)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog3Controller) GraylogWeixin() {
	alert:=Graylog3{}
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG3(alert,3,logsign)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog3Controller) GraylogTxdx() {
	alert:=Graylog3{}
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG3(alert,5,logsign)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog3Controller) GraylogHwdx() {
	alert:=Graylog3{}
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG3(alert,6,logsign)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog3Controller) GraylogTxdh() {
	alert:=Graylog3{}
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG3(alert,4,logsign)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog3Controller) GraylogALYdx() {
	alert:=Graylog3{}
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG3(alert,7,logsign)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog3Controller) GraylogALYdh() {
	alert:=Graylog3{}
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG3(alert,8,logsign)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func SendMessageG3(message Graylog3,typeid int,logsign string)(string)  {
	Title:=beego.AppConfig.String("title")
	Logourl:=beego.AppConfig.String("logourl")
	if len(message.Backlogs)==0 {
		ddurl:=beego.AppConfig.String("ddurl")
		PostToDingDing(Title+"告警信息","## ["+Title+"Graylog3告警信息]("+message.Event.Source+")\n\n"+"#### "+message.Description+"\n\n"+"!["+Title+"]("+Logourl+")", ddurl,logsign)
		wxurl:=beego.AppConfig.String("wxurl")
		PostToWeiXin("["+Title+"Graylog3告警信息]("+message.Event.Source+")\n>**"+message.Description+"**", wxurl,logsign)
		return "告警消息发送完成."
	}
	for _, m := range message.Backlogs {
		DDtext:="## ["+Title+"Graylog3告警信息]("+message.Event.Source+")\n\n"+"#### "+message.Description+"\n\n"+"###### 告警索引："+m.Index+"\n\n"+"###### 开始时间："+GetGraylogCSTtime(m.Timestamp)+" \n\n"+"###### 告警主机："+m.Fields.Gl2RemoteIp+":"+strconv.Itoa(m.Fields.Gl2RemotePort)+"\n\n"+"##### "+m.Message+"\n\n"+"!["+Title+"]("+Logourl+")"
		WXtext:="["+Title+"Graylog3告警信息]("+message.Event.Source+")\n>**"+message.Description+"**\n>`告警索引:`"+m.Index+"\n`开始时间:`"+GetGraylogCSTtime(m.Timestamp)+" \n`告警主机:`"+m.Fields.Gl2RemoteIp+":"+strconv.Itoa(m.Fields.Gl2RemotePort)+"\n**"+m.Message+"**"
		PhoneCallMessage="告警主机 "+m.Fields.Gl2RemoteIp+"端口 "+strconv.Itoa(m.Fields.Gl2RemotePort)+"告警消息 "+m.Message
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
	}
	return "告警消息发送完成."
}