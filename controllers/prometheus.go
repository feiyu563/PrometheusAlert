package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

type PrometheusController struct {
	beego.Controller
}

type Labels struct{
	Alertname string `json:"alertname"`
	Instance string `json:"instance"`
}
type Annotations struct{
	Description string `json:"description"`
	Summary string `json:"summary"`
	Level string `json:"level"`  //2019年2月15日 19:03:07 增加告警级别
	Mobile string `json:"mobile"` //2019年2月25日 19:09:23 增加手机号支持
	Ddurl string `json:"ddurl"` //2019年3月12日 20:33:38 增加多个钉钉告警支持
}
type Alerts struct {
	Labels Labels `json:"labels"`
	Annotations Annotations `json:"annotations"`
	StartsAt string `json:"startsAt"`
	EndsAt string `json:"endsAt"`
	GeneratorUrl string `json:"generatorURL"` //prometheus 告警返回地址
}
type Prometheus struct {
	Status string
	Alerts []Alerts
	Externalurl string `json:"externalURL"` //alertmanage 返回地址
}


// 按照 Alert.Level 从大到小排序
type AlerMessages [] Alerts

func (a AlerMessages) Len() int {         // 重写 Len() 方法
	return len(a)
}
func (a AlerMessages) Swap(i, j int){     // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a AlerMessages) Less(i, j int) bool {    // 重写 Less() 方法， 从大到小排序
	return a[j].Annotations.Level < a[i].Annotations.Level
}

//for prometheus alert
//关于告警级别level共有5个级别,0-4,0 信息,1 警告,2 一般严重,3 严重,4 灾难
func (c *PrometheusController) PrometheusAlert() {
	//{"receiver":"web\\.hook","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"Node_alert","instance":"192.168.10.5:9100","job":"node1","monitor":"node1","node":"alert"},"annotations":{"description":"If one more node goes down the node will be unavailable","summary":"192.168.10.5:9100 node goes down!(current value: 0.2s)"},"startsAt":"2018-08-01T02:01:44.71271343-04:00","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://localhost.localdomain:9090/graph?g0.expr=node_load1+%3E+0.1\u0026g0.tab=1"}],"groupLabels":{"alertname":"Node_alert"},"commonLabels":{"alertname":"Node_alert","instance":"192.168.10.5:9100","job":"node1","monitor":"node1","node":"alert"},"commonAnnotations":{"description":"If one more node goes down the node will be unavailable","summary":"192.168.10.5:9100 node goes down!(current value: 0.2s)"},"externalURL":"http://localhost.localdomain:9093","version":"4","groupKey":"{}:{alertname=\"Node_alert\"}"}
	alert:=Prometheus{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageP(alert)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}
func GetCSTtime(date string)(string)  {
	T1:=date[0:10]
	T2:=date[11:19]
	T3:=T1+" "+T2
	tm2, _ := time.Parse("2006-01-02 15:04:05", T3)
	h, _ := time.ParseDuration("-1h")
	tm3:=tm2.Add(-8*h)
	tm:=tm3.Format("2006-01-02 15:04:05")
	return tm
}
func SendMessageP(message Prometheus)(string)  {
	WebhookType:=beego.AppConfig.String("webhook_type")
	Title:=beego.AppConfig.String("title")
	//Alerturl:=beego.AppConfig.String("alerturl")
	Logourl:=beego.AppConfig.String("logourl")
	Messagelevel,_:=beego.AppConfig.Int("messagelevel")
	PhoneCalllevel,_:=beego.AppConfig.Int("phonecalllevel")
	ddtext:=""
	wxtext:=""
	MobileMessage:=""
	PhoneCallMessage:=""
	titleend:=""
	returnMessage:=""
	//对分组消息进行排序
	AlerMessage:=message.Alerts
	sort.Sort(AlerMessages(AlerMessage))
	//nLevel 为第一条告警信息的告警级别
	nLevel,_:=strconv.Atoi(AlerMessage[0].Annotations.Level)
	//告警级别定义 0 信息,1 警告,2 一般严重,3 严重,4 灾难
	AlertLevel:=[]string{"信息","警告","一般严重","严重","灾难"}
    //nowtime:=time.Now()
	if message.Status=="resolved" {
		titleend="故障恢复信息"
		ddtext="## ["+Title+"Prometheus"+titleend+"]("+AlerMessage[0].GeneratorUrl+")\n\n"+"#### ["+AlerMessage[0].Labels.Alertname+"]("+message.Externalurl+")\n\n"+"###### 告警级别："+AlertLevel[nLevel]+"\n\n"+"###### 开始时间："+GetCSTtime(AlerMessage[0].StartsAt)+"\n\n"+"###### 结束时间："+GetCSTtime(AlerMessage[0].EndsAt)+"\n\n"+"###### 故障主机IP："+AlerMessage[0].Labels.Instance+"\n\n"+"##### "+AlerMessage[0].Annotations.Summary+"\n\n"+"!["+Title+"]("+Logourl+")"
		wxtext="["+Title+"Prometheus"+titleend+"]("+AlerMessage[0].GeneratorUrl+")\n>**["+AlerMessage[0].Labels.Alertname+"]("+message.Externalurl+")**\n>`告警级别:`"+AlertLevel[nLevel]+"\n`开始时间:`"+GetCSTtime(AlerMessage[0].StartsAt)+"\n`结束时间:`"+GetCSTtime(AlerMessage[0].EndsAt)+"\n`故障主机IP:`"+AlerMessage[0].Labels.Instance+"\n**"+AlerMessage[0].Annotations.Summary+"**"
		MobileMessage="\n["+Title+"Prometheus"+titleend+"]\n"+AlerMessage[0].Labels.Alertname+"\n"+"告警级别："+AlertLevel[nLevel]+"\n"+"开始时间："+GetCSTtime(AlerMessage[0].StartsAt)+"\n"+"结束时间："+GetCSTtime(AlerMessage[0].EndsAt)+"\n"+"故障主机IP："+AlerMessage[0].Labels.Instance+"\n"+AlerMessage[0].Annotations.Summary
		PhoneCallMessage="故障主机IP："+AlerMessage[0].Labels.Instance+","+AlerMessage[0].Annotations.Summary
	}else {
		titleend="故障告警信息"
		ddtext="## ["+Title+"Prometheus"+titleend+"]("+AlerMessage[0].GeneratorUrl+")\n\n"+"#### ["+AlerMessage[0].Labels.Alertname+"]("+message.Externalurl+")\n\n"+"###### 告警级别："+AlertLevel[nLevel]+"\n\n"+"###### 开始时间："+GetCSTtime(AlerMessage[0].StartsAt)+"\n\n"+"###### 结束时间："+GetCSTtime(AlerMessage[0].EndsAt)+"\n\n"+"###### 故障主机IP："+AlerMessage[0].Labels.Instance+"\n\n"+"##### "+AlerMessage[0].Annotations.Description+"\n\n"+"!["+Title+"]("+Logourl+")"
		wxtext="["+Title+"Prometheus"+titleend+"]("+AlerMessage[0].GeneratorUrl+")\n>**["+AlerMessage[0].Labels.Alertname+"]("+message.Externalurl+")**\n>`告警级别:`"+AlertLevel[nLevel]+"\n`开始时间:`"+GetCSTtime(AlerMessage[0].StartsAt)+"\n`结束时间:`"+GetCSTtime(AlerMessage[0].EndsAt)+"\n`故障主机IP:`"+AlerMessage[0].Labels.Instance+"\n**"+AlerMessage[0].Annotations.Description+"**"
		MobileMessage="\n["+Title+"Prometheus"+titleend+"]\n"+AlerMessage[0].Labels.Alertname+"\n"+"告警级别："+AlertLevel[nLevel]+"\n"+"开始时间："+GetCSTtime(AlerMessage[0].StartsAt)+"\n"+"结束时间："+GetCSTtime(AlerMessage[0].EndsAt)+"\n"+"故障主机IP："+AlerMessage[0].Labels.Instance+"\n"+AlerMessage[0].Annotations.Description
		PhoneCallMessage="故障主机IP："+AlerMessage[0].Labels.Instance+","+AlerMessage[0].Annotations.Description
	}
	if AlerMessage[0].Annotations.Ddurl==""{
		if WebhookType=="0" {
			url:=beego.AppConfig.String("ddurl")
			returnMessage=returnMessage+"PostToDingDing:"+PostToDingDing(Title+titleend,ddtext,url)+"\n"
		}else {
			url:=beego.AppConfig.String("wxurl")
			//(title,text,Alerturl,Logourl,WXurl string)
			returnMessage=returnMessage+"PostToWeiXin:"+PostToWeiXin(wxtext,url)+"\n"
		}

	}else {
		if WebhookType=="0" {
			Ddurl := strings.Split(AlerMessage[0].Annotations.Ddurl, ",")
			for _, url := range Ddurl {
				returnMessage = returnMessage + "PostToDingDing:" + PostToDingDing(Title+titleend, ddtext, url) + "\n"
			}
		}else {
			Ddurl := strings.Split(AlerMessage[0].Annotations.Ddurl, ",")
			for _, url := range Ddurl {
				returnMessage = returnMessage+"PostToWeiXin:"+PostToWeiXin(wxtext,url) + "\n"
			}
		}
	}
	if (nLevel==Messagelevel) {
		returnMessage=returnMessage+"PostTXmessage:"+PostTXmessage(MobileMessage,AlerMessage[0].Annotations.Mobile)+"\n"
	} else if (nLevel==PhoneCalllevel) {
		returnMessage=returnMessage+"PostTXphonecall:"+PostTXphonecall(PhoneCallMessage,AlerMessage[0].Annotations.Mobile)+"\n"
	}
	return returnMessage
}



