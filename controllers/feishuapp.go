package controllers

import (
	"PrometheusAlert/models"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type FSAPPConf struct {
	WideScreenMode bool `json:"wide_screen_mode"`
	EnableForward  bool `json:"enable_forward"`
}

type FSAPPTe struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type FSAPPElement struct {
	Tag           string         `json:"tag"`
	Text          Te             `json:"text"`
	Content       string         `json:"content"`
	FSAPPElements []FSAPPElement `json:"elements"`
}

type FSAPPTitles struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type FSAPPHeaders struct {
	FSAPPTitle FSAPPTitles `json:"title"`
	Template   string      `json:"template"`
}

type FSAPPCards struct {
	FSAPPConfig   FSAPPConf      `json:"config"`
	FSAPPElements []FSAPPElement `json:"elements"`
	FSAPPHeader   FSAPPHeaders   `json:"header"`
}

type FSContentAPP struct {
	MsgType      string `json:"msg_type"`
	ReceiveId    string `json:"receive_id"` //用户传入的ID，可以是 open_id、user_id、union_id、email、chat_id
	FSAPPContent string `json:"content"`
}

func GetAccessToken(logsign string) (string, error) {
	// https://open.feishu.cn/open-apis/message/v4/batch_send/ 批量发送消息  tenant_access_token
	// 先获取 tenant_access_token
	u := TenantAccessMeg{
		AppId:     beego.AppConfig.String("FEISHU_APPID"),
		AppSecret: beego.AppConfig.String("FEISHU_APPSECRET"),
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
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
	//res, err := client.Post("https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal", "application/json; charset=utf-8", b)
	res, err := http.NewRequest("POST", "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal", b)
	if err != nil {
		logs.Error(logsign, "[feishuapp]", err.Error())
		return "", err
	}
	res.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(res)
	defer res.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(logsign, "[feishuapp]", err.Error())
		return "", err
	}
	resp_json := TenantAccessResp{}
	json.Unmarshal(result, &resp_json)
	if resp_json.Msg != "ok" {
		logs.Error(logsign, "[feishuapp]", resp_json.Msg)
		return "", errors.New(resp_json.Msg)
	}
	logs.Info(logsign, "[feishuapp]", string(result))
	return resp_json.TenantAccessToken, nil
}

func PostToFeiShuApp(title, text, receiveIds, logsign string) string {
	open := beego.AppConfig.String("open-feishuapp")
	if open != "1" {
		logs.Info(logsign, "[feishuapp]", "飞书APP接口未配置未开启状态,请先配置open-feishuapp为1")
		return "飞书APP接口未配置未开启状态,请先配置open-feishuapp为1"
	}
	var color string
	if strings.Count(text, "resolved") > 0 && strings.Count(text, "firing") > 0 {
		color = "orange"
	} else if strings.Count(text, "resolved") > 0 {
		color = "green"
	} else {
		color = "red"
	}
	token, err := GetAccessToken(logsign)
	if err != nil {
		logs.Error(logsign, "[feishuapp]", err.Error())
		return err.Error()
	}
	SendContent := text
	var result []byte
	if receiveIds != "" {
		ReceiveIds := strings.Split(receiveIds, ",")
		fsAppContent :=
			&FSAPPCards{
				FSAPPConfig: FSAPPConf{
					WideScreenMode: true,
					EnableForward:  true,
				},
				FSAPPHeader: FSAPPHeaders{
					FSAPPTitle: FSAPPTitles{
						Content: title,
						Tag:     "plain_text",
					},
					Template: color,
				},
				FSAPPElements: []FSAPPElement{
					FSAPPElement{
						Tag: "div",
						Text: Te{
							Content: SendContent,
							Tag:     "lark_md",
						},
					},
					{
						Tag: "hr",
					},
					{
						Tag: "note",
						FSAPPElements: []FSAPPElement{
							{
								Content: title,
								Tag:     "lark_md",
							},
						},
					},
				},
			}
		contentByte, _ := json.Marshal(fsAppContent)
		fmt.Println("fsAppContent: " + string(contentByte))
		for _, ReceiveId := range ReceiveIds {
			u := FSContentAPP{
				MsgType:      "interactive",
				ReceiveId:    ReceiveId,
				FSAPPContent: string(contentByte),
			}
			var ReceiveType string
			if strings.Contains(ReceiveId, "ou_") {
				ReceiveType = "open_id"
			} else if strings.Contains(ReceiveId, "on_") {
				ReceiveType = "union_id"
			} else if strings.Contains(ReceiveId, "oc_") {
				ReceiveType = "chat_id"
			} else if strings.Contains(ReceiveId, "@") {
				ReceiveType = "email"
			} else {
				ReceiveType = "user_id"
			}
			b := new(bytes.Buffer)
			json.NewEncoder(b).Encode(u)
			logs.Info(logsign, "[feishuapp]", b)
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
			FSUrl := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/messages?receive_id_type=%s", ReceiveType)
			req, err := http.NewRequest("POST", FSUrl, b)
			if err != nil {
				logs.Error(logsign, "[feishuapp]", title+": "+err.Error())
			}
			req.Header.Set("Content-Type", "application/json; charset=utf-8")
			req.Header.Set("Authorization", "Bearer "+token)
			resp, err := client.Do(req)
			if err != nil {
				logs.Error(logsign, "[feishuapp]", err.Error())
			}
			defer resp.Body.Close()
			result, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				logs.Error(logsign, "[feishuapp]", title+": "+err.Error())
			}
			models.AlertToCounter.WithLabelValues("feishuapp").Add(1)
			ChartsJson.Feishu += 1
			logs.Info(logsign, "[feishuapp]", title+": "+string(result))
			//return string(result)
		}
	}
	return string(result)
}
