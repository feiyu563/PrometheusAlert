package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"strconv"
	"encoding/json"
)

type GraylogController struct {
	beego.Controller
}

//graylog告警部分
type Graylog struct {
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
func (c *GraylogController) GraylogAlert() {
	//{"receiver":"web\\.hook","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"Node_alert","instance":"192.168.10.5:9100","job":"node1","monitor":"node1","node":"alert"},"annotations":{"description":"If one more node goes down the node will be unavailable","summary":"192.168.10.5:9100 node goes down!(current value: 0.2s)"},"startsAt":"2018-08-01T02:01:44.71271343-04:00","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://localhost.localdomain:9090/graph?g0.expr=node_load1+%3E+0.1\u0026g0.tab=1"}],"groupLabels":{"alertname":"Node_alert"},"commonLabels":{"alertname":"Node_alert","instance":"192.168.10.5:9100","job":"node1","monitor":"node1","node":"alert"},"commonAnnotations":{"description":"If one more node goes down the node will be unavailable","summary":"192.168.10.5:9100 node goes down!(current value: 0.2s)"},"externalURL":"http://localhost.localdomain:9093","version":"4","groupKey":"{}:{alertname=\"Node_alert\"}"}
	alert:=Graylog{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}

func SendMessageG(message Graylog)(string)  {
	WebhookType:=beego.AppConfig.String("webhook_type")
	Title:=beego.AppConfig.String("title")
	Alerturl:=beego.AppConfig.String("alerturl")
	Logourl:=beego.AppConfig.String("logourl")
	DDtext:="## ["+Title+"Graylog告警信息]("+Alerturl+")\n\n"+"#### "+message.Check_result.Result_description+"\n\n"+"###### 告警名称："+message.Check_result.Triggered_condition.Title+"\n\n"+"###### 告警类型："+message.Check_result.Triggered_condition.Type+"\n\n"+"###### 开始时间："+message.Check_result.Triggered_at+" UTC\n\n"+"###### 持续时间："+strconv.Itoa(message.Check_result.Triggered_condition.Parameters.Time)+"\n\n"+"!["+Title+"]("+Logourl+")"
	WXtext:="["+Title+"Graylog告警信息]("+Alerturl+")\n>**"+message.Check_result.Result_description+"**\n>`告警名称:`"+message.Check_result.Triggered_condition.Title+"\n`告警类型:`"+message.Check_result.Triggered_condition.Type+"\n`开始时间:`"+message.Check_result.Triggered_at+" UTC\n`持续时间:`"+strconv.Itoa(message.Check_result.Triggered_condition.Parameters.Time)+"\n"
	if WebhookType=="0" {
		Ddurl:=beego.AppConfig.String("ddurl")
		return PostToDingDing(Title+"告警信息", DDtext, Ddurl)
	}else {
		Wxurl:=beego.AppConfig.String("wxurl")
		return PostToWeiXin(WXtext, Wxurl)
	}
}