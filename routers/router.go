package routers

import (
	"PrometheusAlert/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/alerttest", &controllers.MainController{},"post:AlertTest")
	//prometheus
	beego.Router("/prometheus/alert", &controllers.PrometheusController{},"post:PrometheusAlert")
    //graylog2
	beego.Router("/graylog2/phone", &controllers.Graylog2Controller{},"post:GraylogPhone")
	beego.Router("/graylog2/dingding", &controllers.Graylog2Controller{},"post:GraylogDingding")
	beego.Router("/graylog2/weixin", &controllers.Graylog2Controller{},"post:GraylogWeixin")
	beego.Router("/graylog2/txdx", &controllers.Graylog2Controller{},"post:GraylogTxdx")
	beego.Router("/graylog2/hwdx", &controllers.Graylog2Controller{},"post:GraylogHwdx")
	//graylog3
	beego.Router("/graylog3/phone", &controllers.Graylog3Controller{},"post:GraylogPhone")
	beego.Router("/graylog3/dingding", &controllers.Graylog3Controller{},"post:GraylogDingding")
	beego.Router("/graylog3/weixin", &controllers.Graylog3Controller{},"post:GraylogWeixin")
	beego.Router("/graylog3/txdx", &controllers.Graylog3Controller{},"post:GraylogTxdx")
	beego.Router("/graylog3/hwdx", &controllers.Graylog3Controller{},"post:GraylogHwdx")
    //grafana
	beego.Router("/grafana/phone", &controllers.GrafanaController{},"post:GrafanaPhone")
	beego.Router("/grafana/dingding", &controllers.GrafanaController{},"post:GrafanaDingding")
	beego.Router("/grafana/weixin", &controllers.GrafanaController{},"post:GrafanaWeixin")
	beego.Router("/grafana/txdx", &controllers.GrafanaController{},"post:GrafanaTxdx")
	beego.Router("/grafana/hwdx", &controllers.GrafanaController{},"post:GrafanaHwdx")
	beego.Router("/tengxun/status", &controllers.TengXunStatusController{},"post:TengXunStatus")
}
