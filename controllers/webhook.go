package controllers

import (
	"PrometheusAlert/model"
	"bytes"
	"crypto/tls"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"net/url"
)

func PostToWebhook(text, WebhookUrl, logsign string) string {
	logs.Info(logsign, "[Webhook]", text)
	JsonMsg := bytes.NewReader([]byte(text))
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
	res, err := client.Post(WebhookUrl, "application/json", JsonMsg)
	if err != nil {
		logs.Error(logsign, "[Webhook]", err.Error())
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error(logsign, "[Webhook]", err.Error())
	}
	defer res.Body.Close()
	model.AlertToCounter.WithLabelValues("Webhook", text, "").Add(1)
	logs.Info(logsign, "[Webhook]", string(result))
	return string(result)
}
