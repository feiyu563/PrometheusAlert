package controllers

import (
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"encoding/json"
)

//腾讯短信接口sha256编码
func getSha256Code(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

//腾讯短信接口消息格式
type Mobiles struct{
	Mobile string `json:"mobile"`
	Nationcode string `json:"nationcode"`
}

type TXmessage struct {
	Ext string `json:"ext"`
	Extend string `json:"extend"`
	Params []string `json:"params"`
	Sig string `json:"sig"`
	Sign string `json:"sign"`
	Tel []Mobiles `json:"tel"`
	Time int `json:"time"`
	Tpl_id int `json:"tpl_id"`
}

//腾讯短信子程序
func PostTXmessage(text string,mobile string)(string)  {
	strAppKey:=beego.AppConfig.String("appkey")
	tpl_id,_:=beego.AppConfig.Int("tpl_id")
	sdkappid:=beego.AppConfig.String("sdkappid")
	//腾讯短信接口算法部分
	//mobile格式:"15395105573,16619875573"
	TXmobile:=Mobiles{}
	TXmobiles:=[]Mobiles{}
	mobiles:=strings.Split(mobile,",")
	for _,m:=range mobiles {
		TXmobile.Mobile=m
		TXmobile.Nationcode="86"
		TXmobiles=append(TXmobiles,TXmobile )
	}
	strRand := "7226249334"
	strTime := strconv.FormatInt(time.Now().Unix(),10)
	intTime,_:=strconv.Atoi(strTime)
	sig := getSha256Code("appkey="+strAppKey+"&random="+strRand+"&time="+strTime+"&mobile="+mobile)
	TXurl:="https://yun.tim.qq.com/v5/tlssmssvr/sendmultisms2?sdkappid="+sdkappid+"&random="+strRand
	u := TXmessage{
		Ext:"",
		Extend:"",
		Params:[]string{text},
		Sig:sig,
		Tel:TXmobiles,
		Time:intTime,
		Tpl_id:tpl_id,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	log.SetPrefix("[DEBUG 2]")
	log.Println(b)
	tr :=&http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	//res,err := http.Post(Ddurl, "application/json", b)
	//resp, err := http.PostForm(url,url.Values{"key": {"Value"}, "id": {"123"}})
	client := &http.Client{Transport: tr}
	res,err  := client.Post(TXurl, "application/json", b)
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

//腾讯语音提醒接口
type TXphonecall struct {
	Ext string `json:"ext"`
	Tpl_id int `json:"tpl_id"`
	Params []string `json:"params"`
	Playtimes int `json:"playtimes"`
	Sig string `json:"sig"`
	Tel Mobiles `json:"tel"`
	Time int `json:"time"`
}

//腾讯语音子程序
func PostTXphonecall(text string,mobile string)(string)  {
	strAppKey:=beego.AppConfig.String("phonecallappkey")
	sdkappid:=beego.AppConfig.String("phonecallsdkappid")
	tpl_id,_:=beego.AppConfig.Int("phonecalltpl_id")
	//腾讯短信接口算法部分
	TXmobile:=Mobiles{}
	mobiles:=strings.Split(mobile,",")
	for _,m:=range mobiles {
		TXmobile.Mobile=m
		TXmobile.Nationcode="86"
		strRand := "7226249334"
		strTime := strconv.FormatInt(time.Now().Unix(),10)
		intTime,_:=strconv.Atoi(strTime)
		sig := getSha256Code("appkey="+strAppKey+"&random="+strRand+"&time="+strTime+"&mobile="+m)
		TXurl:="https://cloud.tim.qq.com/v5/tlsvoicesvr/sendtvoice?sdkappid="+sdkappid+"&random="+strRand
		u := TXphonecall{
			Ext:"",
			Tpl_id:tpl_id,
			Params:[]string{text},
			Playtimes:2,
			Sig:sig,
			Tel:TXmobile,
			Time:intTime,
		}
		log.SetPrefix("[DEBUG 2]")
		log.Println(PhoneCallPost(TXurl,u))
	}
	return "ok"
}

func PhoneCallPost(url string,u TXphonecall)(string) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	log.SetPrefix("[DEBUG 2]")
	log.Println(b)
	tr :=&http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	//res,err := http.Post(Ddurl, "application/json", b)
	//resp, err := http.PostForm(url,url.Values{"key": {"Value"}, "id": {"123"}})
	client := &http.Client{Transport: tr}
	res,err  := client.Post(url, "application/json", b)
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