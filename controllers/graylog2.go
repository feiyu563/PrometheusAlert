package controllers

import (
	"encoding/json"
	"fmt"
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
	MatchingMessages []MatchingMessage `json:"matching_messages"`
	Result_description string `json:"result_description"`
}
type MatchingMessage struct {
	Index string `json:"index"`
	Message string `json:"message"`
	Fields G2Field `json:"fields"`
	Timestamp string `json:"timestamp"`
}
type G2Field struct {
	Gl2RemoteIp string `json:"gl2_remote_ip"`
	Gl2RemotePort int `json:"gl_2_remote_port"`
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
	fmt.Println(len(message.Check_result.MatchingMessages))
	if len(message.Check_result.MatchingMessages)==0 {
		ddurl:=beego.AppConfig.String("ddurl")
		PostToDingDing(Title+"告警信息", "## ["+Title+"Graylog2告警信息]("+Alerturl+")\n\n"+"#### "+message.Check_result.Result_description+"\n\n"+"!["+Title+"]("+Logourl+")", ddurl,logsign)
		wxurl:=beego.AppConfig.String("wxurl")
		PostToWeiXin("["+Title+"Graylog2告警信息]("+Alerturl+")\n>**"+message.Check_result.Result_description+"**", wxurl,logsign)
		return "告警消息发送完成."
	}
	for _, m := range message.Check_result.MatchingMessages{
		DDtext:="## ["+Title+"Graylog2告警信息]("+Alerturl+")\n\n"+"#### "+message.Check_result.Result_description+"\n\n"+"###### 告警索引："+m.Index+"\n\n"+"###### 开始时间："+m.Timestamp+" \n\n"+"###### 告警主机："+m.Fields.Gl2RemoteIp+":"+strconv.Itoa(m.Fields.Gl2RemotePort)+"\n\n"+"##### "+m.Message+"\n\n"+"!["+Title+"]("+Logourl+")"
		WXtext:="["+Title+"Graylog2告警信息]("+Alerturl+")\n>**"+message.Check_result.Result_description+"**\n>`告警索引:`"+m.Index+"\n`开始时间:`"+m.Timestamp+" \n`告警主机:`"+m.Fields.Gl2RemoteIp+":"+strconv.Itoa(m.Fields.Gl2RemotePort)+"\n**"+m.Message+"**"
		PhoneCallMessage="告警主机:"+m.Fields.Gl2RemoteIp+":"+strconv.Itoa(m.Fields.Gl2RemotePort)+"告警消息:"+m.Message
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