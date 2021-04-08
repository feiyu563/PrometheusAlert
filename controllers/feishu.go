package controllers

import (
	"PrometheusAlert/model"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type FSMessage struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func PostToFS(title, text, Fsurl, logsign string) string {
	open := beego.AppConfig.String("open-feishu")
	if open != "1" {
		logs.Info(logsign, "[feishu]", "飞书接口未配置未开启状态,请先配置open-feishu为1")
		return "飞书接口未配置未开启状态,请先配置open-feishu为1"
	}
	RTstring := ""
	if strings.Contains(Fsurl, "/v2/") {
		RTstring = PostToFeiShuv2(title, text, Fsurl, logsign)
	} else {
		RTstring = PostToFeiShu(title, text, Fsurl, logsign)
	}
	return RTstring
}

func PostToFeiShu(title, text, Fsurl, logsign string) string {
	u := FSMessage{Title: title+"告警消息", Text: text}

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

type Element struct {
	Tag      string    `json:"tag"`
	Text     Te        `json:"text"`
	Content  string    `json:"content"`
	Elements []Element `json:"elements"`
}

type Titles struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type Headers struct {
	Title    Titles `json:"title"`
	Template string `json:"template"`
}

type Cards struct {
	Config   Conf      `json:"config"`
	Elements []Element `json:"elements"`
	Header   Headers   `json:"header"`
}

type FSMessagev2 struct {
	MsgType string `json:"msg_type"`
	Email   string `json:"email"`
	Card    Cards  `json:"card"`
}

func PostToFeiShuv2(title, text, Fsurl, logsign string) string {
	var color string
	if strings.Count(text, "resolved") > 0 && strings.Count(text, "firing") > 0 {
		color = "orange"
		title = title+"告警消息"
	} else if strings.Count(text, "resolved") > 0 {
		color = "green"
		title = title+"告警恢复"
	} else {
		color = "red"
		title = title+"告警消息"
	}
	u := FSMessagev2{
		MsgType: "interactive",
		Email:   "244217140@qq.com",
		Card: Cards{
			Config: Conf{
				WideScreenMode: true,
				EnableForward:  true,
			},
			Header: Headers{
				Title: Titles{
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
		logs.Error(logsign, "[feishuv2]", title+": "+err.Error())
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error(logsign, "[feishuv2]", title+": "+err.Error())
	}
	model.AlertToCounter.WithLabelValues("feishuv2", text, "").Add(1)
	logs.Info(logsign, "[feishuv2]", title+": "+string(result))
	return string(result)
}
