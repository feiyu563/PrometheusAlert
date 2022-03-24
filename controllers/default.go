package controllers

import (
	"PrometheusAlert/models"
	"bytes"
	"encoding/json"
	"strconv"
	"text/template"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type DashboardJson struct {
	Telegram        int `json:"telegram"`
	Smoordx         int `json:"smoordx"`
	Smoordh         int `json:"smoordh"`
	Alydx           int `json:"alydx"`
	Alydh           int `json:"alydh"`
	Bdydx           int `json:"bdydx"`
	Bark            int `json:"bark"`
	Dingding        int `json:"dingding"`
	Email           int `json:"email"`
	Feishu          int `json:"feishu"`
	Hwdx            int `json:"hwdx"`
	Rlydx           int `json:"rlydx"`
	Ruliu           int `json:"ruliu"`
	Txdx            int `json:"txdx"`
	Txdh            int `json:"txdh"`
	Webhook         int `json:"webhook"`
	Weixin          int `json:"weixin"`
	Workwechat      int `json:"workwechat"`
	Zabbix          int `json:"zabbix"`
	Grafana         int `json:"grafana"`
	Graylog         int `json:"graylog"`
	Prometheus      int `json:"prometheus"`
	Prometheusalert int `json:"prometheusalert"`
	Aliyun          int `json:"prometheusalert"`
}

var ChartsJson DashboardJson

//取到tpl路径
//fmt.Println(filepath.Join(beego.AppPath,"tpl"))

type MainController struct {
	beego.Controller
}

//func (c *MainController) GetLabels() {
//	c.Data["json"] = Xlabels
//	c.ServeJSON()
//}

//main page
func (c *MainController) Get() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsIndex"] = true
	c.TplName = "index.html"
	c.Data["ChartsJson"] = ChartsJson
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

// Health returns Hello 200
func (c *MainController) Health() {
	c.Ctx.WriteString("Hello!\n")
}

//router
func (c *MainController) AlertRouter() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsAlertRouter"] = true
	c.TplName = "alertrouter.html"

	AlertRouter, err := models.GetAllAlertRouter()
	if err != nil {
		logs.Error(err)
	}
	c.Data["AlertRouter"] = AlertRouter

	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

//router add
func (c *MainController) RouterAdd() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	Template, err := models.GetPromtheusTpl()
	if err != nil {
		logs.Error(err)
	}
	c.Data["Template"] = Template
	c.Data["IsAlertRouter"] = true
	c.TplName = "alertrouter_add.html"
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

func (c *MainController) AddRouter() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	//获取表单信息
	tid := c.Input().Get("id")
	name := c.Input().Get("name")
	tpl_id := c.Input().Get("tpl_id")
	rules := c.Input().Get("rules")
	purl := c.Input().Get("purl")
	pat := c.Input().Get("pat")
	var err error
	if len(tid) == 0 {
		id, _ := strconv.Atoi(tid)
		tpl_id_int, _ := strconv.Atoi(tpl_id)
		err = models.AddAlertRouter(id, tpl_id_int, name, rules, purl, pat)
	} else {
		id, _ := strconv.Atoi(tid)
		tpl_id_int, _ := strconv.Atoi(tpl_id)
		err = models.UpdateAlertRouter(id, tpl_id_int, name, rules, purl, pat)
	}
	var resp interface{}
	resp = err
	if err != nil {
		resp = err.Error()
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

//router edit
func (c *MainController) RouterEdit() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsAlertRouter"] = true
	c.TplName = "alertrouter_edit.html"
	s_id, _ := strconv.Atoi(c.Input().Get("id"))
	AlertRouter, err := models.GetAlertRouter(s_id)
	if err != nil {
		logs.Error(err)
	}
	c.Data["AlertRouter"] = AlertRouter
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

func (c *MainController) RouterDel() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	s_id, _ := strconv.Atoi(c.Input().Get("id"))
	err := models.DelAlertRouter(s_id)
	if err != nil {
		logs.Error(err)
	}
	c.Redirect("/alertrouter", 302)
}

//test page
func (c *MainController) Test() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsTest"] = true
	c.TplName = "test.html"
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

//Record page
func (c *MainController) Record() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsRecord"] = true
	c.TplName = "record.html"
	Record, err := models.GetAllRecord()
	if err != nil {
		logs.Error(err)
	}
	c.Data["Record"] = Record
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

func (c *MainController) RecordClean() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	models.RecordClean()
	c.Redirect("/record", 302)
}

//template page
func (c *MainController) Template() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsTemplate"] = true
	c.TplName = "template.html"
	Template, err := models.GetAllTpl()
	if err != nil {
		logs.Error(err)
	}
	c.Data["Template"] = Template
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

//template add
func (c *MainController) TemplateAdd() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsTemplate"] = true
	c.TplName = "template_add.html"
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}
func (c *MainController) AddTpl() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	//获取表单信息
	tid := c.Input().Get("id")
	name := c.Input().Get("name")
	t_tpye := c.Input().Get("type")
	t_use := c.Input().Get("use")
	content := c.Input().Get("content")
	var err error
	if len(tid) == 0 {
		id, _ := strconv.Atoi(tid)
		err = models.AddTpl(id, name, t_tpye, t_use, content)
	} else {
		id, _ := strconv.Atoi(tid)
		err = models.UpdateTpl(id, name, t_tpye, t_use, content)
	}
	var resp interface{}
	resp = err
	if err != nil {
		resp = err.Error()
	}
	c.Data["json"] = resp
	c.ServeJSON()
}
func (c *MainController) TemplateEdit() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsTemplate"] = true
	c.TplName = "template_edit.html"
	s_id, _ := strconv.Atoi(c.Input().Get("id"))
	Template, err := models.GetTpl(s_id)
	if err != nil {
		logs.Error(err)
	}
	c.Data["Template"] = Template
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

//func (c *MainController) TemplateTest() {
//	c.Data["IsTemplate"]=true
//	c.TplName = "template_test.html"
//	s_id,_:=strconv.Atoi(c.Input().Get("id"))
//	Template, err := models.GetTpl(s_id)
//	if err != nil {
//		logs.Error(err)
//	}
//	c.Data["Template"] = Template
//}
func (c *MainController) TemplateDel() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	s_id, _ := strconv.Atoi(c.Input().Get("id"))
	err := models.DelTpl(s_id)
	if err != nil {
		logs.Error(err)
	}
	c.Redirect("/template", 302)
}

//markdown test
func (c *MainController) MarkdownTest() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	if c.Ctx.Request.Method == "GET" {
		c.Data["IsMarkDownTest"] = true
		c.TplName = "markdown_test.html"
		c.Data["IsLogin"] = CheckAccount(c.Ctx)
	} else {
		var p_json interface{}
		var resp string
		JsonContent := c.Input().Get("jsoncontent")
		TplContent := c.Input().Get("tplcontent")
		json.Unmarshal([]byte(JsonContent), &p_json)

		funcMap := template.FuncMap{
			"GetCSTtime": GetCSTtime,
			"TimeFormat": TimeFormat,
			"GetTime":    GetTime,
		}
		buf := new(bytes.Buffer)
		tpl, err := template.New("").Funcs(funcMap).Parse(TplContent)
		if err != nil {
			resp = err.Error()
		} else {
			err = tpl.Execute(buf, p_json)
			if err != nil {
				resp = err.Error()
			} else {
				resp = buf.String()
			}
		}
		c.Data["json"] = resp
		c.ServeJSON()
	}

}

func (c *MainController) SetupWeixin() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	if c.Ctx.Request.Method == "GET" {
		//c.Data["IsMarkDownTest"] = true
		c.TplName = "setup_weixin.html"
		c.Data["IsLogin"] = CheckAccount(c.Ctx)
	} else {
		var p_json interface{}
		var resp string
		JsonContent := c.Input().Get("jsoncontent")
		TplContent := c.Input().Get("tplcontent")
		json.Unmarshal([]byte(JsonContent), &p_json)

		funcMap := template.FuncMap{
			"GetCSTtime": GetCSTtime,
			"TimeFormat": TimeFormat,
			"GetTime":    GetTime,
		}
		buf := new(bytes.Buffer)
		tpl, err := template.New("").Funcs(funcMap).Parse(TplContent)
		if err != nil {
			resp = err.Error()
		} else {
			err = tpl.Execute(buf, p_json)
			if err != nil {
				resp = err.Error()
			} else {
				resp = buf.String()
			}
		}
		c.Data["json"] = resp
		c.ServeJSON()
	}

}

func LogsSign() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func (c *MainController) AlertTest() {
	MessageData := c.Input().Get("mtype")
	logsign := "[" + LogsSign() + "]"
	switch MessageData {
	case "wx":
		wxtext := "[PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n>**测试告警**\n>`告警级别:`测试\n**PrometheusAlert**"
		ret := PostToWeiXin(wxtext, beego.AppConfig.String("wxurl"), "jikun.zhang", logsign)
		c.Data["json"] = ret
	case "dd":
		ddtext := "## [PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n\n" + "#### 测试告警\n\n" + "###### 告警级别：测试\n\n##### PrometheusAlert\n\n" + "![PrometheusAlert](" + beego.AppConfig.String("logourl") + ")"
		ret := PostToDingDing("PrometheusAlert", ddtext, beego.AppConfig.String("ddurl"), "15395105573", logsign)
		c.Data["json"] = ret
	case "fs":
		fstext := "[PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n\n" + "测试告警\n\n" + "告警级别：测试\n\nPrometheusAlert\n\n" + "![PrometheusAlert](" + beego.AppConfig.String("logourl") + ")"
		ret := PostToFS("PrometheusAlert", fstext, beego.AppConfig.String("fsurl"), "244217140@qq.com", logsign)
		c.Data["json"] = ret
	case "txdx":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := PostTXmessage(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "txdh":
		ret := PostTXphonecall("PrometheusAlertCenter测试告警", beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "hwdx":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := PostHWmessage(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "alydx":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := PostALYmessage(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "alydh":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := PostALYphonecall(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "rlydh":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := PostRLYphonecall(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "email":
		TestEmailMessage := `
            <h3>PrometheusAlert邮件告警测试</h3>
			欢迎使用<a href ="https://feiyu563.gitee.io">PrometheusAlert</a><br>
			`
		ret := SendEmail(TestEmailMessage, beego.AppConfig.String("Default_emails"), logsign)
		c.Data["json"] = ret
	case "7moordx":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := Post7MOORmessage(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "7moordh":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := Post7MOORphonecall(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "tg":
		TgMessage := "PrometheusAlertCenter测试告警"
		ret := SendTG(TgMessage, logsign)
		c.Data["json"] = ret
	case "workwechat":
		WorkwechatMessage := "[PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n" + "测试告警\n" + "告警级别：测试\nPrometheusAlert\n" + "![PrometheusAlert](" + beego.AppConfig.String("logourl") + ")"
		ret := SendWorkWechat(beego.AppConfig.String("WorkWechat_ToUser"), beego.AppConfig.String("WorkWechat_ToParty"), beego.AppConfig.String("WorkWechat_ToTag"), WorkwechatMessage, logsign)
		c.Data["json"] = ret
	case "bddx":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := PostBDYmessage(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "bdrl":
		RLMessage := "## [PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n\n" + "#### 测试告警\n\n" + "###### 告警级别：测试\n\n##### PrometheusAlert\n\n" + "![PrometheusAlert](" + beego.AppConfig.String("logourl") + ")"
		ret := PostToRuLiu(beego.AppConfig.String("BDRL_ID"), RLMessage, beego.AppConfig.String("BDRL_URL"), logsign)
		c.Data["json"] = ret
	case "bark":
		TgMessage := "PrometheusAlertCenter测试告警"
		ret := SendBark(TgMessage, logsign)
		c.Data["json"] = ret
	default:
		c.Data["json"] = "hahaha!"
	}
	c.ServeJSON()
}
