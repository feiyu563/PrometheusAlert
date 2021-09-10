package controllers

import (
	"PrometheusAlert/model"
	"PrometheusAlert/models"
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type PrometheusAlertController struct {
	beego.Controller
}

/*
准备新增阿里云告警回调
Content-Type: application/x-www-form-urlencoded; charset=UTF-8

expression=$Average>=95
&metricName=Host.mem.usedutilization
&instanceName=instance-name-****
&signature=eEq1zHuCUp0XSmLD8p8VtTKF****
&metricProject=acs_ecs
&userId=12****
&curValue=97.39
&alertName=基础监控-ECS-内存使用率
&namespace=acs_ecs
&triggerLevel=WARN
&alertState=ALERT
&preTriggerLevel=WARN
&ruleId=applyTemplateee147e59-664f-4033-a1be-e9595746****
&dimensions={userId=12****), instanceId=i-12****}
&timestamp=1508136760
*/
type AliyunAlert struct {
	Expression      string `json:"expression"`
	MetricName      string `json:"metricName"`
	InstanceName    string `json:"instanceName"`
	Signature       string `json:"signature"`
	MetricProject   string `json:"metricProject"`
	UserId          string `json:"userId"`
	CurValue        string `json:"curValue"`
	AlertName       string `json:"alertName"`
	Namespace       string `json:"namespace"`
	TriggerLevel    string `json:"triggerLevel"`
	AlertState      string `json:"alertState"`
	PreTriggerLevel string `json:"preTriggerLevel"`
	RuleId          string `json:"ruleId"`
	Dimensions      string `json:"dimensions"`
	Timestamp       string `json:"timestamp"`
}

func (c *PrometheusAlertController) PrometheusAlert() {
	logsign := "[" + LogsSign() + "]"
	var p_json interface{}
	logs.Debug(logsign, strings.Replace(string(c.Ctx.Input.RequestBody), "\n", "", -1))
	if c.Input().Get("from") == "aliyun" {
		//阿里云云监控告警消息处理
		AliyunAlertJson := AliyunAlert{}
		AliyunAlertJson.Expression = c.Input().Get("expression")
		AliyunAlertJson.MetricName = c.Input().Get("metricName")
		AliyunAlertJson.InstanceName = c.Input().Get("instanceName")
		AliyunAlertJson.Signature = c.Input().Get("signature")
		AliyunAlertJson.MetricProject = c.Input().Get("metricProject")
		AliyunAlertJson.UserId = c.Input().Get("userId")
		AliyunAlertJson.CurValue = c.Input().Get("curValue")
		AliyunAlertJson.AlertName = c.Input().Get("alertName")
		AliyunAlertJson.Namespace = c.Input().Get("namespace")
		AliyunAlertJson.TriggerLevel = c.Input().Get("triggerLevel")
		AliyunAlertJson.AlertState = c.Input().Get("alertState")
		AliyunAlertJson.PreTriggerLevel = c.Input().Get("preTriggerLevel")
		AliyunAlertJson.RuleId = c.Input().Get("ruleId")
		AliyunAlertJson.Dimensions = c.Input().Get("dimensions")
		AliyunAlertJson.Timestamp = c.Input().Get("timestamp")
		p_json = AliyunAlertJson
	} else {
		json.Unmarshal(c.Ctx.Input.RequestBody, &p_json)
	}
	P_type := c.Input().Get("type")
	P_tpl := c.Input().Get("tpl")
	P_ddurl := c.Input().Get("ddurl")
	if P_ddurl == "" {
		P_ddurl = beego.AppConfig.String("ddurl")
	}
	P_wxurl := c.Input().Get("wxurl")
	if P_wxurl == "" {
		P_wxurl = beego.AppConfig.String("wxurl")
	}
	P_fsurl := c.Input().Get("fsurl")
	if P_fsurl == "" {
		P_fsurl = beego.AppConfig.String("fsurl")
	}
	P_webhookurl := c.Input().Get("webhookurl")
	P_phone := c.Input().Get("phone")
	if P_phone == "" && (P_type == "txdx" || P_type == "hwdx" || P_type == "bddx" || P_type == "alydx" || P_type == "txdh" || P_type == "alydh" || P_type == "rlydh" || P_type == "7moordx" || P_type == "7moordh") {
		P_phone = GetUserPhone(1)
	}
	P_email := c.Input().Get("email")
	if P_email == "" {
		P_email = beego.AppConfig.String("Default_emails")
	}
	P_touser := c.Input().Get("wxuser")
	if P_touser == "" {
		P_touser = beego.AppConfig.String("WorkWechat_ToUser")
	}
	P_toparty := c.Input().Get("wxparty")
	if P_toparty == "" {
		P_toparty = beego.AppConfig.String("WorkWechat_ToParty")
	}
	P_totag := c.Input().Get("wxtag")
	if P_totag == "" {
		P_totag = beego.AppConfig.String("WorkWechat_ToTag")
	}
	P_groupid := c.Input().Get("groupid")
	if P_groupid == "" {
		P_groupid = beego.AppConfig.String("BDRL_ID")
	}
	P_atsomeone := c.Input().Get("at")
	//get tpl
	message := ""
	funcMap := template.FuncMap{
		"GetCSTtime": GetCSTtime,
		"TimeFormat": TimeFormat,
		"GetTime":    GetTime,
	}
	if P_tpl != "" && P_type != "" {
		tpltext, err := models.GetTplOne(P_tpl)
		if err != nil {
			logs.Error(logsign, err)
		}
		buf := new(bytes.Buffer)
		//tpl, err := template.New("").Funcs(template.FuncMap{"GetCSTtime": GetCSTtime}).Parse(tpltext.Tpl)
		tpl, err := template.New("").Funcs(funcMap).Parse(tpltext.Tpl)
		if err != nil {
			logs.Error(logsign, err.Error())
			message = err.Error()
		} else {
			tpl.Execute(buf, p_json)
			message = SendMessagePrometheusAlert(buf.String(), P_type, P_ddurl, P_wxurl, P_fsurl, P_webhookurl, P_phone, P_email, P_touser, P_toparty, P_totag, P_groupid, P_atsomeone, logsign)
		}
	} else {
		message = "接口参数缺失！"
		logs.Error(logsign, message)
	}
	c.Data["json"] = message
	c.ServeJSON()
}

func SendMessagePrometheusAlert(message, ptype, pddurl, pwxurl, pfsurl, pwebhookurl, pphone, email, ptouser, ptoparty, ptotag, pgroupid, patsomeone, logsign string) string {
	Title := beego.AppConfig.String("title")
	ret := ""
	model.AlertsFromCounter.WithLabelValues("PrometheusAlert", message, "", "", "").Add(1)
	switch ptype {
	//微信渠道
	case "wx":
		Wxurl := strings.Split(pwxurl, ",")
		for _, url := range Wxurl {
			ret += PostToWeiXin(message, url, patsomeone, logsign)
		}

	//钉钉渠道
	case "dd":
		Ddurl := strings.Split(pddurl, ",")
		for _, url := range Ddurl {
			ret += PostToDingDing(Title+"告警消息", message, url, patsomeone, logsign)
		}

	//飞书渠道
	case "fs":
		Fsurl := strings.Split(pfsurl, ",")
		for _, url := range Fsurl {
			ret += PostToFS(Title+"告警消息", message, url, patsomeone, logsign)
		}

	//Webhook渠道
	case "webhook":
		Fwebhookurl := strings.Split(pwebhookurl, ",")
		for _, url := range Fwebhookurl {
			ret += PostToWebhook(message, url, logsign)
		}

	//腾讯云短信
	case "txdx":
		ret = PostTXmessage(message, pphone, logsign)
	//华为云短信
	case "hwdx":
		ret = ret + PostHWmessage(message, pphone, logsign)
	//百度云短信
	case "bddx":
		ret = ret + PostBDYmessage(message, pphone, logsign)
	//阿里云短信
	case "alydx":
		ret = ret + PostALYmessage(message, pphone, logsign)
	//腾讯云电话
	case "txdh":
		ret = PostTXphonecall(message, pphone, logsign)
	//阿里云电话
	case "alydh":
		ret = ret + PostALYphonecall(message, pphone, logsign)
	//容联云电话
	case "rlydh":
		ret = ret + PostRLYphonecall(message, pphone, logsign)
	//七陌短信
	case "7moordx":
		ret = ret + Post7MOORmessage(message, pphone, logsign)
	//七陌语音电话
	case "7moordh":
		ret = ret + Post7MOORphonecall(message, pphone, logsign)
	//邮件
	case "email":
		ret = ret + SendEmail(message, email, logsign)
	// Telegram
	case "tg":
		ret = ret + SendTG(message, logsign)
	// Workwechat
	case "workwechat":
		ret = ret + SendWorkWechat(ptouser, ptoparty, ptotag, message, logsign)
	//百度Hi(如流)
	case "rl":
		ret += PostToRuLiu(pgroupid, message, beego.AppConfig.String("BDRL_URL"), logsign)
	// Bark
	case "bark":
		ret = ret + SendBark(message, logsign)
	//异常参数
	default:
		ret = "参数错误"
	}
	return ret
}
