package main

import (
	"PrometheusAlert/model"
	"PrometheusAlert/models"
	_ "PrometheusAlert/routers"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os"
	"path"
)

func init() {
	// 检查数据库文件
	Db_name := "./db/PrometheusAlertDB.db"
	if !com.IsExist(Db_name) {
		os.MkdirAll(path.Dir(Db_name), os.ModePerm)
		os.Create(Db_name)
	}
	// 注册模型
	orm.RegisterModel(new(models.PrometheusAlertDB))
	// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	// 注册默认数据库
	orm.RegisterDataBase("default", "sqlite3", Db_name, 10)
	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true
	logtype := beego.AppConfig.String("logtype")
	if logtype == "console" {
		logs.SetLogger(logtype)
	} else if logtype == "file" {
		logpath := beego.AppConfig.String("logpath")
		logs.SetLogger(logtype, `{"filename":"`+logpath+`"}`)
	}
	model.MetricsInit()
	beego.Handler("/metrics", promhttp.Handler())
	beego.Run()
}
