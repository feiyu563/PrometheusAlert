package controllers

import (
	"PrometheusAlert/model"
	"PrometheusAlert/models"
	"bytes"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
	"text/template"
	"encoding/json"
)

type PrometheusAlertController struct {
	beego.Controller
}

func (c *PrometheusAlertController) PrometheusAlert() {
	logsign:="["+LogsSign()+"]"
	var p_json interface{}
	logs.Debug(logsign,strings.Replace(string(c.Ctx.Input.RequestBody),"\n","",-1))
	json.Unmarshal(c.Ctx.Input.RequestBody,&p_json)
	P_type:=c.Input().Get("type")
	P_tpl:=c.Input().Get("tpl")
	P_ddurl:=c.Input().Get("ddurl")
	P_wxurl:=c.Input().Get("wxurl")
	P_fsurl:=c.Input().Get("fsurl")
	P_phone:=c.Input().Get("phone")
	//get tpl
	message:=""
	if P_tpl!="" && P_type!="" {
		tpltext, err := models.GetTplOne(P_tpl)
		if err!=nil {
			logs.Error(logsign,err)
		}
		buf := new(bytes.Buffer)
		tpl,_:=template.New("").Parse(tpltext.Tpl)
		tpl.Execute(buf,p_json)
		message=SendMessagePrometheusAlert(buf.String(),P_type,P_ddurl,P_wxurl,P_fsurl,P_phone,logsign)
	} else {
		message="接口参数缺失！"
		logs.Error(logsign,message)
	}
	c.Data["json"]=message
	c.ServeJSON()
}

func SendMessagePrometheusAlert(message,ptype,pddurl,pwxurl,pfsurl,pphone,logsign string) (string){
	ret:=""
	model.AlertsFromCounter.WithLabelValues("PrometheusAlert",message,"","","").Add(1)
	switch ptype {
	//微信渠道
	case "wx":
		ret=PostToWeiXin(message,pwxurl,logsign)
	//钉钉渠道
	case "dd":
		ret=PostToDingDing("告警消息",message,pddurl,logsign)
	//飞书渠道
	case "fs":
		ret=PostToFeiShu("告警消息",message,pfsurl,logsign)
	//腾讯云短信
	case "txdx":
		ret=PostTXmessage(message,pphone,logsign)
	//华为云短信
	case "hwdx":
		ret=ret+PostHWmessage(message,pphone,logsign)
	//阿里云短信
	case "alydx":
		ret=ret+PostALYmessage(message,pphone,logsign)
	//腾讯云电话
	case "txdh":
		ret=PostTXphonecall(message,pphone,logsign)
	//阿里云电话
	case "alydh":
		ret=ret+PostALYphonecall(message,pphone,logsign)
	//容联云电话
	case "rlydh":
		ret=ret+PostRLYphonecall(message,pphone,logsign)
	//异常参数
	default:
		ret="参数错误"
	}
	return ret
}