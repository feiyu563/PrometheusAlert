package routers

import (
	"PrometheusAlert/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//page
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/template", &controllers.MainController{}, "get:Template")
	beego.Router("/template/add", &controllers.MainController{}, "get:TemplateAdd")
	beego.Router("/template/addtpl", &controllers.MainController{}, "post:AddTpl")
	beego.Router("/template/edit", &controllers.MainController{}, "get,post:TemplateEdit")
	beego.Router("/template/del", &controllers.MainController{}, "get:TemplateDel")
	beego.Router("/template/import", &controllers.MainController{}, "post:ImportTpl")

	beego.Router("/alerttest", &controllers.MainController{}, "post:AlertTest")
	beego.Router("/test", &controllers.MainController{}, "get:Test")
	beego.Router("/markdowntest", &controllers.MainController{}, "get,post:MarkdownTest")

	beego.Router("/record", &controllers.MainController{}, "get:Record")
	beego.Router("/record/data", &controllers.MainController{}, "get:RecordData")
	beego.Router("/record/filters", &controllers.MainController{}, "get:RecordFilters")
	beego.Router("/record/clean", &controllers.MainController{}, "get:RecordClean")

	//alertrouter
	beego.Router("/alertrouter", &controllers.MainController{}, "get:AlertRouter")
	beego.Router("/alertrouter/add", &controllers.MainController{}, "get:RouterAdd")
	beego.Router("/alertrouter/edit", &controllers.MainController{}, "get:RouterEdit")
	beego.Router("/alertrouter/addrouter", &controllers.MainController{}, "post:AddRouter")
	beego.Router("/alertrouter/changestatus", &controllers.MainController{}, "post:RouterChangeStatus")
	beego.Router("/alertrouter/del", &controllers.MainController{}, "get:RouterDel")

	beego.Router("/oncall", &controllers.MainController{}, "get:OnCall")
	beego.Router("/oncall/add", &controllers.MainController{}, "get:OnCallAdd")
	beego.Router("/oncall/edit", &controllers.MainController{}, "get:OnCallEdit")
	beego.Router("/oncall/save", &controllers.MainController{}, "post:SaveOnCall")
	beego.Router("/oncall/del", &controllers.MainController{}, "get:OnCallDel")

	//system config
	beego.Router("/system/config", &controllers.MainController{}, "get:SystemConfig")
	beego.Router("/system/config/save", &controllers.MainController{}, "post:SaveSystemConfig")

	//service log
	beego.Router("/servicelog", &controllers.MainController{}, "get:ServiceLog")
	beego.Router("/servicelog/data", &controllers.MainController{}, "get:ServiceLogData")

	// health
	beego.Router("/health", &controllers.MainController{}, "get:Health")

	beego.Router("/tengxun/status", &controllers.TengXunStatusController{}, "post:TengXunStatus")
	//zabbix
	beego.Router("/zabbix/alert", &controllers.ZabbixController{}, "post:ZabbixAlert")

	//webhook
	beego.Router("/prometheusalert", &controllers.PrometheusAlertController{}, "get,post:PrometheusAlert")

	// gitlab
	beego.Router("/gitlab/weixin", &controllers.GitlabController{}, "post:GitlabWeixin")
	beego.Router("/gitlab/dingding", &controllers.GitlabController{}, "post:GitlabDingding")
	beego.Router("/gitlab/feishu", &controllers.GitlabController{}, "post:GitlabFeishu")

	// hotreload
	beego.Router("/-/reload", &controllers.ConfigController{}, "post:Reload")
}
