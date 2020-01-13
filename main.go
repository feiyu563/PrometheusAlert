package main

import (
	_ "PrometheusAlert/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logtype:=beego.AppConfig.String("logtype")
	if logtype=="console" {
		logs.SetLogger(logtype)
	}else if logtype=="file" {
		logpath:=beego.AppConfig.String("logpath")
		logs.SetLogger(logtype, `{"filename":"`+logpath+`"}`)
	}

	beego.Run()
}

