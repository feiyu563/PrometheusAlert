package controllers

import (
	"PrometheusAlert/model"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
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
	ddurl:=c.GetString("ddurl")
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,2,logsign,ddurl,"","","","","","","","")
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogWeixin() {
	alert:=Graylog2{}
	wxurl:=c.GetString("wxurl")
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,3,logsign,"",wxurl,"","","","","","","")
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogTxdx() {
	alert:=Graylog2{}
	phone:=c.GetString("phone")
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,5,logsign,"","","",phone,"","","","","")
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogHwdx() {
	alert:=Graylog2{}
	phone:=c.GetString("phone")
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,6,logsign,"","","","","",phone,"","","")
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogTxdh() {
	alert:=Graylog2{}
	phone:=c.GetString("phone")
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,4,logsign,"","","","",phone,"","","","")
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogALYdx() {
	alert:=Graylog2{}
	phone:=c.GetString("phone")
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,7,logsign,"","","","","","","",phone,"")
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogALYdh() {
	alert:=Graylog2{}
	phone:=c.GetString("phone")
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,8,logsign,"","","","","","","","",phone)
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}
func (c *Graylog2Controller) GraylogRLYdh() {
	alert:=Graylog2{}
	phone:=c.GetString("phone")
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,9,logsign,"","","","","","",phone,"","")
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}

func (c *Graylog2Controller) GraylogFeishu() {
	alert:=Graylog2{}
	fsurl:=c.GetString("fsurl")
	logsign:="["+LogsSign()+"]"
	logs.Info(logsign,string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageG(alert,10,logsign,"","",fsurl,"","","","","","")
	logs.Info(logsign,c.Data["json"])
	c.ServeJSON()
}

//强制转换时间为cst  2019-09-26T15:27:49.644Z
func GetGraylogCSTtime(date string)(string)  {
	T1:=date[0:10] //取日期
	T2:=date[11:23] //取时间
	T3:=T1+" "+T2
	tm2, _ := time.Parse("2006-01-02 15:04:05.000", T3)
	h, _ := time.ParseDuration("-1h")
	tm3:=tm2.Add(-8*h)
	tm:=tm3.Format("2006-01-02 15:04:05.000")
	return tm
}

func SendMessageG(message Graylog2,typeid int,logsign,ddurl,wxurl,fsurl,txdx,txdh,hwdx,rlydh,alydx,alydh string)(string)  {
	Title:=beego.AppConfig.String("title")
	Alerturl:=beego.AppConfig.String("GraylogAlerturl")
	Logourl:=beego.AppConfig.String("logourl")
	if len(message.Check_result.MatchingMessages)==0 {
		model.AlertsFromCounter.WithLabelValues("graylog2",message.Check_result.Result_description,"4","","").Add(1)
		if ddurl=="" {
			ddurl=beego.AppConfig.String("ddurl")
		}
		PostToDingDing(Title+"告警信息", "## ["+Title+"Graylog2告警信息]("+Alerturl+")\n\n"+"#### "+message.Check_result.Result_description+"\n\n"+"!["+Title+"]("+Logourl+")", ddurl,logsign)
		if fsurl=="" {
			fsurl=beego.AppConfig.String("fsurl")
		}
		PostToFeiShu(Title+"告警信息", "["+Title+"Graylog2告警信息]("+Alerturl+")\n\n"+""+message.Check_result.Result_description+"\n\n"+"!["+Title+"]("+Logourl+")", fsurl,logsign)
		if wxurl=="" {
			wxurl=beego.AppConfig.String("wxurl")
		}
		PostToWeiXin("["+Title+"Graylog2告警信息]("+Alerturl+")\n>**"+message.Check_result.Result_description+"**", wxurl,logsign)
		return "告警消息发送完成."
	}
	for _, m := range message.Check_result.MatchingMessages{
		model.AlertsFromCounter.WithLabelValues("graylog2",message.Check_result.Result_description,"4",m.Fields.Gl2RemoteIp,"firing").Add(1)
		DDtext:="## ["+Title+"Graylog2告警信息]("+Alerturl+")\n\n"+"#### "+message.Check_result.Result_description+"\n\n"+"###### 告警索引："+m.Index+"\n\n"+"###### 开始时间："+GetGraylogCSTtime(m.Timestamp)+" \n\n"+"###### 告警主机："+m.Fields.Gl2RemoteIp+":"+strconv.Itoa(m.Fields.Gl2RemotePort)+"\n\n"+"##### "+m.Message+"\n\n"+"!["+Title+"]("+Logourl+")"
		FStext:="["+Title+"Graylog2告警信息]("+Alerturl+")\n\n"+""+message.Check_result.Result_description+"\n\n"+"告警索引："+m.Index+"\n\n"+"开始时间："+GetGraylogCSTtime(m.Timestamp)+" \n\n"+"告警主机："+m.Fields.Gl2RemoteIp+":"+strconv.Itoa(m.Fields.Gl2RemotePort)+"\n\n"+""+m.Message+"\n\n"+"!["+Title+"]("+Logourl+")"
		WXtext:="["+Title+"Graylog2告警信息]("+Alerturl+")\n>**"+message.Check_result.Result_description+"**\n>`告警索引:`"+m.Index+"\n`开始时间:`"+GetGraylogCSTtime(m.Timestamp)+" \n`告警主机:`"+m.Fields.Gl2RemoteIp+":"+strconv.Itoa(m.Fields.Gl2RemotePort)+"\n**"+m.Message+"**"
		PhoneCallMessage="告警主机 "+m.Fields.Gl2RemoteIp+"端口 "+strconv.Itoa(m.Fields.Gl2RemotePort)+"告警消息 "+m.Message
		//触发钉钉
		if typeid==2 {
			if ddurl=="" {
				ddurl=beego.AppConfig.String("ddurl")
			}
			PostToDingDing(Title+"告警信息", DDtext, ddurl,logsign)
		}
		//触发微信
		if typeid==3 {
			if wxurl=="" {
				wxurl=beego.AppConfig.String("wxurl")
			}
			PostToWeiXin(WXtext, wxurl,logsign)
		}
		//出发飞书
		if typeid==10 {
			if fsurl=="" {
				fsurl=beego.AppConfig.String("fsurl")
			}
			PostToFeiShu(Title+"告警信息", FStext, fsurl,logsign)
		}
		//触发电话告警
		if typeid==4 {
			if txdh=="" {
				txdh=GetUserPhone(1)
			}
			PostTXphonecall(PhoneCallMessage,txdh,logsign)
		}
		//触发腾讯云短信告警
		if typeid==5 {
			if txdx=="" {
				txdx=GetUserPhone(1)
			}
			PostTXmessage(PhoneCallMessage,txdx,logsign)
		}
		//触发华为云短信告警
		if typeid==6 {
			if hwdx=="" {
				hwdx=GetUserPhone(1)
			}
			PostHWmessage(PhoneCallMessage,hwdx,logsign)
		}
		//触发阿里云短信告警
		if typeid==7 {
			if alydx=="" {
				alydx=GetUserPhone(1)
			}
			PostALYmessage(PhoneCallMessage,alydx,logsign)
		}
		//触发阿里云电话告警
		if typeid==8 {
			if alydh=="" {
				alydh=GetUserPhone(1)
			}
			PostALYphonecall(PhoneCallMessage,alydh,logsign)
		}
		//触发容联云电话告警
		if typeid==9 {
			if rlydh=="" {
				rlydh=GetUserPhone(1)
			}
			PostRLYphonecall(PhoneCallMessage,rlydh,logsign)
		}

	}
	return "告警消息发送完成."
}