package controllers

import (
	"PrometheusAlert/models"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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

func PostToDingDing(title, text, Ddurl, AtSomeOne, logsign string) string {
	open := beego.AppConfig.String("open-dingding")
	if open != "1" {
		logs.Info(logsign, "[dingding]", "钉钉接口未配置未开启状态,请先配置open-dingding为1")
		return "钉钉接口未配置未开启状态,请先配置open-dingding为1"
	}
	// dingding sign
	if openSecret := beego.AppConfig.String("open-dingding-secret"); openSecret == "1" {
		Ddurl = dingdingSign(Ddurl)
	}

	Isatall, _ := beego.AppConfig.Int("dd_isatall")
	Atall := true
	if Isatall == 0 {
		Atall = false
	}
	atMobile := []string{"15888888888"}
	SendText := text
	if AtSomeOne != "" {
		atMobile = strings.Split(AtSomeOne, ",")
		AtText := ""
		for _, phoneN := range atMobile {
			AtText += " @" + phoneN
		}
		SendText += AtText
		Atall = false
	}

	u := DDMessage{
		Msgtype: "markdown",
		Markdown: struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		}{Title: title, Text: SendText},
		At: struct {
			AtMobiles []string `json:"atMobiles"`
			IsAtAll   bool     `json:"isAtAll"`
		}{AtMobiles: atMobile, IsAtAll: Atall},
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
	result, err := io.ReadAll(res.Body)
	if err != nil {
		logs.Error(logsign, "[dingding]", err.Error())
	}
	models.AlertToCounter.WithLabelValues("dingding").Add(1)
	ChartsJson.Dingding += 1
	logs.Info(logsign, "[dingding]", string(result))
	return string(result)
}

// dingdingSign adds sign and timestamp parms to dingding webhook url
// docs: https://open.dingtalk.com/document/orgapp/custom-bot-creation-and-installation
func dingdingSign(ddurl string) string {
	timestamp := time.Now()
	timestampMs := timestamp.UnixNano() / int64(time.Millisecond)
	tsMsStr := strconv.FormatInt(timestampMs, 10)
	// parse ddurl parms
	u, err := url.Parse(ddurl)
	if err != nil {
		logs.Info("[dingdingSign]", "配置文件已开启钉钉加签，钉钉机器人地址解析加签参数 secret 失败，将使用不加签的地址！")
		return ddurl
	}
	// get parm secret
	queryParams := u.Query()
	secret := queryParams.Get("secret")
	if len(secret) == 0 {
		logs.Info("[dingdingSign]", "配置文件已开启钉钉加签，钉钉机器人地址解析加签参数 secret 为空，将使用不加签的地址！")
		return ddurl
	}
	// sign string
	signStr := tsMsStr + "\n" + secret
	// HmacSHA256
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(signStr))
	signature := h.Sum(nil)
	// Base64
	sign := base64.StdEncoding.EncodeToString(signature)
	// splice url
	delete(queryParams, "secret")
	queryParams.Add("timestamp", tsMsStr)
	queryParams.Add("sign", sign)
	u.RawQuery = queryParams.Encode()
	signURL := u.String()

	return signURL
}
