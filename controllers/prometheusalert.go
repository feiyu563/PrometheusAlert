package controllers

import (
	"PrometheusAlert/models"
	"PrometheusAlert/models/elastic"
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	tmplhtml "html/template"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
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

//var Xlabels = make(map[string][]string)

func (c *PrometheusAlertController) PrometheusAlert() {
	logsign := "[" + LogsSign() + "]"
	var p_json interface{}
	//针对prometheus的消息特殊处理
	p_alertmanager_json := make(map[string]interface{})
	pMsg := PrometheusAlertMsg{}
	logs.Debug(logsign, strings.Replace(string(c.Ctx.Input.RequestBody), "\n", "", -1))
	if c.Input().Get("from") == "aliyun" {
		models.AlertsFromCounter.WithLabelValues("aliyun").Add(1)
		ChartsJson.Aliyun += 1
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
		//针对prometheus的消息特殊处理
		json.Unmarshal(c.Ctx.Input.RequestBody, &p_alertmanager_json)
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
	//该配置仅适用于alertmanager的消息,用于判断是否需要拆分alertmanager告警消息
	pMsg.Split = c.Input().Get("split")
	//模版加载进内存处理,防止告警过多频繁查库
	var PrometheusAlertTpl *models.PrometheusAlertDB
	if GlobalPrometheusAlertTpl == nil {
		GlobalPrometheusAlertTpl, _ = models.GetAllTpl()
	}
	for _, Tpl := range GlobalPrometheusAlertTpl {
		if Tpl.Tplname == pMsg.Tpl {
			PrometheusAlertTpl = Tpl
		}
	}

	var message string
	if pMsg.Type != "" && PrometheusAlertTpl != nil {
		if pMsg.Split != "false" && PrometheusAlertTpl.Tpluse == "Prometheus" {
			//判断告警路由AlertRouter列表是否为空
			if GlobalAlertRouter == nil {
				//刷新告警路由AlertRouter
				GlobalAlertRouter, _ = models.GetAllAlertRouter()
			}
			Alerts_Value, _ := p_alertmanager_json["alerts"].([]interface{})
			//拆分告警消息
			for _, AlertValue := range Alerts_Value {
				p_alertmanager_json["alerts"] = Alerts_Value[0:0]
				p_alertmanager_json["alerts"] = append(p_alertmanager_json["alerts"].([]interface{}), AlertValue)
				go SetRecord(AlertValue)
				//提取 prometheus 告警消息中的 label，用于和告警路由比对
				xalert := AlertValue.(map[string]interface{})
				//路由处理,可能存在多个路由都匹配成功，所以这里返回的是个列表sMsg
				pMsgs := AlertRouterSet(xalert, pMsg)

				for _, send_msg := range pMsgs {
					//发送消息
					err, msg := TransformAlertMessage(p_alertmanager_json, &send_msg, PrometheusAlertTpl.Tpl, logsign)
					if err != nil {
						logs.Error(logsign, err.Error())
						message = err.Error()
					} else {
						message = msg
					}
				}

			}
		} else {
			err, msg := TransformAlertMessage(p_json, &pMsg, PrometheusAlertTpl.Tpl, logsign)
			if err != nil {
				logs.Error(logsign, err.Error())
				message = err.Error()
			} else {
				message = msg
			}
		}

	} else {
		message = "自定义模板接口参数异常！"
		logs.Error(logsign, message)
	}
	c.Data["json"] = message
	c.ServeJSON()
}

//路由处理
func AlertRouterSet(xalert map[string]interface{}, PMsg PrometheusAlertMsg) []PrometheusAlertMsg {
	return_Msgs := []PrometheusAlertMsg{}
	//循环检测现有的路由规则，找到匹配的目标后，替换发送目标参数
	for _, router_value := range GlobalAlertRouter {
		//先判断是否需要进行路由处理
		if PMsg.Type == router_value.Tpl.Tpltype {
			rules := strings.Split(router_value.Rules, ",")
			rules_num := len(rules)
			rules_num_match := 0

			for _, rule := range rules {
				for label_key, label_value := range xalert["labels"].(map[string]interface{}) {
					if rule == (label_key + "=" + label_value.(string)) {
						rules_num_match += 1
					}
				}
			}

			//判断如果路由规则匹配，需要替换url到现有的参数中
			if rules_num == rules_num_match {
				switch router_value.Tpl.Tpltype {
				case "wx":
					PMsg.Wxurl = router_value.UrlOrPhone
					PMsg.AtSomeOne = router_value.AtSomeOne
				//钉钉渠道
				case "dd":
					PMsg.Ddurl = router_value.UrlOrPhone
					PMsg.AtSomeOne = router_value.AtSomeOne
				//飞书渠道
				case "fs":
					PMsg.Fsurl = router_value.UrlOrPhone
					PMsg.AtSomeOne = router_value.AtSomeOne
				//Webhook渠道
				case "webhook":
					PMsg.WebHookUrl = router_value.UrlOrPhone
					PMsg.AtSomeOne = router_value.AtSomeOne
				//邮件
				case "email":
					PMsg.Email = router_value.UrlOrPhone
				//百度Hi(如流)
				case "rl":
					PMsg.GroupId = router_value.UrlOrPhone
				//短信、电话
				case "txdx", "hwdx", "bddx", "alydx", "txdh", "alydh", "rlydh", "7moordx", "7moordh":
					PMsg.Phone = router_value.UrlOrPhone
				//异常参数
				default:
					logs.Info("暂未支持的路由！")
				}

				//匹配路由完成加入返回列表
				return_Msgs = append(return_Msgs, PMsg)
			}

		}
	}

	//如果没有路由匹配，则将传入的PMsg直接加入返回列表，等于是原路返回
	if len(return_Msgs) == 0 {
		return_Msgs = append(return_Msgs, PMsg)
	}
	return return_Msgs
}

func SetRecord(AlertValue interface{}) {
	var Alertname, Level, Instance, Job, Summary, Description string
	xalert := AlertValue.(map[string]interface{})
	for label_key, label_value := range xalert["labels"].(map[string]interface{}) {
		if label_key == "alertname" {
			Alertname = label_value.(string)
		}
		if label_key == "level" {
			Level = label_value.(string)
		}
		if label_key == "instance" {
			Instance = label_value.(string)
		}
		if label_key == "job" {
			Job = label_value.(string)
		}
		//SetXlabels(label_key, label_value.(string))
	}
	for annotation_key, annotation_value := range xalert["annotations"].(map[string]interface{}) {
		if annotation_key == "description" {
			Description = annotation_value.(string)
		}
		if annotation_key == "summary" {
			Summary = annotation_value.(string)
		}
	}

	if beego.AppConfig.String("AlertRecord") == "1" && !models.GetRecordExist(Alertname, Level, Instance, Job, xalert["startsAt"].(string), xalert["endsAt"].(string), Summary, Description, xalert["status"].(string)) {
		models.AddAlertRecord(Alertname,
			Level,
			Instance,
			Job,
			xalert["startsAt"].(string),
			xalert["endsAt"].(string),
			Summary,
			Description,
			xalert["status"].(string))
	}

	// 告警写入ES
	if beego.AppConfig.DefaultString("alert_to_es", "0") == "1" {
		dt := time.Now()
		dty, dtm := dt.Year(), int(dt.Month())
		esIndex := "prometheusalert-" + strconv.Itoa(dty) + strconv.Itoa(dtm)
		alert := &elastic.AlertES{
			Alertname:   Alertname,
			Status:      xalert["status"].(string),
			Instance:    Instance,
			Level:       Level,
			Job:         Job,
			Summary:     Summary,
			Description: Description,
			StartsAt:    xalert["startsAt"].(string),
			EndsAt:      xalert["endsAt"].(string),
			Created:     dt,
		}
		elastic.Insert(esIndex, alert)
	}
}

//func SetXlabels(keys, values string) {
//	if len(Xlabels["labels"]) > 0 {
//		y := 0
//		for _, x := range Xlabels["labels"] {
//			if x == keys {
//				y = 1
//				break
//			}
//		}
//		if y == 0 {
//			Xlabels["labels"] = append(Xlabels["labels"], keys)
//		}
//	} else {
//		Xlabels["labels"] = append(Xlabels["labels"], keys)
//	}
//
//	if len(Xlabels["values"]) > 0 {
//		y := 0
//		for _, x := range Xlabels["values"] {
//			if x == values {
//				y = 1
//				break
//			}
//		}
//		if y == 0 {
//			Xlabels["values"] = append(Xlabels["values"], values)
//		}
//	} else {
//		Xlabels["values"] = append(Xlabels["values"], values)
//	}
//}

//消息模版化并发送告警
func TransformAlertMessage(p_json interface{}, pmsg *PrometheusAlertMsg, tpltext, logsign string) (error error, msg string) {
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

	buf := new(bytes.Buffer)
	tpl, err := template.New("").Funcs(funcMap).Parse(tpltext)
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
	models.AlertsFromCounter.WithLabelValues("/prometheusalert").Add(1)
	ChartsJson.Prometheusalert += 1
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
