package controllers

import (
	"PrometheusAlert/models"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
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

type FSAPP struct {
	MsgType       string     `json:"msg_type"`
	UnionIds      []string   `json:"union_ids"`      //@所使用字段 支持自定义部门ID，和open_department_id，列表长度小于等于 200 注：部门下的所有子部门包含的成员也会收到消息 示例值：["3dceba33a33226","d502aaa9514059", "od-5b91c9affb665451a16b90b4be367efa"]
	UserIds       []string   `json:"user_ids"`       //@所使用字段 用户 user_id 列表，长度小于等于 200 （对应 V3 接口的 employee_ids ） 示例值：["7cdcc7c2","ca51d83b"]
	OpenIds       []string   `json:"open_ids"`       //@所使用字段 用户 open_id 列表，长度小于等于 200 示例值：["ou_18eac85d35a26f989317ad4f02e8bbbb","ou_461cf042d9eedaa60d445f26dc747d5e"]
	DepartmentIds []string   `json:"department_ids"` //@所使用字段 用户 union_ids 列表，长度小于等于 200 示例值：["on_cad4860e7af114fb4ff6c5d496d1dd76","on_gdcq860e7af114fb4ff6c5d496dabcet"]
	FSAPPCard     FSAPPCards `json:"card"`
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

func PostToFeiShuApp(title, text, userIds, logsign string) string {
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
	SendContent := text
	SendContentJson := []string{}
	if userIds != "" {
		UserIds := strings.Split(userIds, ",")
		UserIdtext := ""
		for _, UserId := range UserIds {
			UserIdtext += "<at user_id=" + UserId + "></at>"
			SendContentJson = append(SendContentJson, UserId)
		}

		SendContent += UserIdtext
	}

	u := FSAPP{
		MsgType:       "interactive",
		UnionIds:      SendContentJson,
		UserIds:       SendContentJson,
		OpenIds:       SendContentJson,
		DepartmentIds: SendContentJson,
		FSAPPCard: FSAPPCards{
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
		},
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
	req, err := http.NewRequest("POST", "https://open.feishu.cn/open-apis/message/v4/batch_send/", b)
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
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(logsign, "[feishuapp]", title+": "+err.Error())
	}
	models.AlertToCounter.WithLabelValues("feishuapp").Add(1)
	ChartsJson.Feishu += 1
	logs.Info(logsign, "[feishuapp]", title+": "+string(result))
	return string(result)
}
