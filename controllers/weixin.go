package controllers

import (
	"PrometheusAlert/model"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Mark struct {
	Content string `json:"content"`
}
type WXMessage struct {
	Msgtype  string `json:"msgtype"`
	Markdown Mark   `json:"markdown"`
}

func PostToWeiXin(text, WXurl, logsign string) string {
	open := beego.AppConfig.String("open-weixin")
	if open != "1" {
		logs.Info(logsign, "[weixin]", "企业微信接口未配置未开启状态,请先配置open-weixin为1")
		return "企业微信接口未配置未开启状态,请先配置open-weixin为1"
	}
	u := WXMessage{
		Msgtype:  "markdown",
		Markdown: Mark{Content: text},
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	logs.Info(logsign, "[weixin]", b)
	//url="http://127.0.0.1:8081"
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
	res, err := client.Post(WXurl, "application/json", b)
	if err != nil {
		logs.Error(logsign, "[weixin]", err.Error())
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error(logsign, "[weixin]", err.Error())
	}
	model.AlertToCounter.WithLabelValues("weixin", text, "").Add(1)
	logs.Info(logsign, "[weixin]", string(result))
	return string(result)
}
