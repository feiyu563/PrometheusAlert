package controllers

import (
	"PrometheusAlert/model"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/baidubce/bce-sdk-go/services/sms"
	"github.com/baidubce/bce-sdk-go/services/sms/api"
)

func PostBDYmessage(Messages, PhoneNumbers, logsign string) string {
	open := beego.AppConfig.String("open-baidudx")
	if open != "1" {
		logs.Info(logsign, "[bdymessage]", "百度云短信接口未配置未开启状态,请先配置open-baidudx为1")
		return "百度云短信接口未配置未开启状态,请先配置open-baidudx为1"
	}
	// 用户的Access Key ID和Secret Access Key
	ACCESS_KEY_ID, SECRET_ACCESS_KEY := beego.AppConfig.String("BDY_DX_AK"), beego.AppConfig.String("BDY_DX_SK")
	// 用户指定的Endpoint
	ENDPOINT := beego.AppConfig.String("BDY_DX_ENDPOINT")
	// 初始化一个SmsClient
	smsClient, _ := sms.NewClient(ACCESS_KEY_ID, SECRET_ACCESS_KEY, ENDPOINT)
	// 配置不进行重试，默认为Back Off重试
	//smsClient.Config.Retry = bce.NewNoRetryPolicy()
	// 配置连接超时时间为30秒
	smsClient.Config.ConnectionTimeoutInMillis = 30 * 1000
	contentMap := make(map[string]interface{})
	contentMap["code"] = Messages
	mobiles := strings.Split(PhoneNumbers, ",")
	for _, m := range mobiles {
		sendSmsArgs := &api.SendSmsArgs{
			Mobile:      m,
			Template:    beego.AppConfig.String("BDY_DX_TEMPLATE_ID"),
			SignatureId: beego.AppConfig.String("TXY_DX_SIGNATURE_ID"),
			ContentVar:  contentMap,
		}
		result, err := smsClient.SendSms(sendSmsArgs)
		if err != nil {
			logs.Error(logsign, "[bdymessage]", "send sms to %s error, %s", m, err)
		}
		logs.Info(logsign, "[bdymessage]", "send sms success to %s . %s", m, result)
	}
	model.AlertToCounter.WithLabelValues("bdydx", Messages, PhoneNumbers).Add(1)
	return PhoneNumbers + " SendMessages Over."
}
