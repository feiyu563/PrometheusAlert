package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"sort"
	"strconv"
	"strings"
)

type PrometheusController struct {
	beego.Controller
}

type Labels struct{
	Alertname string `json:"alertname"`
	Instance string `json:"instance"`
	Level string `json:"level"`  //2019年11月20日 16:03:10更改告警级别定义位置,适配prometheus alertmanager rule
}
type Annotations struct{
	Description string `json:"description"`
	Summary string `json:"summary"`
	//Level string `json:"level"`  //2019年11月20日 16:04:04 删除Annotations level,改用label中的level
	Mobile string `json:"mobile"` //2019年2月25日 19:09:23 增加手机号支持
	Ddurl string `json:"ddurl"` //2019年3月12日 20:33:38 增加多个钉钉告警支持
	Wxurl string `json:"wxurl"` //2019年3月12日 20:33:38 增加多个钉钉告警支持
}
type Alerts struct {
	Status string
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
	return a[j].Labels.Level < a[i].Labels.Level
}

//for prometheus alert
//关于告警级别level共有5个级别,0-4,0 信息,1 警告,2 一般严重,3 严重,4 灾难
func (c *PrometheusController) PrometheusAlert() {
	//{"receiver":"prometheus-alert-center","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"KubePodCrashLooping","app":"kube-state-metrics","container":"emqx-test","instance":"172.28.58.250:8080","job":"kubernetes-service-endpoints","kubernetes_name":"kube-state-metrics","kubernetes_namespace":"monitor","level":"3","namespace":"test-mars","pod":"emqx-test-0"},"annotations":{"description":"Pod CrashLooping,test-mars/emqx-test-0 的容器 emqx-test 10分钟重启了5次","level":"3","timestamp":"@2019-12-06 03:00:00.516 +0000 UTC"},"startsAt":"2019-12-06T02:57:50.516115711Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"https://xprometheus-dev.i-tetris.com/graph?g0.expr=increase%28kube_pod_container_status_restarts_total%7Bnamespace%21~%22.%2Adev.%2A%22%7D%5B10m%5D%29+%3E+5\u0026g0.tab=1"},{"status":"firing","labels":{"alertname":"PodNotReady","app":"kube-state-metrics","class":"basic","container":"emqx-test","infrastructure":"true","instance":"172.28.58.250:8080","job":"kubernetes-service-endpoints","kubernetes_name":"kube-state-metrics","kubernetes_namespace":"monitor","level":"2","namespace":"test-mars","pod":"emqx-test-0"},"annotations":{"description":"test-mars/emqx-test-0 was not ready more than 120s.","level":"2","timestamp":"@2019-12-06 02:59:57.829 +0000 UTC"},"startsAt":"2019-12-06T02:12:17.830216179Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"https://xprometheus-dev.i-tetris.com/graph?g0.expr=kube_pod_container_status_ready+%21%3D+1\u0026g0.tab=1"},{"status":"firing","labels":{"alertname":"PodNotReady","app":"kube-state-metrics","class":"basic","container":"emqx-test","infrastructure":"true","instance":"172.28.58.250:8080","job":"kubernetes-service-endpoints","kubernetes_name":"kube-state-metrics","kubernetes_namespace":"monitor","level":"3","namespace":"test-mars","pod":"emqx-test-0"},"annotations":{"description":"test-mars/emqx-test-0 was not ready more than 300s.","level":"3","timestamp":"@2019-12-06 02:59:42.829 +0000 UTC"},"startsAt":"2019-12-06T02:15:17.830216179Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"https://xprometheus-dev.i-tetris.com/graph?g0.expr=kube_pod_container_status_ready%7Bnamespace%21~%22.%2Adev.%2A%22%7D+%21%3D+1\u0026g0.tab=1"}],"groupLabels":{"instance":"172.28.58.250:8080"},"commonLabels":{"app":"kube-state-metrics","container":"emqx-test","instance":"172.28.58.250:8080","job":"kubernetes-service-endpoints","kubernetes_name":"kube-state-metrics","kubernetes_namespace":"monitor","namespace":"test-mars","pod":"emqx-test-0"},"commonAnnotations":{},"externalURL":"https://xalertmanager-dev.i-tetris.com","version":"4","groupKey":"{}/{job=~\"^(?:.*)$\"}:{instance=\"172.28.58.250:8080\"}"}
	alert:=Prometheus{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageP(alert)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}
//func GetCSTtime(date string)(string)  {
//	T1:=date[0:10]
//	T2:=date[11:19]
//	T3:=T1+" "+T2
//	tm2, _ := time.Parse("2006-01-02 15:04:05", T3)
//	h, _ := time.ParseDuration("-1h")
//	tm3:=tm2.Add(-8*h)
//	tm:=tm3.Format("2006-01-02 15:04:05")
//	return tm
//}
func SendMessageP(message Prometheus)(string)  {
	Title:=beego.AppConfig.String("title")
	Logourl:=beego.AppConfig.String("logourl")
	Messagelevel,_:=beego.AppConfig.Int("messagelevel")
	PhoneCalllevel,_:=beego.AppConfig.Int("phonecalllevel")
	PhoneCallResolved,_:=beego.AppConfig.Int("phonecallresolved")
	Silent,_:=beego.AppConfig.Int("silent")
	var ddtext,wxtext,MobileMessage,PhoneCallMessage,titleend,returnMessage string
	//对分组消息进行排序
	AlerMessage:=message.Alerts
	sort.Sort(AlerMessages(AlerMessage))
	//告警级别定义 0 信息,1 警告,2 一般严重,3 严重,4 灾难
	AlertLevel:=[]string{"信息","警告","一般严重","严重","灾难"}
    //遍历消息
	for _, RMessage := range AlerMessage {
		nLevel,_:=strconv.Atoi(RMessage.Labels.Level)
		if RMessage.Status=="resolved" {
			titleend="故障恢复信息"
			ddtext="## ["+Title+"Prometheus"+titleend+"]("+RMessage.GeneratorUrl+")\n\n"+"#### ["+RMessage.Labels.Alertname+"]("+message.Externalurl+")\n\n"+"###### 告警级别："+AlertLevel[nLevel]+"\n\n"+"###### 开始时间："+RMessage.StartsAt+"\n\n"+"###### 结束时间："+RMessage.EndsAt+"\n\n"+"###### 故障主机IP："+RMessage.Labels.Instance+"\n\n"+"##### "+RMessage.Annotations.Description+"\n\n"+"!["+Title+"]("+Logourl+")"
			wxtext="["+Title+"Prometheus"+titleend+"]("+RMessage.GeneratorUrl+")\n>**["+RMessage.Labels.Alertname+"]("+message.Externalurl+")**\n>`告警级别:`"+AlertLevel[nLevel]+"\n`开始时间:`"+RMessage.StartsAt+"\n`结束时间:`"+RMessage.EndsAt+"\n`故障主机IP:`"+RMessage.Labels.Instance+"\n**"+RMessage.Annotations.Description+"**"
			MobileMessage="\n["+Title+"Prometheus"+titleend+"]\n"+RMessage.Labels.Alertname+"\n"+"告警级别："+AlertLevel[nLevel]+"\n"+"开始时间："+RMessage.StartsAt+"\n"+"结束时间："+RMessage.EndsAt+"\n"+"故障主机IP："+RMessage.Labels.Instance+"\n"+RMessage.Annotations.Description
			PhoneCallMessage="故障主机IP："+RMessage.Labels.Instance+","+RMessage.Annotations.Description+"已经恢复"
		}else {
			titleend="故障告警信息"
			ddtext="## ["+Title+"Prometheus"+titleend+"]("+RMessage.GeneratorUrl+")\n\n"+"#### ["+RMessage.Labels.Alertname+"]("+message.Externalurl+")\n\n"+"###### 告警级别："+AlertLevel[nLevel]+"\n\n"+"###### 开始时间："+RMessage.StartsAt+"\n\n"+"###### 结束时间："+RMessage.EndsAt+"\n\n"+"###### 故障主机IP："+RMessage.Labels.Instance+"\n\n"+"##### "+RMessage.Annotations.Description+"\n\n"+"!["+Title+"]("+Logourl+")"
			wxtext="["+Title+"Prometheus"+titleend+"]("+RMessage.GeneratorUrl+")\n>**["+RMessage.Labels.Alertname+"]("+message.Externalurl+")**\n>`告警级别:`"+AlertLevel[nLevel]+"\n`开始时间:`"+RMessage.StartsAt+"\n`结束时间:`"+RMessage.EndsAt+"\n`故障主机IP:`"+RMessage.Labels.Instance+"\n**"+RMessage.Annotations.Description+"**"
			MobileMessage="\n["+Title+"Prometheus"+titleend+"]\n"+RMessage.Labels.Alertname+"\n"+"告警级别："+AlertLevel[nLevel]+"\n"+"开始时间："+RMessage.StartsAt+"\n"+"结束时间："+RMessage.EndsAt+"\n"+"故障主机IP："+RMessage.Labels.Instance+"\n"+RMessage.Annotations.Description
			PhoneCallMessage="故障主机IP："+RMessage.Labels.Instance+","+RMessage.Annotations.Description
		}
		//发送消息到钉钉
		if RMessage.Annotations.Ddurl==""{
			url:=beego.AppConfig.String("ddurl")
			returnMessage = returnMessage + "PostToDingDing:" + PostToDingDing(Title+titleend, ddtext, url) + "\n"
		}else {
			Ddurl := strings.Split(RMessage.Annotations.Ddurl, ",")
			for _, url := range Ddurl {
				returnMessage = returnMessage + "PostToDingDing:" + PostToDingDing(Title+titleend, ddtext, url) + "\n"
			}
		}
		//发送消息到微信
		if RMessage.Annotations.Wxurl=="" {
			url := beego.AppConfig.String("wxurl")
			returnMessage = returnMessage + "PostToWeiXin:" + PostToWeiXin(wxtext, url) + "\n"
		}else {
			Wxurl := strings.Split(RMessage.Annotations.Wxurl, ",")
			for _, url := range Wxurl {
				returnMessage = returnMessage + "PostToWeiXin:" + PostToWeiXin(wxtext, url) + "\n"
			}
		}
		//发送消息到短信
		if (nLevel==Messagelevel) {
			if RMessage.Annotations.Mobile=="" {
				phone:=GetUserPhone(1)
				returnMessage = returnMessage + "PostTXmessage:" + PostTXmessage(MobileMessage, phone) + "\n"
				returnMessage = returnMessage + "PostTXmessage:" + PostHWmessage(MobileMessage, phone) + "\n"
			}else {
				returnMessage = returnMessage + "PostTXmessage:" + PostTXmessage(MobileMessage, RMessage.Annotations.Mobile) + "\n"
				returnMessage = returnMessage + "PostTXmessage:" + PostHWmessage(MobileMessage, RMessage.Annotations.Mobile) + "\n"
			}
		}
		//发送消息到语音
		if (nLevel==PhoneCalllevel) {
			//判断如果是恢复信息且PhoneCallResolved
			if (RMessage.Status=="resolved" && PhoneCallResolved!=1) {
				returnMessage="告警恢复消息已经关闭\n"
			}else {
				if RMessage.Annotations.Mobile=="" {
					phone:=GetUserPhone(1)
					returnMessage = returnMessage + "PostTXphonecall:" + PostTXphonecall(PhoneCallMessage, phone) + "\n"
				}else {
					returnMessage = returnMessage + "PostTXphonecall:" + PostTXphonecall(PhoneCallMessage, RMessage.Annotations.Mobile) + "\n"
				}
			}
		}
		//告警抑制开启就直接跳出循环
		if Silent==1 {
			break
		}
	}
	return returnMessage
}



