package controllers

import (
	"PrometheusAlert/model"
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//定义语音请求body结构体
type BodyStr struct {
	To          string `json:"to"`
	MediaName   string `json:"mediaName"`
	MediaTxt    string `json:"mediaTxt"`
	AppId       string `json:"appId"`
	DisplayNum  string `json:"displayNum"`
	PlayTimes   string `json:"playTimes"`
	RespUrl     string `json:"respUrl"`
	UserData    string `json:"userData"`
	MaxCallTime string `json:"maxCallTime"`
	Speed       string `json:"speed"`
	Volume      string `json:"volume"`
	Pitch       string `json:"pitch"`
	Bgsound     string `json:"bgsound"`
}

//生成sig，auth
func GetSigAuth() (string, string) {

	accountSid := beego.AppConfig.String("RLY_ACCOUNT_SID")
	accountToken := beego.AppConfig.String("RLY_ACCOUNT_TOKEN")
	now := time.Now()
	calltime := now.Format("20060102150405")
	signature := strings.Join([]string{accountSid, accountToken, calltime}, "")
	sigdata := []byte(signature)
	sighas := md5.Sum(sigdata)
	sig := strings.ToUpper(fmt.Sprintf("%x", sighas))
	auth := strings.Join([]string{accountSid, calltime}, ":")

	return sig, auth
}

func PostRLYphonecall(CallMessage, PhoneNumber, logsign string) string {
	var (
		body []byte
	)
	open := beego.AppConfig.String("RLY_DH_open-rlydh")
	accountSid := beego.AppConfig.String("RLY_ACCOUNT_SID")

	if open != "1" {
		logs.Info(logsign, "[rlyphonecall]", "容联云语音接口未配置未开启状态,请先配置open-rlydh为1")
		return "容联云语音接口未配置未开启状态,请先配置open-txdh为1"
	}

	appId := beego.AppConfig.String("RLY_APP_ID")
	url := beego.AppConfig.String("RLY_URL")
	//整合body字符
	var Body BodyStr
	Body.To = PhoneNumber
	Body.MediaTxt = CallMessage
	Body.AppId = appId
	Body.PlayTimes = "3" //此处定义默认重复三次
	BodyData, err := json.Marshal(Body)
	if err != nil {
		logs.Error(logsign, "[rlyphonecall]", err)
	}

	//获取sig，auth信息
	sigStr, authStr := GetSigAuth()

	//整合auth信息
	encodeAuth := base64.StdEncoding.EncodeToString([]byte(authStr))

	//整合url请求
	urlPath := strings.Join([]string{url, accountSid, "/Calls/LandingCalls?sig=", sigStr}, "")

	//开始请求
	client := &http.Client{}
	req, err := http.NewRequest("POST", urlPath, bytes.NewBuffer([]byte(string(BodyData))))
	if err != nil {
		// handle error
		logs.Error(logsign, "[rlyphonecall]", err)
	}
	//设置表头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", encodeAuth)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		logs.Error(logsign, "[rlyphonecall]", err)
	}
	model.AlertToCounter.WithLabelValues("rlydx", CallMessage, PhoneNumber).Add(1)
	return string(body)
}
