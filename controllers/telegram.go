// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package controllers

import (
	"PrometheusAlert/model"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

// SendTG 发送电报消息
func SendTG(msg, logsign string) string {
	open := beego.AppConfig.String("open-tg")
	if open != "1" {
		logs.Info(logsign, "[tg]", "telegram未配置未开启状态,请先配置open-tg为1")
		return "telegram未配置未开启状态,请先配置open-tg为1"
	}
	tgbottoken := beego.AppConfig.String("TG_TOKEN")
	tgmsgmode := beego.AppConfig.String("TG_MODE_CHAN")
	tguserid, _ := beego.AppConfig.Int64("TG_USERID")
	tgchanname := beego.AppConfig.String("TG_CHANNAME")
	tgapi := beego.AppConfig.String("TG_API_PROXY")

	botapi := newBot(tgbottoken, logsign, tgapi)
	var err error
	if tgmsgmode == "0" {
		// 推送给个人
		tgusermsg := tgbotapi.NewMessage(tguserid, msg)
		_, err = botapi.Send(tgusermsg)
	} else {
		// 推送给channel
		if !strings.HasPrefix(tgchanname, "@") {
			tgchanname = fmt.Sprintf("@%v", tgchanname)
		}
		tgchanmsg := tgbotapi.NewMessageToChannel(tgchanname, msg)
		_, err = botapi.Send(tgchanmsg)
	}
	if err != nil {
		logs.Error(logsign, "[tg]", err.Error())
	}
	model.AlertToCounter.WithLabelValues("telegram", "", "").Add(1)
	logs.Info(logsign, "[tg]", "tg send ok.")
	return "tg send ok"
}

func newBot(token, logsign string, api ...string) *tgbotapi.BotAPI {
	endpoint := tgbotapi.APIEndpoint
	if len(api) != 0 && strings.HasPrefix(api[0], "http") {
		endpoint = api[0]
	}
	bot, err := tgbotapi.NewBotAPIWithAPIEndpoint(token, endpoint)
	if err != nil {
		logs.Error(logsign, "[tg]", err)
		return nil
	}
	runmode := beego.AppConfig.String("runmode")
	if runmode == "dev" {
		bot.Debug = true
	}
	return bot
}
