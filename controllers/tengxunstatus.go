package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"encoding/json"
)

type TengXunStatusController struct {
	beego.Controller
}

type CallBack struct {
	Voiceprompt_callback Vcallback `json:"voiceprompt_callback"`
}

type Vcallback struct {
	Result string `json:"result"`
	Accept_time string `json:"accept_time"`
	Call_from string `json:"call_from"`
	Callid string `json:"callid"`
	End_calltime string `json:"end_calltime"`
	Fee string `json:"fee"`
	Mobile string `json:"mobile"`
	Nationcode string `json:"nationcode"`
	Start_calltime string `json:"start_calltime"`
}

type Re struct {
	Result int `json:"result"`
	Errmsg string `json:"errmsg"`
} 

func (c *TengXunStatusController) TengXunStatus() {
	TengXunReturn:=CallBack{}
	log.SetPrefix("[DEBUG tengxun]")
	log.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &TengXunReturn)
	if TengXunReturn.Voiceprompt_callback.Result!="0" && TengXunReturn.Voiceprompt_callback.Mobile==GetUserPhone(1) {
		CallOthers(3)
	}
	if TengXunReturn.Voiceprompt_callback.Result!="0" && TengXunReturn.Voiceprompt_callback.Mobile==GetUserPhone(3) {
		CallOthers(5)
	}
	result:=Re{}
	result.Result=0
	result.Errmsg="ok"
	c.Data["json"]=result
	c.ServeJSON()
}

func CallOthers(Num int)  {
	phone:=GetUserPhone(Num)
	LogMessage:="PostTXphonecall:"+PostTXphonecall(PhoneCallMessage,phone)+"\n"
	log.Println(LogMessage)
}