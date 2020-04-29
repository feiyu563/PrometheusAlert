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

type FSMessage struct {
	Title string `json:"title"`
	Text string `json:"text"`
}

func PostToFeiShu(title,text,Fsurl,logsign string)(string)  {
	open:=beego.AppConfig.String("open-feishu")
	if open=="0" {
		logs.Info(logsign,"[dingding]","飞书接口未配置未开启状态,请先配置open-feishu为1")
		return "飞书接口未配置未开启状态,请先配置open-feishu为1"
	}

	u := FSMessage{Title: title, Text: text}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	logs.Info(logsign,"[feishu]",b)
	var tr *http.Transport
	if proxyUrl := beego.AppConfig.String("proxy");proxyUrl != ""{
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		}
		tr = &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			Proxy: proxy,
		}
	}else{
		tr = &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		}
	}
	client := &http.Client{Transport: tr}
	res,err  := client.Post(Fsurl, "application/json", b)
	if err != nil {
		logs.Error(logsign,"[feishu]",err.Error())
	}
	defer res.Body.Close()
	result,err:=ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error(logsign,"[feishu]",err.Error())
	}
	model.AlertToCounter.WithLabelValues("feishu",text,"").Add(1)
	logs.Info(logsign,"[feishu]",string(result))
	return string(result)
}