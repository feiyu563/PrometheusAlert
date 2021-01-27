// Package controllers is a beego controller, it implements send sms message, or make a phonecall by 7moor.
package controllers

import (
	"PrometheusAlert/model"
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// messageBody is 7moordx message
type messageBody struct {
	Num         string `json:"num"`
	TemplateNum string `json:"templateNum"`
	Var1        string `json:"var1"`
}

// phonecallBody is 7moordh webcall
type phonecallBody struct {
	Action    string `json:"Action"`
	ServiceNo string `json:"ServiceNo"`
	Exten     string `json:"Exten"`
	Variable  string `json:"Variable"`
}

// Get7MoorSigAuth generates 7moor Authorization and Sig
func Get7MoorSigAuth() (string, string) {
	accountID := beego.AppConfig.String("7MOOR_ACCOUNT_ID")
	accountAPISecret := beego.AppConfig.String("7MOOR_ACCOUNT_APISECRET")
	now := time.Now()
	// 7moor time format
	calltime := now.Format("20060102150405")
	// generate sig
	signature := strings.Join([]string{accountID, accountAPISecret, calltime}, "")
	sigdata := []byte(signature)
	sighas := md5.Sum(sigdata)
	sigMd5 := strings.ToUpper(fmt.Sprintf("%x", sighas))
	// generate authorization
	authorization := strings.Join([]string{accountID, calltime}, ":")
	authBase64 := base64.StdEncoding.EncodeToString([]byte(authorization))

	return sigMd5, authBase64

}

// Post7MOORmessage sends sms message by 7moor
func Post7MOORmessage(Messages string, PhoneNumbers, logsign string) string {
	open := beego.AppConfig.String("open-7moordx")
	if open != "1" {
		logs.Info(logsign, "[7moordx]", "七陌短信接口未配置为开启状态，请先配置open-7moordx为1")
		return "陌短信接口未配置为开启状态，请先配置open-7moordx为1"
	}
	accountID := beego.AppConfig.String("7MOOR_ACCOUNT_ID")
	templateNum := beego.AppConfig.String("7MOOR_DX_TEMPLATENUM")
	sigMd5, authBase64 := Get7MoorSigAuth()
	// generate url
	urlPath := "http://apis.7moor.com/v20160818/sms/sendInterfaceTemplateSms/" + accountID + "?sig=" + sigMd5
	// construct body
	u := messageBody{
		Num:         PhoneNumbers,
		TemplateNum: templateNum,
		Var1:        Messages,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	logs.Info(logsign, "[7moordx]", b)

	var tr *http.Transport
	if proxyUrl := beego.AppConfig.String("proxy"); proxyUrl != "" {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		}
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           proxy,
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	client := &http.Client{Transport: tr}
	// need set headers
	req, err := http.NewRequest("POST", urlPath, b)
	if err != nil {
		logs.Error(logsign, "[7moordx]", err.Error())
	}
	// set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authBase64)
	resp, err := client.Do(req)
	if err != nil {
		logs.Error(logsign, "[7moordx]", err.Error())
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(logsign, "[7moordx]", err.Error())
	}
	model.AlertToCounter.WithLabelValues("7moordx", Messages, PhoneNumbers).Add(1)
	logs.Info(logsign, "[7moordx]", string(result))

	return string(result)
}

// Post7MOORphonecall sends mutiple phonenumbers
func Post7MOORphonecall(Messages string, PhoneNumbers, logsign string) string {
	mobiles := strings.Split(PhoneNumbers, ",")
	for _, mobile := range mobiles {
		go webcallPost(Messages, mobile, logsign)
	}
	model.AlertToCounter.WithLabelValues("7moordh", Messages, PhoneNumbers).Add(1)

	return PhoneNumbers + " called over."
}

// webcallPost makes phonecall by 7moor webcall
func webcallPost(Messages string, PhoneNumber, logsign string) string {
	open := beego.AppConfig.String("open-7moordh")
	if open != "1" {
		logs.Info(logsign, "[7moorphonecall]", "七陌语音通知接口未配置为开启状态，请先配置open-7moordh为1")
		return "七陌语音通知接口未配置为开启状态，请先配置open-7moordh为1"
	}
	accountID := beego.AppConfig.String("7MOOR_ACCOUNT_ID")
	serviceNo := beego.AppConfig.String("7MOOR_WEBCALL_SERVICENO")
	voiceVar := beego.AppConfig.String("7MOOR_WEBCALL_VOICE_VAR")
	sigMd5, authBase64 := Get7MoorSigAuth()
	// generate url
	urlPath := "https://apis.7moor.com/v20160818/webCall/webCall/" + accountID + "?sig=" + sigMd5
	// construct body
	u := phonecallBody{
		Action:    "Webcall",
		ServiceNo: serviceNo,
		Exten:     PhoneNumber,
		Variable:  voiceVar + ":" + Messages,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	logs.Info(logsign, "[7moordh]", b)

	var tr *http.Transport
	if proxyUrl := beego.AppConfig.String("proxy"); proxyUrl != "" {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		}
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           proxy,
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", urlPath, b)
	if err != nil {
		logs.Error(logsign, "[7moordh]", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authBase64)
	resp, err := client.Do(req)
	if err != nil {
		logs.Error(logsign, "[7moordh]", err.Error())
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(logsign, "[7moordh]", err.Error())
	}
	logs.Info(logsign, "[7moordh]", string(result))

	return string(result)
}
