package routers

import (
	"PrometheusAlert/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	//prometheus alert
	beego.Router("/prometheus/alert", &controllers.MainController{},"post:PrometheusAlert")
	beego.Router("/graylog/alert", &controllers.MainController{},"post:GraylogAlert")
}
