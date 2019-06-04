package controllers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type DDMessage struct {
	Msgtype string `json:"msgtype"`
	Markdown struct{
		Title string `json:"title"`
		Text string `json:"text"`
	} `json:"markdown"`
	At struct{
		AtMobiles []string `json:"atMobiles"`
		IsAtAll bool `json:"isAtAll"`
	} `json:"at"`
}

func PostToDingDing(title,text,Ddurl string)(string)  {
	u := DDMessage{
		Msgtype:"markdown",
		Markdown: struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		}{Title: title, Text: text},
		At: struct {
			AtMobiles []string `json:"atMobiles"`
			IsAtAll   bool `json:"isAtAll"`
		}{AtMobiles:[]string{"15395105573"} , IsAtAll:true },
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
	res,err  := client.Post(Ddurl, "application/json", b)
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