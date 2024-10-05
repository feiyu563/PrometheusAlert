package controllers

import (
	"PrometheusAlert/models"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	tb "gopkg.in/telebot.v3"
)

// SendTG 发送电报消息
func SendTG(msg, logsign string) string {
	open := beego.AppConfig.String("open-tg")
	if open != "1" {
		logs.Info(logsign, "[tg]", "telegram未配置未开启状态,请先配置open-tg为1")
		return "telegram未配置未开启状态,请先配置open-tg为1"
	}
	tgbottoken := beego.AppConfig.String("TG_TOKEN")
	tgId, _ := beego.AppConfig.Int64("TG_SENDID")
	tgapi := beego.AppConfig.String("TG_API_PROXY")
	botapi := newBot(tgbottoken, logsign, tgapi)
	opt := &tb.SendOptions{}
	tgParseMode := beego.AppConfig.String("TG_PARSE_MODE")
	if tgParseMode == "1" {
		// 设置解析模式为Markdown
		opt.ParseMode = tb.ModeMarkdown
	} else if tgParseMode == "2" {
		opt.ParseMode = tb.ModeMarkdownV2
	}
	if _, err := botapi.Send(tb.ChatID(tgId), msg, opt); err != nil {
		logs.Error(logsign, "[tg]", err.Error())
	}
	models.AlertToCounter.WithLabelValues("telegram").Add(1)
	ChartsJson.Telegram += 1
	logs.Info(logsign, "[tg]", "tg send ok.")
	return "tg send ok"
}

func newBot(token, logsign string, api ...string) *tb.Bot {
	tbcfg := tb.Settings{
		Poller: &tb.LongPoller{Timeout: time.Second * 15},
		Token:  token,
	}
	if len(api) > 0 && strings.HasPrefix(api[0], "http") {
		tbcfg.URL = api[0]
	}
	runmode := beego.AppConfig.String("runmode")
	if runmode == "dev" {
		tbcfg.Verbose = true
	}
	// tgParseMode := beego.AppConfig.String("TG_PARSE_MODE")
	// if tgParseMode == "1" {
	// 	tbcfg.ParseMode = tb.ModeMarkdown // 设置解析模式为Markdown
	// }
	bot, err := tb.NewBot(tbcfg)
	if err != nil {
		logs.Error(logsign, "[tg]", err)
		return nil
	}
	return bot
}
