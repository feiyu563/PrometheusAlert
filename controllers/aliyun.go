package controllers

import (
	"PrometheusAlert/models"
	"fmt"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dyvmsapi"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func PostALYmessage(Messages, PhoneNumbers, logsign string) string {
	open := beego.AppConfig.String("open-alydx")
	if open != "1" {
		logs.Warn(logsign, "[alymessage] [channel-disabled] Alibaba Cloud SMS is not enabled (open-alydx != 1)")
		return "阿里云短信接口未配置未开启状态,请先配置open-alydx为1"
	}
	AccessKeyId := beego.AppConfig.String("ALY_DX_AccessKeyId")
	AccessSecret := beego.AppConfig.String("ALY_DX_AccessSecret")
	SignName := beego.AppConfig.String("ALY_DX_SignName")
	Template := beego.AppConfig.String("ALY_DX_Template")

	logs.Info(logsign, "[alymessage] [send-attempt] Target PhoneNumbers:", PhoneNumbers, "| SignName:", SignName, "| Template:", Template, "| Param:", Messages)
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", AccessKeyId, AccessSecret)
	if err != nil {
		logs.Error(logsign, "[alymessage] [client-init-failed] Error:", err.Error())
		return err.Error()
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = PhoneNumbers
	request.SignName = SignName
	request.TemplateCode = Template
	request.TemplateParam = `{"code":"` + Messages + `"}`
	response, err := client.SendSms(request)

	if err != nil {
		logs.Error(logsign, "[alymessage] [send-failed] Target PhoneNumbers:", PhoneNumbers, "| Error:", err.Error())
		return err.Error()
	}
	logs.Info(logsign, "[alymessage] [send-success] Target PhoneNumbers:", PhoneNumbers, "| Code:", response.Code, "| Message:", response.Message, "| RequestId:", response.RequestId)
	models.AlertToCounter.WithLabelValues("alydx").Add(1)
	ChartsJson.Alydx += 1
	return response.Message
}
func PostALYphonecall(Messages string, PhoneNumbers, logsign string) string {
	open := beego.AppConfig.String("open-alydh")
	if open != "1" {
		logs.Warn(logsign, "[alyphonecall] [channel-disabled] Alibaba Cloud Voice is not enabled (open-alydh != 1)")
		return "阿里云电话接口未配置未开启状态,请先配置open-alydh为1"
	}
	AccessKeyId := beego.AppConfig.String("ALY_DH_AccessKeyId")
	AccessSecret := beego.AppConfig.String("ALY_DH_AccessSecret")
	CalledShowNumber := beego.AppConfig.String("ALY_DX_CalledShowNumber")
	TtsCode := beego.AppConfig.String("ALY_DH_TtsCode")

	var errorCollector []string
	mobiles := strings.Split(PhoneNumbers, ",")
	for _, m := range mobiles {
		logs.Info(logsign, "[alyphonecall] [send-attempt] Target CalledNumber:", m, "| CalledShowNumber:", CalledShowNumber, "| TtsCode:", TtsCode, "| Param Msg:", Messages)
		client, err := dyvmsapi.NewClientWithAccessKey("cn-hangzhou", AccessKeyId, AccessSecret)
		if err != nil {
			logs.Error(logsign, "[alyphonecall] [client-init-failed] Target CalledNumber:", m, "| Error:", err.Error())
			errorCollector = append(errorCollector, fmt.Sprintf("to %s client-init-failed: %s", m, err.Error()))
			continue
		}
		request := dyvmsapi.CreateSingleCallByTtsRequest()
		request.Scheme = "https"
		request.CalledShowNumber = CalledShowNumber
		request.CalledNumber = m
		request.TtsCode = TtsCode
		request.TtsParam = `{"code":"` + Messages + `"}`
		request.PlayTimes = requests.NewInteger(2)

		response, err := client.SingleCallByTts(request)
		if err != nil {
			logs.Error(logsign, "[alyphonecall] [send-failed] Target CalledNumber:", m, "| Error:", err.Error())
			errorCollector = append(errorCollector, fmt.Sprintf("to %s send-failed: %s", m, err.Error()))
			continue
		}
		if response.Code != "OK" {
			errorCollector = append(errorCollector, fmt.Sprintf("to %s failed: code %s, msg %s", m, response.Code, response.Message))
			logs.Error(logsign, "[alyphonecall] [send-failed] Target CalledNumber:", m, "| Code:", response.Code, "| Message:", response.Message)
		}
		logs.Info(logsign, "[alyphonecall] [send-success] Target CalledNumber:", m, "| Code:", response.Code, "| Message:", response.Message, "| CallId:", response.CallId, "| RequestId:", response.RequestId)
	}
	models.AlertToCounter.WithLabelValues("alydh").Add(1)
	ChartsJson.Alydh += 1
	if len(errorCollector) > 0 {
		return strings.Join(errorCollector, "; ")
	}
	return PhoneNumbers + " all called successfully."
}
