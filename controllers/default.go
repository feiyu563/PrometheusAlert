package controllers

import (
	"PrometheusAlert/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//取到tpl路径
//fmt.Println(filepath.Join(beego.AppPath,"tpl"))

type MainController struct {
	beego.Controller
}

//main page
func (c *MainController) Get() {
	c.Data["IsIndex"] = true
	c.TplName = "index.html"
}

//test page
func (c *MainController) Test() {
	c.Data["IsTest"] = true
	c.TplName = "test.html"
}

//template page
func (c *MainController) Template() {
	c.Data["IsTemplate"] = true
	c.TplName = "template.html"
	Template, err := models.GetAllTpl()
	if err != nil {
		logs.Error(err)
	}
	c.Data["Template"] = Template
}

//template add
func (c *MainController) TemplateAdd() {
	c.Data["IsTemplate"] = true
	c.TplName = "template_add.html"
}
func (c *MainController) AddTpl() {
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
	c.Data["IsTemplate"] = true
	c.TplName = "template_edit.html"
	s_id, _ := strconv.Atoi(c.Input().Get("id"))
	Template, err := models.GetTpl(s_id)
	if err != nil {
		logs.Error(err)
	}
	c.Data["Template"] = Template
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
	s_id, _ := strconv.Atoi(c.Input().Get("id"))
	err := models.DelTpl(s_id)
	if err != nil {
		logs.Error(err)
	}
	c.Redirect("/template", 302)
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
		ret := PostToWeiXin(wxtext, beego.AppConfig.String("wxurl"), logsign)
		c.Data["json"] = ret
	case "dd":
		ddtext := "## [PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n\n" + "#### 测试告警\n\n" + "###### 告警级别：测试\n\n##### PrometheusAlert\n\n" + "![PrometheusAlert](" + beego.AppConfig.String("logourl") + ")"
		ret := PostToDingDing("PrometheusAlert", ddtext, beego.AppConfig.String("ddurl"), logsign)
		c.Data["json"] = ret
	case "fs":
		fstext := "[PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n\n" + "测试告警\n\n" + "告警级别：测试\n\nPrometheusAlert\n\n" + "![PrometheusAlert](" + beego.AppConfig.String("logourl") + ")"
		ret := PostToFS("PrometheusAlert", fstext, beego.AppConfig.String("fsurl"), logsign)
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
		ret := SendWorkWechat(beego.AppConfig.String("WorkWechat_ToUser"),beego.AppConfig.String("WorkWechat_ToParty"), beego.AppConfig.String("WorkWechat_ToTag"),WorkwechatMessage, logsign)
		c.Data["json"] = ret
	case "bddx":
		MobileMessage := "PrometheusAlertCenter测试告警"
		ret := PostBDYmessage(MobileMessage, beego.AppConfig.String("defaultphone"), logsign)
		c.Data["json"] = ret
	case "bdrl":
		RLMessage := "## [PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n\n" + "#### 测试告警\n\n" + "###### 告警级别：测试\n\n##### PrometheusAlert\n\n" + "![PrometheusAlert](" + beego.AppConfig.String("logourl") + ")"
		ret := PostToRuLiu(beego.AppConfig.String("BDRL_ID"),RLMessage,beego.AppConfig.String("BDRL_URL"), logsign)
		c.Data["json"] = ret
	default:
		c.Data["json"] = "hahaha!"
	}
	c.ServeJSON()
}
