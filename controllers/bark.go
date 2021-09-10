package controllers

import (
	"PrometheusAlert/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// SendBark 发送消息至iPhone
func SendBark(msg, logsign string) string {
	open := beego.AppConfig.String("open-bark")
	if open != "1" {
		logs.Info(logsign, "[bark]", "bark未配置未开启状态,请先配置open-bark为1")
		return "bark未配置未开启状态,请先配置open-bark为1"
	}
	senduser := beego.AppConfig.String("BARK_KEYS")
	sendusers := strings.Split(senduser, "-")
	for _, u := range sendusers {
		// 处理发送消息
		urlprefix := generateGetUrlPrefix(msg, u)
		barkcopy := beego.AppConfig.String("BARK_COPY")
		if barkcopy == "1" {
			urlprefix += fmt.Sprintf("?copy=%s", msg)
			urlprefix += "&automaticallyCopy=1"
		}
		barkarchive := beego.AppConfig.String("BARK_ARCHIVE")
		if barkarchive == "1" {
			urlprefix += "&isArchive=1"
		}
		urlprefix += fmt.Sprintf("&group=%s", beego.AppConfig.String("BARK_GROUP"))
		get, err := sendBark(urlprefix)
		if err != nil {
			logs.Error(logsign, "[bark]", fmt.Errorf("send to %s, err: %v", u, err))
		}
		if get.Code != 200 {
			logs.Error(logsign, "[bark]", fmt.Errorf("send to %s, get code: %d", u, get.Code))
		}
	}
	model.AlertToCounter.WithLabelValues("bark", "", "").Add(1)
	logs.Info(logsign, "[bark]", "bark send ok.")
	return "bark send ok"
}

type responseMessage struct {
	Code    int64  `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func sendBark(url string) (responseMessage, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return responseMessage{}, err
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return responseMessage{}, err
		}
	}

	message := responseMessage{}
	err = json.Unmarshal(result.Bytes(), &message)
	if err != nil {
		return responseMessage{}, err
	}

	return message, nil
}

func generateGetUrlPrefix(msg, userkey string) string {
	barkserver := beego.AppConfig.String("BARK_URL")
	barktitle := beego.AppConfig.String("BARK_TITLE")
	if len(barktitle) == 0 {
		barktitle = "Bark推送测试"
	}
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("%s/%s/%s/%s", barkserver, userkey, barktitle, msg))
	return buffer.String()
}
