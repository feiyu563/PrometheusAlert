package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type TengXunStatusController struct {
	beego.Controller
}

type CallBack struct {
	Voiceprompt_callback Vcallback `json:"voiceprompt_callback"`
}

type Vcallback struct {
	Result         string `json:"result"`
	Accept_time    string `json:"accept_time"`
	Call_from      string `json:"call_from"`
	Callid         string `json:"callid"`
	End_calltime   string `json:"end_calltime"`
	Fee            string `json:"fee"`
	Mobile         string `json:"mobile"`
	Nationcode     string `json:"nationcode"`
	Start_calltime string `json:"start_calltime"`
}

type Re struct {
	Result int    `json:"result"`
	Errmsg string `json:"errmsg"`
}

func (c *TengXunStatusController) TengXunStatus() {
	TengXunReturn := CallBack{}
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &TengXunReturn)
	DefaultPhone := beego.AppConfig.String("defaultphone")
	//判断失败号码是否是defaultphone中配置的号码,如果不是则继续,是则跳过,不执行重试逻辑
	if TengXunReturn.Voiceprompt_callback.Result != "0" && TengXunReturn.Voiceprompt_callback.Mobile != DefaultPhone {
		//如果失败号码不是user.csv或者defaultphone的号码,那么就直接拨打返回的第一个号码
		if TengXunReturn.Voiceprompt_callback.Result != "0" && TengXunReturn.Voiceprompt_callback.Mobile != GetUserPhone(1) {
			CallOthers(1, logsign)
			//如果失败号码是user.csv的第一个号码,那么就直接拨打user.csv的第2个号码
		} else if TengXunReturn.Voiceprompt_callback.Result != "0" && TengXunReturn.Voiceprompt_callback.Mobile == GetUserPhone(1) {
			CallOthers(3, logsign)
			//如果失败号码是user.csv的第2个号码,那么就直接拨打user.csv的第3个号码
		} else if TengXunReturn.Voiceprompt_callback.Result != "0" && TengXunReturn.Voiceprompt_callback.Mobile == GetUserPhone(3) {
			CallOthers(5, logsign)
		}
	}
	result := Re{}
	result.Result = 0
	result.Errmsg = "ok"
	c.Data["json"] = result
	c.ServeJSON()
}

func CallOthers(Num int, logsign string) {
	phone := GetUserPhone(Num)
	PostTXphonecall(PhoneCallMessage, phone, logsign)
	logs.Info(logsign, "失败重试号码: "+phone)
}
