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
	c.Data["json"]=SendMessageGrafana(alert,2)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}

func (c *GrafanaController) GrafanaWeixin() {
	alert:=Grafana{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageGrafana(alert,3)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}

func (c *GrafanaController) GrafanaTxdx() {
	alert:=Grafana{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageGrafana(alert,5)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}
func (c *GrafanaController) GrafanaHwdx() {
	alert:=Grafana{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=SendMessageGrafana(alert,6)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}



//typeid 为0,触发电话告警和钉钉告警, typeid 为1 仅触发dingding告警
func SendMessageGrafana(message Grafana,typeid int)(string)  {
	Title:=beego.AppConfig.String("title")
	Logourl:=beego.AppConfig.String("logourl")
	DDtext:=""
	WXtext:=""
	titleend:=""
	//返回的内容
	returnMessage:=""
	//告警级别定义 0 信息,1 警告,2 一般严重,3 严重,4 灾难
	AlertLevel:=[]string{"信息","警告","一般严重","严重","灾难"}
	//拆分用户号码和告警消息
	fullMessage:=strings.Split(message.Message,"&&url")
	if message.State=="ok" {
		titleend="故障恢复信息"
		DDtext="## ["+Title+"Grafana"+titleend+"]("+message.RuleUrl+")\n\n#### "+message.RuleName+"\n\n###### 告警级别："+AlertLevel[4]+"\n\n###### 开始时间："+time.Now().Format("2006-01-02 15:04:05")+"\n\n##### "+fullMessage[0]+" 已经恢复正常\n\n"+"!["+Title+"]("+Logourl+")"
		WXtext="["+Title+"Grafana"+titleend+"]("+message.RuleUrl+")\n>**"+message.RuleName+"**\n>`告警级别:`"+AlertLevel[4]+"\n`开始时间:`"+time.Now().Format("2006-01-02 15:04:05")+"\n"+fullMessage[0]+" 已经恢复正常\n"
		PhoneCallMessage=fullMessage[0]+" 已经恢复正常"
	}else {
		titleend="故障告警信息"
		DDtext="## ["+Title+"Grafana"+titleend+"]("+message.RuleUrl+")\n\n"+"#### "+message.RuleName+"\n\n"+"###### 告警级别："+AlertLevel[4]+"\n\n"+"###### 开始时间："+time.Now().Format("2006-01-02 15:04:05")+"\n\n"+"##### "+fullMessage[0]+"\n\n"+"!["+Title+"]("+Logourl+")"
		WXtext="["+Title+"Grafana"+titleend+"]("+message.RuleUrl+")\n>**"+message.RuleName+"**\n>`告警级别:`"+AlertLevel[4]+"\n`开始时间:`"+time.Now().Format("2006-01-02 15:04:05")+"\n"+fullMessage[0]+"\n"
		PhoneCallMessage=fullMessage[0]
	}
	//触发钉钉
	if typeid==2 {
		if len(fullMessage)<2 {
			ddurl:=beego.AppConfig.String("ddurl")
			returnMessage=returnMessage+"PostToDingDing:"+PostToDingDing(Title+titleend,DDtext,ddurl)+"\n"
		} else {
			DD:=strings.Split(fullMessage[1], ",")
			for _,d:=range DD {
				returnMessage=returnMessage+"PostToDingDing:"+PostToDingDing(Title+titleend,DDtext,d)+"\n"
			}
		}
	}
	//触发微信
	if typeid==3 {
		if len(fullMessage)<2 {
			wxurl:=beego.AppConfig.String("wxurl")
			returnMessage=returnMessage+"PostToWeiXin:"+PostToWeiXin(WXtext, wxurl)+"\n"
		} else {
			DD:=strings.Split(fullMessage[1], ",")
			for _,d:=range DD {
				returnMessage=returnMessage+"PostToWeiXin:"+PostToWeiXin(WXtext, d)+"\n"
			}
		}
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
//获取用户号码
func GetUserPhone(neednum int) string  {
	//判断是否存在user.csv文件
	Num:=beego.AppConfig.String("defaultphone")
	Today:=time.Now()
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
                                break
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
