package controllers

import (
	"bytes"
	"crypto/tls"
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
/*
{
    "msgtype": "markdown",
    "markdown": {
        "content": "[Prometheus故障告警信息](https://prometheus-dev.gn.i-tetris.com/alerts)\n>**For Test Messages,请忽略**\n>`告警级别：``严重`\n`开始时间:``2019-08-13 23:52:19`\n`结束时间:``2019-08-14 12:58:09`\n`故障主机IP:``172.16.83.164:8080`\n**部署 default/demo 当前状态不可用.**"
    }
}
*/
func PostToWeiXin(text,WXurl string)(string)  {
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
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	//res,err := http.Post(Ddurl, "application/json", b)
	//resp, err := http.PostForm(url,url.Values{"key": {"Value"}, "id": {"123"}})
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