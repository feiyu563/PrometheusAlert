package controllers

import (
	"PrometheusAlert/model"
	"crypto/tls"
	"encoding/base64"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//华为云短信子程序
func PostHWmessage(Messages string, PhoneNumbers, logsign string) string {
	open := beego.AppConfig.String("open-hwdx")
	if open != "1" {
		logs.Info(logsign, "[hwmessage]", "华为云短信接口未配置未开启状态,请先配置open-hwdx为1")
		return "华为云短信接口未配置未开启状态,请先配置open-hwdx为1"
	}
	hwappkey := beego.AppConfig.String("HWY_DX_APP_Key")
	hwappsecret := beego.AppConfig.String("HWY_DX_APP_Secret")
	hwappurl := beego.AppConfig.String("HWY_DX_APP_Url")
	hwtplid := beego.AppConfig.String("HWY_DX_Templateid")
	hwsign := beego.AppConfig.String("HWY_DX_Signature")
	sender := beego.AppConfig.String("HWY_DX_Sender")
	//mobile格式:"15395105573,16619875573"
	//生成header
	now := time.Now().Format("2006-01-02T15:04:05Z")
	nonce := "7226249334"
	digest := getSha256Code(nonce + now + hwappsecret)
	//digestBase64:=base64.URLEncoding.EncodeToString([]byte(digest))
	digestBase64 := base64.StdEncoding.EncodeToString([]byte(digest))
	xheader := `"UsernameToken Username="` + hwappkey + `",PasswordDigest="` + digestBase64 + `",Nonce="` + nonce + `",Created="` + now + `"`
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
	req, _ := http.NewRequest("POST", hwappurl+"/sms/batchSendSms/v1", strings.NewReader(url.Values{"from": {sender}, "to": {PhoneNumbers}, "templateId": {hwtplid}, "templateParas": {"[" + Messages + "]"}, "signature": {hwsign}, "statusCallback": {""}, "extend": {logsign}}.Encode()))
	req.Header.Set("Authorization", `WSSE realm="SDP",profile="UsernameToken",type="Appkey"`)
	req.Header.Set("X-WSSE", xheader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		logs.Error(logsign, "[hwmessage]", err.Error())
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(logsign, "[hwmessage]", err.Error())
	}
	model.AlertToCounter.WithLabelValues("hwdx", Messages, PhoneNumbers).Add(1)
	logs.Info(logsign, "[hwmessage]", string(result))
	return string(result)
}
