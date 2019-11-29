package controllers

import (
	"bytes"
	"crypto/tls"
	"github.com/astaxie/beego"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

type Mark struct {
	Content string `json:"content"`
}
type WXMessage struct {
	Msgtype string `json:"msgtype"`
	Markdown Mark `json:"markdown"`
}
func PostToWeiXin(text,WXurl string)(string)  {
	open:=beego.AppConfig.String("open-weixin")
	if open=="0" {
		return "企业微信接口未配置未开启状态,请先配置open-weixin为1"
	}
	u := WXMessage{
		Msgtype:"markdown",
		Markdown:Mark{Content:text},
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	log.SetPrefix("[DEBUG 2]")
	log.Println(b)
	//url="http://127.0.0.1:8081"
	tr :=&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res,err  := client.Post(WXurl, "application/json", b)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()
	result,err:=ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	return string(result)
}