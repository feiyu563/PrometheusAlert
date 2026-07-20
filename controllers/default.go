package controllers

import (
	"PrometheusAlert/models"
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
	Rlydx           int `json:"rlydx"`
	Ruliu           int `json:"ruliu"`
	Txdx            int `json:"txdx"`
	Txdh            int `json:"txdh"`
	Webhook         int `json:"webhook"`
	Weixin          int `json:"weixin"`
	Workwechat      int `json:"workwechat"`
	Voice           int `json:"voice"`
	Zabbix          int `json:"zabbix"`
	Grafana         int `json:"grafana"`
	Graylog         int `json:"graylog"`
	Prometheus      int `json:"prometheus"`
	Prometheusalert int `json:"prometheusalert"`
	Aliyun          int `json:"prometheusalert"`
}

var ChartsJson DashboardJson
var PhoneCallMessage = ""
var GlobalAlertRouter []*models.AlertRouter
var GlobalPrometheusAlertTpl []*models.PrometheusAlertDB

//取到tpl路径
//fmt.Println(filepath.Join(beego.AppPath,"tpl"))

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	title := beego.AppConfig.String("title")
	if title == "" {
		title = "PrometheusAlert"
	}
	c.Data["AppTitle"] = title
}

//main page
func (c *MainController) Get() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsIndex"] = true
	c.TplName = "index.html"
	c.Data["IsLogin"] = CheckAccount(c.Ctx)

	stats, err := models.GetDashboardStats()
	if err != nil {
		logs.Error("Failed to get dashboard statistics:", err.Error())
	}
	c.Data["Stats"] = stats
}

// Health returns Hello 200
func (c *MainController) Health() {
	c.Ctx.WriteString("Hello!\n")
}
