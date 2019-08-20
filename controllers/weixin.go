package controllers

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Url string `json:"url"`
	Picurl string `json:"picurl"`
}
type New struct {
	Articles []Article `json:"articles"`
}
type WXMessage struct {
	Msgtype string `json:"msgtype"`
	News New `json:"news"`
}
/*
{
    "msgtype": "news",
    "news": {
       "articles" : [
           {
               "title" : "For Test Messages,请忽略",
               "description" : "告警级别:严重\n开始时间:2019-08-13 23:52:19\n结束时间:2019-08-14 12:58:09\n故障主机IP:172.16.83.164:8080\n部署 default/demo 当前状态不可用.",
               "url" : "https://prometheus-dev.gn.i-tetris.com/alerts",
               "picurl" : "https://www.megatronix.cn/logo.png?v=1"
           }
        ]
    }
}
*/
func PostToWeiXin(title,text,alerturl,logourl,WXurl string)(string)  {
	u := WXMessage{
		Msgtype:"news",
		News:New{
			Articles: []Article{
				Article{
					Title:title,
					Description:text,
					Url:alerturl,
					Picurl:logourl,
				},
			},
		},
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