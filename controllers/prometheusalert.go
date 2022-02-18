package controllers

import (
	"PrometheusAlert/model"
	"PrometheusAlert/models"
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	tmplhtml "html/template"
	"regexp"
	"strings"
	"text/template"
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

type PrometheusAlertMsg struct {
	Tpl        string
	Type       string
	Ddurl      string
	Wxurl      string
	Fsurl      string
	Phone      string
	WebHookUrl string
	ToUser     string
	Email      string
	ToParty    string
	ToTag      string
	GroupId    string
	AtSomeOne  string
	RoundRobin string
	Split      string
}

type AlertManagerMsg struct {
	Receiver          string        `json:"receiver"`
	Status            string        `json:"status"`
	Alerts            []interface{} `json:"alerts"`
	ExternalURL       string        `json:"externalURL"` //alertmanage 返回地址
	GroupLabels       interface{}   `json:"groupLabels"`
	CommonLabels      interface{}   `json:"commonLabels"`
	CommonAnnotations interface{}   `json:"commonAnnotations"`
	Version           string        `json:"version"`
	GroupKey          string        `json:"groupKey"`
	TruncatedAlerts   interface{}   `json:"truncatedAlerts"`
}

func (c *PrometheusAlertController) PrometheusAlert() {
	logsign := "[" + LogsSign() + "]"
	var p_json interface{}
	p_alertmanager_json := AlertManagerMsg{}
	pMsg := PrometheusAlertMsg{}
	logs.Debug(logsign, strings.Replace(string(c.Ctx.Input.RequestBody), "\n", "", -1))
	//该配置仅适用于alertmanager的消息,用于判断是否需要拆分alertmanager告警消息
	pMsg.Split = c.Input().Get("split")

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
		if pMsg.Split == "true" {
			json.Unmarshal(c.Ctx.Input.RequestBody, &p_alertmanager_json)
		}
	}

	pMsg.Type = c.Input().Get("type")
	pMsg.Tpl = c.Input().Get("tpl")
	pMsg.Ddurl = c.Input().Get("ddurl")
	if pMsg.Ddurl == "" {
		pMsg.Ddurl = beego.AppConfig.String("ddurl")
	}
	pMsg.Wxurl = c.Input().Get("wxurl")
	if pMsg.Wxurl == "" {
		pMsg.Wxurl = beego.AppConfig.String("wxurl")
	}
	pMsg.Fsurl = c.Input().Get("fsurl")
	if pMsg.Fsurl == "" {
		pMsg.Fsurl = beego.AppConfig.String("fsurl")
	}
	pMsg.WebHookUrl = c.Input().Get("webhookurl")
	pMsg.Phone = c.Input().Get("phone")
	if pMsg.Phone == "" && (pMsg.Type == "txdx" || pMsg.Type == "hwdx" || pMsg.Type == "bddx" || pMsg.Type == "alydx" || pMsg.Type == "txdh" || pMsg.Type == "alydh" || pMsg.Type == "rlydh" || pMsg.Type == "7moordx" || pMsg.Type == "7moordh") {
		pMsg.Phone = GetUserPhone(1)
	}
	pMsg.Email = c.Input().Get("email")
	if pMsg.Email == "" {
		pMsg.Email = beego.AppConfig.String("Default_emails")
	}
	pMsg.ToUser = c.Input().Get("wxuser")
	if pMsg.ToUser == "" {
		pMsg.ToUser = beego.AppConfig.String("WorkWechat_ToUser")
	}
	pMsg.ToParty = c.Input().Get("wxparty")
	if pMsg.ToParty == "" {
		pMsg.ToParty = beego.AppConfig.String("WorkWechat_ToParty")
	}
	pMsg.ToTag = c.Input().Get("wxtag")
	if pMsg.ToTag == "" {
		pMsg.ToTag = beego.AppConfig.String("WorkWechat_ToTag")
	}
	pMsg.GroupId = c.Input().Get("groupid")
	if pMsg.GroupId == "" {
		pMsg.GroupId = beego.AppConfig.String("BDRL_ID")
	}
	pMsg.AtSomeOne = c.Input().Get("at")
	pMsg.RoundRobin = c.Input().Get("rr")

	var message string
	var err error
	var msg string
	if pMsg.Tpl != "" && pMsg.Type != "" {
		if pMsg.Split == "true" {
			p_alertmanager_json_copy := p_alertmanager_json
			for _, Split_alert := range p_alertmanager_json.Alerts {
				//清空p_alertmanager_json_copy的Alerts
				p_alertmanager_json_copy.Alerts = p_alertmanager_json_copy.Alerts[0:0]
				//重新将拆分后的Alerts插入到p_alertmanager_json_copy
				p_alertmanager_json_copy.Alerts = append(p_alertmanager_json_copy.Alerts, Split_alert)
				//直接使用p_alertmanager_json_copy去匹配自定义模板会出现报错，主要是由于struct的首字母大写问题，故需要定义一个交换用的json interface，用于将p_alertmanager_json_copy转换成interface
				var p_alertmanager_interface interface{}
				p_a_json, _ := json.Marshal(p_alertmanager_json_copy)
				json.Unmarshal(p_a_json, &p_alertmanager_interface)

				err, msg = TransformAlertMessage(p_alertmanager_interface, &pMsg, logsign)
				if err != nil {
					logs.Error(logsign, err.Error())
					message = err.Error()
				} else {
					message = msg
				}
			}
		} else {
			err, msg = TransformAlertMessage(p_json, &pMsg, logsign)
			if err != nil {
				logs.Error(logsign, err.Error())
				message = err.Error()
			} else {
				message = msg
			}
		}

	} else {
		message = "自定义模板接口参数不全！"
		logs.Error(logsign, message)
	}
	c.Data["json"] = message
	c.ServeJSON()
}

//消息模版化并发送告警
func TransformAlertMessage(p_json interface{}, pmsg *PrometheusAlertMsg, logsign string) (error error, msg string) {
	funcMap := template.FuncMap{
		"GetCSTtime": GetCSTtime,
		"TimeFormat": TimeFormat,
		"GetTime":    GetTime,
		"toUpper":    strings.ToUpper,
		"toLower":    strings.ToLower,
		"title":      strings.Title,
		// join is equal to strings.Join but inverts the argument order
		// for easier pipelining in templates.
		"join": func(sep string, s []string) string {
			return strings.Join(s, sep)
		},
		"match": regexp.MatchString,
		"safeHtml": func(text string) tmplhtml.HTML {
			return tmplhtml.HTML(text)
		},
		"reReplaceAll": func(pattern, repl, text string) string {
			re := regexp.MustCompile(pattern)
			return re.ReplaceAllString(text, repl)
		},
		"stringSlice": func(s ...string) []string {
			return s
		},
	}

	tpltext, err := models.GetTplOne(pmsg.Tpl)
	if err != nil {
		return err, ""
	}
	buf := new(bytes.Buffer)

	tpl, err := template.New("").Funcs(funcMap).Parse(tpltext.Tpl)
	if err != nil {
		return err, ""
	}

	err = tpl.Execute(buf, p_json)
	if err != nil {
		return err, ""
	}

	ReturnMsg := SendMessagePrometheusAlert(buf.String(), pmsg, logsign)
	return nil, ReturnMsg
}

func SendMessagePrometheusAlert(message string, pmsg *PrometheusAlertMsg, logsign string) string {
	Title := beego.AppConfig.String("title")
	var ReturnMsg string
	model.AlertsFromCounter.WithLabelValues("PrometheusAlert", message, "", "", "").Add(1)
	switch pmsg.Type {
	//微信渠道
	case "wx":
		Wxurl := strings.Split(pmsg.Wxurl, ",")
		if pmsg.RoundRobin == "true" {
			ReturnMsg += PostToWeiXin(message, DoBalance(Wxurl), pmsg.AtSomeOne, logsign)
		} else {
			for _, url := range Wxurl {
				ReturnMsg += PostToWeiXin(message, url, pmsg.AtSomeOne, logsign)
			}
		}

	//钉钉渠道
	case "dd":
		Ddurl := strings.Split(pmsg.Ddurl, ",")
		if pmsg.RoundRobin == "true" {
			ReturnMsg += PostToDingDing(Title+"告警消息", message, DoBalance(Ddurl), pmsg.AtSomeOne, logsign)
		} else {
			for _, url := range Ddurl {
				ReturnMsg += PostToDingDing(Title+"告警消息", message, url, pmsg.AtSomeOne, logsign)
			}
		}

	//飞书渠道
	case "fs":
		Fsurl := strings.Split(pmsg.Fsurl, ",")
		if pmsg.RoundRobin == "true" {
			ReturnMsg += PostToFS(Title+"告警消息", message, DoBalance(Fsurl), pmsg.AtSomeOne, logsign)
		} else {
			for _, url := range Fsurl {
				ReturnMsg += PostToFS(Title+"告警消息", message, url, pmsg.AtSomeOne, logsign)
			}
		}

	//Webhook渠道
	case "webhook":
		Fwebhookurl := strings.Split(pmsg.WebHookUrl, ",")
		if pmsg.RoundRobin == "true" {
			ReturnMsg += PostToWebhook(message, DoBalance(Fwebhookurl), logsign)
		} else {
			for _, url := range Fwebhookurl {
				ReturnMsg += PostToWebhook(message, url, logsign)
			}
		}

	//腾讯云短信
	case "txdx":
		ReturnMsg += PostTXmessage(message, pmsg.Phone, logsign)
	//华为云短信
	case "hwdx":
		ReturnMsg += PostHWmessage(message, pmsg.Phone, logsign)
	//百度云短信
	case "bddx":
		ReturnMsg += PostBDYmessage(message, pmsg.Phone, logsign)
	//阿里云短信
	case "alydx":
		ReturnMsg += PostALYmessage(message, pmsg.Phone, logsign)
	//腾讯云电话
	case "txdh":
		ReturnMsg += PostTXphonecall(message, pmsg.Phone, logsign)
	//阿里云电话
	case "alydh":
		ReturnMsg += PostALYphonecall(message, pmsg.Phone, logsign)
	//容联云电话
	case "rlydh":
		ReturnMsg += PostRLYphonecall(message, pmsg.Phone, logsign)
	//七陌短信
	case "7moordx":
		ReturnMsg += Post7MOORmessage(message, pmsg.Phone, logsign)
	//七陌语音电话
	case "7moordh":
		ReturnMsg += Post7MOORphonecall(message, pmsg.Phone, logsign)
	//邮件
	case "email":
		ReturnMsg += SendEmail(message, pmsg.Email, logsign)
	// Telegram
	case "tg":
		ReturnMsg += SendTG(message, logsign)
	// Workwechat
	case "workwechat":
		ReturnMsg += SendWorkWechat(pmsg.ToUser, pmsg.ToParty, pmsg.ToTag, message, logsign)
	//百度Hi(如流)
	case "rl":
		ReturnMsg += PostToRuLiu(pmsg.GroupId, message, beego.AppConfig.String("BDRL_URL"), logsign)
	// Bark
	case "bark":
		ReturnMsg += SendBark(message, logsign)
	//异常参数
	default:
		ReturnMsg = "参数错误"
	}
	return ReturnMsg
}
