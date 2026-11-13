package controllers

import (
	"PrometheusAlert/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"net"
)

// SendVoice
func SendVoice(message, logsign string) string {
	open := beego.AppConfig.String("open-voice")
	if open != "1" {
		logs.Warn(logsign, "[voicemessage] [channel-disabled] Custom Voice broadcast is not enabled (open-voice != 1)")
		return "语音播报配置未开启,请先配置open-voice为1"
	}
	v_ip := beego.AppConfig.String("VOICE_IP")
	v_port := beego.AppConfig.String("VOICE_PORT")
	logs.Info(logsign, "[voicemessage] [connect-attempt] Target TCP Addr:", v_ip+":"+v_port)
	//发送tcp语音消息文本
	v_addr, err := net.ResolveTCPAddr("tcp", v_ip+":"+v_port)
	if err != nil {
		logs.Error(logsign, "[voicemessage] [connect-failed] Target TCP Addr:", v_ip+":"+v_port, "| Error:", err.Error())
		return "语音组件连接初始化失败：" + err.Error()
	}
	conn, err := net.DialTCP("tcp", nil, v_addr)
	if err != nil {
		logs.Error(logsign, "[voicemessage] [connect-failed] Target TCP Addr:", v_ip+":"+v_port, "| Error:", err.Error())
		return "语音组件连接失败：" + err.Error()
	}
	logs.Info(logsign, "[voicemessage] [send-attempt] Sending payload over TCP.")
	_, err = conn.Write([]byte(message))
	if err != nil {
		logs.Error(logsign, "[voicemessage] [send-failed] Target TCP Addr:", v_ip+":"+v_port, "| Error:", err.Error())
		conn.Close()
		return "语音组件发送消息失败：" + err.Error()
	}
	logs.Info(logsign, "[voicemessage] [send-success] Message sent successfully. Payload:", message)
	conn.Close()
	models.AlertToCounter.WithLabelValues("voice").Add(1)
	ChartsJson.Voice += 1
	return message + "  语音播报消息发送成功"
}
