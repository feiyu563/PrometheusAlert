package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
//main page
func (c *MainController) Get() {
	c.Data["Email"] = "244217140@qq.com"
	c.TplName = "index.tpl"
}

func (c *MainController)AlertTest()  {
	MessageData:=c.Input().Get("mtype")
	switch MessageData {
	case "wx":
		wxtext:="[PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n>**测试告警**\n>`告警级别:`测试\n**PrometheusAlert**"
		ret:=PostToWeiXin(wxtext,beego.AppConfig.String("wxurl"))
		c.Data["json"]=ret
	case "dd":
		ddtext:="## [PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)\n\n"+"#### 测试告警\n\n"+"###### 告警级别：测试\n\n##### PrometheusAlert\n\n"+"![PrometheusAlert]("+beego.AppConfig.String("logourl")+")"
	    ret:=PostToDingDing("PrometheusAlert",ddtext,beego.AppConfig.String("ddurl"))
		c.Data["json"]=ret
	case "txdx":
		MobileMessage:="\n[PrometheusAlert]\n测试告警\n"+"告警级别：测试\nPrometheusAlert"
		ret:=PostTXmessage(MobileMessage,beego.AppConfig.String("defaultphone"))
		c.Data["json"]=ret
	case "txdh":
		ret:=PostTXphonecall("PrometheusAlert测试告警",beego.AppConfig.String("defaultphone"))
		c.Data["json"]=ret
	default:
		c.Data["json"]="fuck you!"
	}
	c.ServeJSON()
}