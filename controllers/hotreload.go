package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ConfigController struct {
	beego.Controller
}

var configPath = "conf/app.conf"

// hot reload config
func (c *ConfigController) Reload() {
	logsign := "[" + LogsSign() + "]"
	if open := beego.AppConfig.String("open-hotreload"); open != "1" {
		logs.Info(logsign, "[hotreload]", "open-hotreload is 0.")
		c.Ctx.WriteString("open-hotreload is disable.\n")
		return
	}

	// apply config file
	err := beego.LoadAppConfig("ini", configPath)
	if err != nil {
		logs.Error(logsign, "[hotreload]", err.Error())
		c.Ctx.ResponseWriter.WriteHeader(500)
		c.Ctx.WriteString("Error reloading config: " + err.Error() + "\n")
		return
	}

	logs.Info(logsign, "[hotreload]", "Config reloaded successfully.")
	c.Ctx.WriteString("Config reloaded successfully.\n")
}

// Add auth middleware or not?
