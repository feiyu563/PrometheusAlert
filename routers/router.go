package routers

import (
	"PrometheusAlert/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//alert
	beego.Router("/prometheus/alert", &controllers.PrometheusController{},"post:PrometheusAlert")
	beego.Router("/graylog/alert", &controllers.GraylogController{},"post:GraylogAlert")
	beego.Router("/grafana/alert", &controllers.GrafanaController{},"post:GrafanaAlert")
	beego.Router("/tengxun/status", &controllers.TengXunStatusController{},"post:TengXunStatus")
}
