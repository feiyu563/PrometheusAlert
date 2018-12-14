package controllers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type MainController struct {
	beego.Controller
}

type Labels struct{
	Alertname string `json:"alertname"`
	Instance string `json:"instance"`
}
type Annotations struct{
	Description string `json:"description"`
	Summary string `json:"summary"`
}
type Alerts struct {
	Labels Labels `json:"labels"`
	Annotations Annotations `json:"annotations"`
	StartsAt string `json:"startsAt"`
	EndsAt string `json:"endsAt"`
}
type Prometheus struct {
	Status string
	Alerts []Alerts
}

type DDMessage struct {
	Msgtype string `json:"msgtype"`
	Markdown struct{
		Title string `json:"title"`
		Text string `json:"text"`
	} `json:"markdown"`
	At struct{
		AtMobiles []string `json:"atMobiles"`
		IsAtAll bool `json:"isAtAll"`
	} `json:"at"`
}
//main page
func (c *MainController) Get() {
	c.Data["Email"] = "244217140@qq.com"
	c.TplName = "index.tpl"
}
//for prometheus alert
func (c *MainController) PrometheusAlert() {
	//{"receiver":"web\\.hook","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"Node_alert","instance":"192.168.10.5:9100","job":"node1","monitor":"node1","node":"alert"},"annotations":{"description":"If one more node goes down the node will be unavailable","summary":"192.168.10.5:9100 node goes down!(current value: 0.2s)"},"startsAt":"2018-08-01T02:01:44.71271343-04:00","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://localhost.localdomain:9090/graph?g0.expr=node_load1+%3E+0.1\u0026g0.tab=1"}],"groupLabels":{"alertname":"Node_alert"},"commonLabels":{"alertname":"Node_alert","instance":"192.168.10.5:9100","job":"node1","monitor":"node1","node":"alert"},"commonAnnotations":{"description":"If one more node goes down the node will be unavailable","summary":"192.168.10.5:9100 node goes down!(current value: 0.2s)"},"externalURL":"http://localhost.localdomain:9093","version":"4","groupKey":"{}:{alertname=\"Node_alert\"}"}
	alert:=Prometheus{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=PostToDingDingP(alert)
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
func PostToDingDingP(message Prometheus)(string)  {
	Ddurl:=beego.AppConfig.String("ddurl")
	Title:=beego.AppConfig.String("title")
	Alerturl:=beego.AppConfig.String("alerturl")
	Logourl:=beego.AppConfig.String("logourl")
	text:=""
    titleend:=""
	if message.Status=="resolved" {
		text="## ["+Title+"云Prometheus故障恢复信息]("+Alerturl+")\n\n"+"#### "+message.Alerts[0].Labels.Alertname+"\n\n"+"###### 告警级别：故障恢复\n\n"+"###### 开始时间："+GetCSTtime(message.Alerts[0].StartsAt)+"\n\n"+"###### 结束时间："+GetCSTtime(message.Alerts[0].EndsAt)+"\n\n"+"###### 故障主机IP："+message.Alerts[0].Labels.Instance+"\n\n"+"##### "+message.Alerts[0].Annotations.Summary+"\n\n"+"!["+Title+"]("+Logourl+")"
		titleend="云故障恢复信息"
	}else {
		text="## ["+Title+"云Prometheus故障告警信息]("+Alerturl+")\n\n"+"#### "+message.Alerts[0].Labels.Alertname+"\n\n"+"###### 告警级别：故障告警\n\n"+"###### 开始时间："+GetCSTtime(message.Alerts[0].StartsAt)+"\n\n"+"###### 结束时间："+GetCSTtime(message.Alerts[0].EndsAt)+"\n\n"+"###### 故障主机IP："+message.Alerts[0].Labels.Instance+"\n\n"+"##### "+message.Alerts[0].Annotations.Description+"\n\n"+"!["+Title+"]("+Logourl+")"
		titleend="云故障告警信息"
	}

	u := DDMessage{
		Msgtype:"markdown",
		Markdown: struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		}{Title: Title+titleend, Text: text},
		At: struct {
			AtMobiles []string `json:"atMobiles"`
			IsAtAll   bool `json:"isAtAll"`
		}{AtMobiles:[]string{"15395105573"} , IsAtAll:true },
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	log.SetPrefix("[DEBUG 2]")
	log.Println(b)
	//url="http://127.0.0.1:8081"
	tr :=&http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	//res,err := http.Post(Ddurl, "application/json", b)
	//resp, err := http.PostForm(url,url.Values{"key": {"Value"}, "id": {"123"}})
	client := &http.Client{Transport: tr}
	res,err  := client.Post(Ddurl, "application/json", b)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()
	result,err:=ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	return string(result)
}
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
func (c *MainController) GraylogAlert() {
	//{"receiver":"web\\.hook","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"Node_alert","instance":"192.168.10.5:9100","job":"node1","monitor":"node1","node":"alert"},"annotations":{"description":"If one more node goes down the node will be unavailable","summary":"192.168.10.5:9100 node goes down!(current value: 0.2s)"},"startsAt":"2018-08-01T02:01:44.71271343-04:00","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://localhost.localdomain:9090/graph?g0.expr=node_load1+%3E+0.1\u0026g0.tab=1"}],"groupLabels":{"alertname":"Node_alert"},"commonLabels":{"alertname":"Node_alert","instance":"192.168.10.5:9100","job":"node1","monitor":"node1","node":"alert"},"commonAnnotations":{"description":"If one more node goes down the node will be unavailable","summary":"192.168.10.5:9100 node goes down!(current value: 0.2s)"},"externalURL":"http://localhost.localdomain:9093","version":"4","groupKey":"{}:{alertname=\"Node_alert\"}"}
	alert:=Graylog{}
	log.SetPrefix("[DEBUG 1]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &alert)
	c.Data["json"]=PostToDingDingG(alert)
	log.SetPrefix("[DEBUG 3]")
	log.Println(c.Data["json"])
	c.ServeJSON()
}

func PostToDingDingG(message Graylog)(string)  {
	Ddurl:=beego.AppConfig.String("ddurl")
	Title:=beego.AppConfig.String("title")
	Alerturl:=beego.AppConfig.String("alerturl")
	Logourl:=beego.AppConfig.String("logourl")
	text:="## ["+Title+"云Graylog告警信息]("+Alerturl+")\n\n"+"#### "+message.Check_result.Result_description+"\n\n"+"###### 告警名称："+message.Check_result.Triggered_condition.Title+"\n\n"+"###### 告警类型："+message.Check_result.Triggered_condition.Type+"\n\n"+"###### 开始时间："+message.Check_result.Triggered_at+" UTC\n\n"+"###### 持续时间："+strconv.Itoa(message.Check_result.Triggered_condition.Parameters.Time)+"\n\n"+"!["+Title+"]("+Logourl+")"
	u := DDMessage{
		Msgtype:"markdown",
		Markdown: struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		}{Title: Title+"云告警信息", Text: text},
		At: struct {
			AtMobiles []string `json:"atMobiles"`
			IsAtAll   bool `json:"isAtAll"`
		}{AtMobiles:[]string{"15395105573"} , IsAtAll:true },
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	log.SetPrefix("[DEBUG 2]")
	log.Println(b)
	//url="http://127.0.0.1:8081"
	tr :=&http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res,err  := client.Post(Ddurl, "application/json", b)
	//res,err := http.Post(Ddurl, "application/json", b)
	//resp, err := http.PostForm(url,url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()
	result,err:=ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	return string(result)
}