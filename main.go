package main

import (
	"PrometheusAlert/models"
	_ "PrometheusAlert/routers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron/v3"
	"os"
	"path"
	"runtime"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Infos are set at build time use ldflags.
var (
	Version   string
	Revision  string
	BuildUser string
	BuildDate string
	GoVersion = runtime.Version()
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func init() {
	dbDriver := beego.AppConfig.String("db_driver")
	switch dbDriver {
	case "sqlite3":
		// 检查数据库文件
		DbName := "./db/PrometheusAlertDB.db"
		if !IsExist(DbName) {
			err := os.MkdirAll(path.Dir(DbName), os.ModePerm)
			if err != nil {
				logs.Error(err)
				return
			}
			_, err = os.Create(DbName)
			if err != nil {
				logs.Error(err)
				return
			}
		}
		// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
		err := orm.RegisterDriver("db_driver", orm.DRSqlite)
		if err != nil {
			logs.Error(err)
			return
		}
		// 注册默认数据库
		err = orm.RegisterDataBase("default", "sqlite3", DbName, 10)
		if err != nil {
			logs.Error(err)
			return
		}
	case "mysql":
		err := orm.RegisterDriver("mysql", orm.DRMySQL)
		if err != nil {
			logs.Error(err)
			return
		}
		dataSource := beego.AppConfig.String("db_user") + ":" + beego.AppConfig.String("db_password") + "@tcp(" + beego.AppConfig.String("db_host") + ":" + beego.AppConfig.String("db_port") + ")/" + beego.AppConfig.String("db_name") + "?charset=utf8mb4"
		err = orm.RegisterDataBase("default", "mysql", dataSource)
		if err != nil {
			logs.Error(err)
			return
		}
		orm.SetMaxIdleConns("default", 10)
		orm.SetMaxOpenConns("default", 10)
	case "postgres":
		err := orm.RegisterDriver("postgres", orm.DRPostgres)
		if err != nil {
			logs.Error(err)
			return
		}
		dataSource := "user=" + beego.AppConfig.String("db_user") + " password=" + beego.AppConfig.String("db_password") + " dbname=" + beego.AppConfig.String("db_name") + " host=" + beego.AppConfig.String("db_host") + " port=" + beego.AppConfig.String("db_port") + " sslmode=disable"
		err = orm.RegisterDataBase("default", "postgres", dataSource)
		if err != nil {
			logs.Error(err)
			return
		}
		orm.SetMaxIdleConns("default", 10)
		orm.SetMaxOpenConns("default", 10)
	default:
		// 检查数据库文件
		Db_name := "./db/PrometheusAlertDB.db"
		if !IsExist(Db_name) {
			err := os.MkdirAll(path.Dir(Db_name), os.ModePerm)
			if err != nil {
				logs.Error(err)
				return
			}
			_, err = os.Create(Db_name)
			if err != nil {
				logs.Error(err)
				return
			}
		}
		// 注册驱动（“sqlite3” 属于默认注册，此处代码可省略）
		err := orm.RegisterDriver("db_driver", orm.DRSqlite)
		if err != nil {
			logs.Error(err)
			return
		}
		// 注册默认数据库
		err = orm.RegisterDataBase("default", "sqlite3", Db_name, 10)
		if err != nil {
			logs.Error(err)
			return
		}
	}
	// 注册模型
	orm.RegisterModel(new(models.PrometheusAlertDB), new(models.AlertRecord), new(models.AlertRouter))
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Error(err)
		return
	}
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
	// 输出应用信息
	logs.Info("[main] 构建的Go版本: %s", GoVersion)
	logs.Info("[main] 应用当前版本: %s", Version)
	logs.Info("[main] 应用当前提交: %s", Revision)
	logs.Info("[main] 应用构建时间: %s", BuildDate)
	logs.Info("[main] 应用构建用户: %s", BuildUser)

	// 定时删除日志
	RecordLive, _ := beego.AppConfig.Int("RecordLive")
	if RecordLive == 1 {
		RecordLiveDay, _ := beego.AppConfig.Int("RecordLiveDay")
		logs.Info("[main] 告警记录生存周期: %dd", RecordLiveDay)
		c := cron.New(cron.WithSeconds())
		_, err := c.AddFunc("0 0 1 * * *", func() { models.RecordCleanByTime(RecordLiveDay) })
		if err != nil {
			logs.Error(err)
			return
		}
		c.Start()
	}
	models.MetricsInit()
	beego.Handler("/metrics", promhttp.Handler())
	beego.Run()
}
