package main

import (
	_ "PrometheusAlert/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

