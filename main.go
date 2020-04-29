package main

import (
	"PrometheusAlert/model"
	_ "PrometheusAlert/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logtype:=beego.AppConfig.String("logtype")
	if logtype=="console" {
		logs.SetLogger(logtype)
	}else if logtype=="file" {
		logpath:=beego.AppConfig.String("logpath")
		logs.SetLogger(logtype, `{"filename":"`+logpath+`"}`)
	}
	model.MetricsInit()
	beego.Handler("/metrics", promhttp.Handler())
	beego.Run()
}

