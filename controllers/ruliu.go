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
	"strings"
)

type RuLiuMessage struct {
	Message struct{
		Header struct{
			Toid []int `json:"toid"`
		} `json:"header"`
		Body []struct{
			Type string `json:"type"`
			Content string `json:"content"`
		} `json:"body"`
	} `json:"message"`
}

func PostToRuLiu(Gids,text, RLurl, logsign string) string {
	open := beego.AppConfig.String("open-ruliu")
	if open != "1" {
		logs.Info(logsign, "[ruliu]", "钉钉接口未配置未开启状态,请先配置open-ruliu为1")
		return "如流接口未配置未开启状态,请先配置open-ruliu为1"
	}
	GroupIds:=[]int{}
	sGid := strings.Split(Gids, ",")
	for Gid:=range sGid{
		GroupIds=append(GroupIds, Gid)
	}
	u := RuLiuMessage{
		Message: struct {
			Header struct{ Toid []int `json:"toid"` }
			Body []struct {
				Type    string `json:"type"`
				Content string `json:"content"`
			}
		}{Header: struct {
			Toid []int `json:"toid"`
		}{GroupIds}, Body: []struct {
			Type    string `json:"type"`
			Content string `json:"content"`
		}{{"MD",text}}},
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	logs.Info(logsign, "[ruliu]", b)
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
	res, err := client.Post(RLurl, "application/json", b)
	if err != nil {
		logs.Error(logsign, "[ruliu]", err.Error())
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error(logsign, "[ruliu]", err.Error())
	}
	model.AlertToCounter.WithLabelValues("ruliu", text, "").Add(1)
	logs.Info(logsign, "[ruliu]", string(result))
	return string(result)
}
