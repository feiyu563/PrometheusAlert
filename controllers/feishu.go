package controllers

import (
	"PrometheusAlert/model"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type FSMessage struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func PostToFeiShu(title, text, Fsurl, logsign string) string {
	open := beego.AppConfig.String("open-feishu")
	if open == "0" {
		logs.Info(logsign, "[feishu]", "飞书接口未配置未开启状态,请先配置open-feishu为1")
		return "飞书接口未配置未开启状态,请先配置open-feishu为1"
	}

	u := FSMessage{Title: title, Text: text}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	logs.Info(logsign, "[feishu]", b)
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
	res, err := client.Post(Fsurl, "application/json", b)
	if err != nil {
		logs.Error(logsign, "[feishu]", err.Error())
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error(logsign, "[feishu]", err.Error())
	}
	model.AlertToCounter.WithLabelValues("feishu", text, "").Add(1)
	logs.Info(logsign, "[feishu]", string(result))
	return string(result)
}

type Conf struct {
	WideScreenMode bool `json:"wide_screen_mode"`
	EnableForward  bool `json:"enable_forward"`
}

type Te struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type Header struct {
	Title    Te     `json:"title"`
	Template string `json:"template"`
}

type Element struct {
	Tag      string    `json:"tag"`
	Elements []Element `json:"elements"`
	Content  string    `json:"content"`
	Text     Te        `json:"text"`
}

type Cards struct {
	Config   Conf      `json:"config"`
	Header   Header    `json:"header"`
	Elements []Element `json:"elements"`
}

type FSMessagev2 struct {
	MsgType string `json:"msg_type"`
	Email   string `json:"email"`
	Card    Cards  `json:"card"`
}

func PostToFeiShuv2(title, status, text, Fsurl, logsign string) string {
	open := beego.AppConfig.String("open-feishuv2")
	if open == "0" {
		logs.Info(logsign, "[feishuv2]", "飞书v2接口未配置未开启状态,请先配置open-feishuv2为1")
		return "飞书v2接口未配置未开启状态,请先配置open-feishuv2为1"
	}

	loc, _ := time.LoadLocation("Asia/Chongqing")
	tStr := fmt.Sprintln(time.Now().In(loc).Format("2006-01-02 15:04:05"))
	var color string

	if status == "resolved" {
		color = "green"
	} else {
		color = "red"
	}

	u := FSMessagev2{
		MsgType: "interactive",
		Email:   "244217140@qq.com",
		Card: Cards{
			Config: Conf{
				WideScreenMode: true,
				EnableForward:  true,
			},
			Header: Header{
				Title: Te{
					Content: title,
					Tag:     "plain_text",
				},
				Template: color,
			},
			Elements: []Element{
				Element{
					Tag: "div",
					Text: Te{
						Content: text,
						Tag:     "lark_md",
					},
				},
				{
					Tag: "hr",
				},
				{
					Tag: "note",
					Elements: []Element{
						{
							Content: "PrometheusAlert    ",
							Tag:     "lark_md",
						},
						{
							Content: tStr,
							Tag:     "lark_md",
						},
					},
				},
			},
		},
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	logs.Info(logsign, "[feishuv2]", b)
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
	res, err := client.Post(Fsurl, "application/json", b)
	if err != nil {
		logs.Error(logsign, "[feishuv2]", err.Error())
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error(logsign, "[feishuv2]", err.Error())
	}
	model.AlertToCounter.WithLabelValues("feishuv2", text, "").Add(1)
	logs.Info(logsign, "[feishuv2]", string(result))
	return string(result)
}
