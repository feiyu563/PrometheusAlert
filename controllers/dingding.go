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

type DDMessage struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

func PostToDingDing(title, text, Ddurl, logsign string) string {
	open := beego.AppConfig.String("open-dingding")
	if open == "0" {
		logs.Info(logsign, "[dingding]", "钉钉接口未配置未开启状态,请先配置open-dingding为1")
		return "钉钉接口未配置未开启状态,请先配置open-dingding为1"
	}
	Isatall, _ := beego.AppConfig.Int("dd_isatall")
	Atall := true
	if Isatall == 0 {
		Atall = false
	}
	u := DDMessage{
		Msgtype: "markdown",
		Markdown: struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		}{Title: title, Text: text},
		At: struct {
			AtMobiles []string `json:"atMobiles"`
			IsAtAll   bool     `json:"isAtAll"`
		}{AtMobiles: []string{"15395105573"}, IsAtAll: Atall},
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	logs.Info(logsign, "[dingding]", b)
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
	res, err := client.Post(Ddurl, "application/json", b)
	if err != nil {
		logs.Error(logsign, "[dingding]", err.Error())
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error(logsign, "[dingding]", err.Error())
	}
	model.AlertToCounter.WithLabelValues("dingding", text, "").Add(1)
	logs.Info(logsign, "[dingding]", string(result))
	return string(result)
}
