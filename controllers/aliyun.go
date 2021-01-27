package controllers

import (
	"PrometheusAlert/model"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dyvmsapi"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
)

func PostALYmessage(Messages, PhoneNumbers, logsign string) string {
	open := beego.AppConfig.String("open-alydx")
	if open != "1" {
		logs.Info(logsign, "[alymessage]", "阿里云短信接口未配置未开启状态,请先配置open-alydx为1")
		return "阿里云短信接口未配置未开启状态,请先配置open-alydx为1"
	}
	AccessKeyId := beego.AppConfig.String("ALY_DX_AccessKeyId")
	AccessSecret := beego.AppConfig.String("ALY_DX_AccessSecret")
	SignName := beego.AppConfig.String("ALY_DX_SignName")
	Template := beego.AppConfig.String("ALY_DX_Template")
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", AccessKeyId, AccessSecret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = PhoneNumbers
	request.SignName = SignName
	request.TemplateCode = Template
	request.TemplateParam = `{"code":"` + Messages + `"}`
	response, err := client.SendSms(request)

	if err != nil {
		logs.Error(logsign, "[alymessage]", err.Error())
	}
	logs.Info(logsign, "[alymessage]", response)
	model.AlertToCounter.WithLabelValues("alydx", Messages, PhoneNumbers).Add(1)
	return response.Message
}
func PostALYphonecall(Messages string, PhoneNumbers, logsign string) string {
	open := beego.AppConfig.String("open-alydh")
	if open != "1" {
		logs.Info(logsign, "[alyphonecall]", "阿里云电话接口未配置未开启状态,请先配置open-alydh为1")
		return "阿里云电话接口未配置未开启状态,请先配置open-alydh为1"
	}
	AccessKeyId := beego.AppConfig.String("ALY_DH_AccessKeyId")
	AccessSecret := beego.AppConfig.String("ALY_DH_AccessSecret")
	CalledShowNumber := beego.AppConfig.String("ALY_DX_CalledShowNumber")
	TtsCode := beego.AppConfig.String("ALY_DH_TtsCode")

	mobiles := strings.Split(PhoneNumbers, ",")
	for _, m := range mobiles {
		client, err := dyvmsapi.NewClientWithAccessKey("cn-hangzhou", AccessKeyId, AccessSecret)
		request := dyvmsapi.CreateSingleCallByTtsRequest()
		request.Scheme = "https"
		request.CalledShowNumber = CalledShowNumber
		request.CalledNumber = m
		request.TtsCode = TtsCode
		request.TtsParam = `{"code":"` + Messages + `"}`
		request.PlayTimes = requests.NewInteger(2)

		response, err := client.SingleCallByTts(request)
		if err != nil {
			logs.Error(logsign, "[alyphonecall]", err.Error())
		}
		logs.Info(logsign, "[alyphonecall]", response)
	}
	model.AlertToCounter.WithLabelValues("alydh", Messages, PhoneNumbers).Add(1)
	return PhoneNumbers + "Called Over."
}
