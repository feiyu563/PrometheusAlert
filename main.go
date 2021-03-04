package main

import (
	"PrometheusAlert/model"
	"PrometheusAlert/models"
	_ "PrometheusAlert/routers"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os"
	"path"
)

func init() {
	db_driver:=beego.AppConfig.String("db_driver")
	switch db_driver {
	case "sqlite3":
		// 检查数据库文件
		Db_name := "./db/PrometheusAlertDB.db"
		if !com.IsExist(Db_name) {
			os.MkdirAll(path.Dir(Db_name), os.ModePerm)
			os.Create(Db_name)
		}
		// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
		orm.RegisterDriver("db_driver", orm.DRSqlite)
		// 注册默认数据库
		orm.RegisterDataBase("default", "sqlite3", Db_name, 10)
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("db_user")+":"+beego.AppConfig.String("db_password")+"@tcp("+beego.AppConfig.String("db_host")+":"+beego.AppConfig.String("db_port")+")/"+beego.AppConfig.String("db_name")+"?charset=utf8mb4")
	case "postgres":
		orm.RegisterDriver("postgres", orm.DRPostgres)
		orm.RegisterDataBase("default", "postgres", "user="+beego.AppConfig.String("db_user")+" password="+beego.AppConfig.String("db_password")+" dbname="+beego.AppConfig.String("db_name")+" host="+beego.AppConfig.String("db_host")+" port="+beego.AppConfig.String("db_port")+" sslmode=disable")
	default:
		// 检查数据库文件
		Db_name := "./db/PrometheusAlertDB.db"
		if !com.IsExist(Db_name) {
			os.MkdirAll(path.Dir(Db_name), os.ModePerm)
			os.Create(Db_name)
		}
		// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
		orm.RegisterDriver("db_driver", orm.DRSqlite)
		// 注册默认数据库
		orm.RegisterDataBase("default", "sqlite3", Db_name, 10)
	}
	// 注册模型
	orm.RegisterModel(new(models.PrometheusAlertDB))
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
