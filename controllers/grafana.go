package controllers

import (
	"bufio"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"os"
	"strings"
	"time"
)

type GrafanaController struct {
	beego.Controller
}

var PhoneCallMessage=""
// {"evalMatches":[],"message":"5分钟内申请云服务流量低于100","ruleId":6,"ruleName":"云服务任务成功数量过低","ruleUrl":"http://grafana.haimacloud.com/d/pH9lfnrmk/ias-p3-ji-gao-jing-xiang?fullscreen=true\u0026edit=true\u0026tab=alert\u0026panelId=28\u0026orgId=1","state":"ok","title":"[OK] 云服务任务成功数量过低"}
//{"evalMatches":[{"value":0,"metric":"Count","tags":{}}],"message":"5分钟内申请云服务流量低于100","ruleId":6,"ruleName":"云服务任务成功数量过低","ruleUrl":"http://grafana.haimacloud.com/d/pH9lfnrmk/ias-p3-ji-gao-jing-xiang?fullscreen=true\u0026edit=true\u0026tab=alert\u0026panelId=28\u0026orgId=1","state":"alerting","title":"[Alerting] 云服务任务成功数量过低"}

type Grafana struct {
	Message string `json:"message"`
	RuleName string `json:"ruleName"`
	RuleUrl string `json:"ruleUrl"`
	State string `json:"state"`
}

func (c *GrafanaController) GrafanaPhone() {
	alert:=Grafana{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageGrafana(alert,4)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}

func (c *GrafanaController) GrafanaDingding() {
	alert:=Grafana{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageGrafana(alert,3)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}

//typeid 为0,触发电话告警和钉钉告警, typeid 为1 仅触发dingding告警
func SendMessageGrafana(message Grafana,typeid int)(string)  {
	Title:=beego.AppConfig.String("title")
	Logourl:=beego.AppConfig.String("logourl")
	Defaultphone:=beego.AppConfig.String("defaultphone")
	text:=""
	//MobileMessage:=""

	titleend:=""
	//返回的内容
	returnMessage:=""

	//拨号的手机号码
	phone:=""
	//告警级别定义 0 信息,1 警告,2 一般严重,3 严重,4 灾难
	AlertLevel:=[]string{"信息","警告","一般严重","严重","灾难"}

	//拆分用户号码和告警消息
	fullMessage:=strings.Split(message.Message,"&&ddurl")

	if message.State=="ok" {
		titleend="故障恢复信息"
		text="## ["+Title+"Grafana"+titleend+"]("+message.RuleUrl+")\n\n"+"#### "+message.RuleName+"\n\n"+"###### 告警级别："+AlertLevel[4]+"\n\n"+"###### 开始时间："+time.Now().Format("2006-01-02 15:04:05")+"\n\n"+"##### "+fullMessage[0]+" 已经恢复正常\n\n"+"!["+Title+"]("+Logourl+")"
		//MobileMessage="\n["+Title+"Grafana"+titleend+"]\n"+message.RuleName+"\n"+"告警级别："+AlertLevel[4]+"\n"+"开始时间："+time.Now().Format("2006-01-02 15:04:05")+"\n"+message.Message+" 已经恢复正常"
		PhoneCallMessage=fullMessage[0]+" 已经恢复正常"
	}else {
		titleend="故障告警信息"
		text="## ["+Title+"Grafana"+titleend+"]("+message.RuleUrl+")\n\n"+"#### "+message.RuleName+"\n\n"+"###### 告警级别："+AlertLevel[4]+"\n\n"+"###### 开始时间："+time.Now().Format("2006-01-02 15:04:05")+"\n\n"+"##### "+fullMessage[0]+"\n\n"+"!["+Title+"]("+Logourl+")"
		//MobileMessage="\n["+Title+"Grafana"+titleend+"]\n"+message.RuleName+"\n"+"告警级别："+AlertLevel[4]+"\n"+"开始时间："+time.Now().Format("2006-01-02 15:04:05")+"\n"+message.Message
		PhoneCallMessage=fullMessage[0]
	}

	//根据不同nLevel 优先级告警
	//if (nLevel==Messagelevel) {
	//	return PostTXmessage(MobileMessage,AlerMessage[0].Annotations.Mobile)
	//} else if (nLevel==PhoneCalllevel) {
	//	return PostTXphonecall(PhoneCallMessage,AlerMessage[0].Annotations.Mobile)
	//}
	//return PostToDingDing(Title+titleend,text)

	//拆分ddurl
	url:=beego.AppConfig.String("ddurl")
	//判断发送到默认钉钉 还是多个钉钉
	if len(fullMessage)<2 {
		returnMessage=returnMessage+"PostToDingDing:"+PostToDingDing(Title+titleend,text,url)+"\n"
	} else {
		DD:=strings.Split(fullMessage[1], ",")
		for _,d:=range DD {
			returnMessage=returnMessage+"PostToDingDing:"+PostToDingDing(Title+titleend,text,d)+"\n"
		}
	}
	//取到手机号
	if Defaultphone=="" {
		phone=GetUserPhone(1)
	} else {
		phone=Defaultphone
	}

	//触发电话告警
	if typeid==4 {
		returnMessage=returnMessage+"PostTXphonecall:"+PostTXphonecall(PhoneCallMessage,phone)+"\n"
	}
	return returnMessage
}

//获取用户号码
func GetUserPhone(neednum int) string  {
	//判断是否存在user.csv文件
	Num:=beego.AppConfig.String("backupphone")
	Today:=time.Now()
	//Today:=time.Now().Format("2006年1月2日")
	//判断当前时间是否大于10点,大于10点取当天值班号码,小于10点取前一天值班号码
	DayString:=""
	if time.Now().Hour()>=10 {
		//取当天值班号码
		DayString=Today.Format("2006年1月2日")
	} else {
		//取前一天值班号码
		DayString=Today.AddDate(0,0,-1).Format("2006年1月2日")
	}

	_, err := os.Stat("user.csv")
	if err == nil {
		f, err := os.Open("user.csv")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		rd := bufio.NewReader(f)
		for {
			line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
			if err!=nil {
				log.SetPrefix("[DEBUG 3]")
				log.Println(err.Error())
			}
			if strings.Contains(line,DayString ) {
				x:=strings.Split(line, ",")
				Num=x[neednum]
				break
				}
		}
	}
	return Num
}