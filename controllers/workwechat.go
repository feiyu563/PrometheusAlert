// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package controllers

import (
	"PrometheusAlert/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ysicing/workwxbot"
)

// SendWorkWechat 发送微信企业应用消息
func SendWorkWechat(touser,toparty,totag,msg, logsign string) string {
	open := beego.AppConfig.String("open-workwechat")
	if open != "1" {
		logs.Info(logsign, "[workwechat]", "workwechat未配置未开启状态,请先配置open-workwechat为1")
		return "workwechat未配置未开启状态,请先配置open-workwechat为1"
	}
	cropid := beego.AppConfig.String("WorkWechat_CropID")
	agentid, _ := beego.AppConfig.Int64("WorkWechat_AgentID")
	agentsecret := beego.AppConfig.String("WorkWechat_AgentSecret")

	//touser := beego.AppConfig.String("WorkWechat_ToUser")
	//toparty := beego.AppConfig.String("WorkWechat_ToParty")
	//totag := beego.AppConfig.String("WorkWechat_ToTag")

	workwxapi := workwxbot.Client{
		CropID:      cropid,
		AgentID:     agentid,
		AgentSecret: agentsecret,
	}
	workwxmsg := workwxbot.Message{
		ToUser:   touser,
		ToParty:  toparty,
		ToTag:    totag,
		MsgType:  "markdown",
		Markdown: workwxbot.Content{Content: msg},
	}
	if err := workwxapi.Send(workwxmsg); err != nil {
		logs.Error(logsign, "[workwechat]", err.Error())
	}
	model.AlertToCounter.WithLabelValues("workwechat", "", "").Add(1)
	logs.Info(logsign, "[workwechat]", "workwechat send ok.")
	return "workwechat send ok"
}

