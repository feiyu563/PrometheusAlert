package controllers

import (
	"PrometheusAlert/model"
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//腾讯短信接口sha256编码
func getSha256Code(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

//腾讯短信接口消息格式
type Mobiles struct {
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
}

type TXmessage struct {
	Ext    string    `json:"ext"`
	Extend string    `json:"extend"`
	Params []string  `json:"params"`
	Sig    string    `json:"sig"`
	Sign   string    `json:"sign"`
	Tel    []Mobiles `json:"tel"`
	Time   int       `json:"time"`
	Tpl_id int       `json:"tpl_id"`
}

//腾讯短信子程序
func PostTXmessage(Messages string, PhoneNumbers, logsign string) string {
	open := beego.AppConfig.String("open-txdx")
	if open != "1" {
		logs.Info(logsign, "[txmessage]", "腾讯短信接口未配置未开启状态,请先配置open-txdx为1")
		return "腾讯短信接口未配置未开启状态,请先配置open-txdx为1"
	}
	strAppKey := beego.AppConfig.String("TXY_DX_appkey")
	tpl_id, _ := beego.AppConfig.Int("TXY_DX_tpl_id")
	sdkappid := beego.AppConfig.String("TXY_DX_sdkappid")
	sign := beego.AppConfig.String("TXY_DX_sign")
	//腾讯短信接口算法部分
	//mobile格式:"15395105573,16619875573"
	TXmobile := Mobiles{}
	TXmobiles := []Mobiles{}
	mobiles := strings.Split(PhoneNumbers, ",")
	for _, m := range mobiles {
		TXmobile.Mobile = m
		TXmobile.Nationcode = "86"
		TXmobiles = append(TXmobiles, TXmobile)
	}
	strRand := "7226249334"
	strTime := strconv.FormatInt(time.Now().Unix(), 10)
	intTime, _ := strconv.Atoi(strTime)
	sig := getSha256Code("appkey=" + strAppKey + "&random=" + strRand + "&time=" + strTime + "&mobile=" + PhoneNumbers)
	TXurl := "https://yun.tim.qq.com/v5/tlssmssvr/sendmultisms2?sdkappid=" + sdkappid + "&random=" + strRand
	u := TXmessage{
		Ext:    logsign,
		Extend: "",
		Params: []string{Messages},
		Sig:    sig,
		Sign:   sign,
		Tel:    TXmobiles,
		Time:   intTime,
		Tpl_id: tpl_id,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	logs.Info(logsign, "[txmessage]", b)

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

	//res,err := http.Post(Ddurl, "application/json", b)
	//resp, err := http.PostForm(url,url.Values{"key": {"Value"}, "id": {"123"}})
	client := &http.Client{Transport: tr}
	res, err := client.Post(TXurl, "application/json", b)

	if err != nil {
		logs.Error(logsign, "[txmessage]", err.Error())
	}

	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error(logsign, "[txmessage]", err.Error())
	}

	model.AlertToCounter.WithLabelValues("txdx", Messages, PhoneNumbers).Add(1)
	logs.Info(logsign, "[txmessage]", string(result))
	return string(result)
}

//腾讯语音提醒接口
type TXphonecall struct {
	Ext       string   `json:"ext"`
	Tpl_id    int      `json:"tpl_id"`
	Params    []string `json:"params"`
	Playtimes int      `json:"playtimes"`
	Sig       string   `json:"sig"`
	Tel       Mobiles  `json:"tel"`
	Time      int      `json:"time"`
}

//腾讯语音子程序
func PostTXphonecall(Messages string, PhoneNumbers, logsign string) string {
	open := beego.AppConfig.String("open-txdh")
	if open != "1" {
		logs.Info(logsign, "[txphonecall]", "腾讯语音接口未配置未开启状态,请先配置open-txdh为1")
		return "腾讯语音接口未配置未开启状态,请先配置open-txdh为1"
	}
	strAppKey := beego.AppConfig.String("TXY_DH_phonecallappkey")
	sdkappid := beego.AppConfig.String("TXY_DH_phonecallsdkappid")
	tpl_id, _ := beego.AppConfig.Int("TXY_DH_phonecalltpl_id")
	//腾讯短信接口算法部分
	TXmobile := Mobiles{}
	mobiles := strings.Split(PhoneNumbers, ",")
	for _, m := range mobiles {
		TXmobile.Mobile = m
		TXmobile.Nationcode = "86"
		strRand := "7226249334"
		strTime := strconv.FormatInt(time.Now().Unix(), 10)
		intTime, _ := strconv.Atoi(strTime)
		sig := getSha256Code("appkey=" + strAppKey + "&random=" + strRand + "&time=" + strTime + "&mobile=" + m)
		TXurl := "https://cloud.tim.qq.com/v5/tlsvoicesvr/sendtvoice?sdkappid=" + sdkappid + "&random=" + strRand
		u := TXphonecall{
			Ext:       logsign,
			Tpl_id:    tpl_id,
			Params:    []string{Messages},
			Playtimes: 2,
			Sig:       sig,
			Tel:       TXmobile,
			Time:      intTime,
		}
		res := PhoneCallPost(TXurl, u, logsign)
		logs.Info(logsign, "[txphonecall]", res)
	}
	model.AlertToCounter.WithLabelValues("txdh", Messages, PhoneNumbers).Add(1)
	return PhoneNumbers + " Called Over."
}

func PhoneCallPost(url string, u TXphonecall, logsign string) string {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	logs.Info(logsign, "[txphonecall]", b)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Post(url, "application/json", b)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	return string(result)
}
