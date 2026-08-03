package main

import (
	"PrometheusAlert/models"
	_ "PrometheusAlert/routers"
	"io"
	"os"
	"path"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron/v3"

	"github.com/astaxie/beego"
	beecontext "github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func sqliteDBPath() string {
	return beego.AppConfig.DefaultString("db_path", "./db/PrometheusAlertDB.db")
}

func ensureSQLiteDB(dbName string) error {
	if IsExist(dbName) {
		return nil
	}

	if err := os.MkdirAll(path.Dir(dbName), os.ModePerm); err != nil {
		return err
	}

	const seedDB = "./db/PrometheusAlertDB.db"
	if dbName != seedDB && IsExist(seedDB) {
		source, err := os.Open(seedDB)
		if err != nil {
			return err
		}
		defer source.Close()

		target, err := os.Create(dbName)
		if err != nil {
			return err
		}
		defer target.Close()

		_, err = io.Copy(target, source)
		return err
	}

	file, err := os.Create(dbName)
	if err != nil {
		return err
	}
	return file.Close()
}

func init() {
	dbDriver := beego.AppConfig.String("db_driver")
	switch dbDriver {
	case "sqlite3":
		// 检查数据库文件
		DbName := sqliteDBPath()
		if err := ensureSQLiteDB(DbName); err != nil {
			logs.Error(err)
			return
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
		Db_name := sqliteDBPath()
		if err := ensureSQLiteDB(Db_name); err != nil {
			logs.Error(err)
			return
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
	orm.RegisterModel(new(models.PrometheusAlertDB), new(models.AlertRecord), new(models.AlertRouter), new(models.SysConfig))
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Error(err)
		return
	}
	models.AlertRouterDBInit()
	models.AlertTplDBInit()
	models.InitConfigCache()
	models.HookBeegoAppConfig()
}

type CustomAccessLogFilter struct{}

func (f *CustomAccessLogFilter) Filter(ctx *beecontext.Context) bool {
	return ctx.Request.URL.Path == "/health"
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
	beego.BConfig.Log.AccessLogs = true
	beego.DefaultAccessLogFilter = &CustomAccessLogFilter{}
	beego.Run()
}
