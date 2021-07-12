package controllers

import (
	"PrometheusAlert/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/go-gomail/gomail"
	"strings"
)

// SendEmail
func SendEmail(EmailBody, Emails, logsign string) string {
	open := beego.AppConfig.String("open-email")
	if open != "1" {
		logs.Info(logsign, "[email]", "email未配置未开启状态,请先配置open-email为1")
		return "eamil未配置未开启状态,请先配置open-email为1"
	}
	serverHost := beego.AppConfig.String("Email_host")
	serverPort, _ := beego.AppConfig.Int("Email_port")
	fromEmail := beego.AppConfig.String("Email_user")
	Passwd := beego.AppConfig.String("Email_password")
	EmailTitle := beego.AppConfig.String("Email_title")
	//Emails= xxx1@qq.com,xxx2@qq.com,xxx3@qq.com
	SendToEmails := []string{}
	m := gomail.NewMessage()
	if len(Emails) == 0 {
		return "收件人不能为空"
	}
	for _, Email := range strings.Split(Emails, ",") {
		SendToEmails = append(SendToEmails, strings.TrimSpace(Email))
	}
	// 收件人,...代表打散列表填充不定参数
	m.SetHeader("To", SendToEmails...)
	// 发件人
	m.SetAddressHeader("From", fromEmail, EmailTitle)
	// 主题
	m.SetHeader("Subject", EmailTitle)
	// 正文
	m.SetBody("text/html", EmailBody)
	d := gomail.NewDialer(serverHost, serverPort, fromEmail, Passwd)
	// 发送
	err := d.DialAndSend(m)
	model.AlertToCounter.WithLabelValues("email", EmailBody, Emails).Add(1)
	if err != nil {
		logs.Error(logsign, "[email]", err.Error())
	}
	logs.Info(logsign, "[email]", "email send ok to "+Emails)
	return "email send ok to " + Emails
}
