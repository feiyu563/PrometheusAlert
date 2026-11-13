package elastic

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var (
	r           map[string]interface{}
	esClient    *elasticsearch.Client
	PCstTime, _ = beego.AppConfig.Int("prometheus_cst_time")
)

// AlertES is a alert structure used for serializing data in ES.
// 将 created 定义为 es 默认的 @timestamp 时间戳字段
type AlertES struct {
	Source          string    `json:"source"`           // 消息来源 (Prometheus, Zabbix, GitLab等)
	Channel         string    `json:"channel"`          // 转发渠道 (wx, dd, fs, email, alydh等)
	Status          string    `json:"status"`           // 发送状态 ("success", "failed")
	Result          string    `json:"result"`           // 详细结果 (返回的response或Error信息)
	Summary         string    `json:"summary"`          // 告警摘要
	OriginalPayload string    `json:"original_payload"` // 原始消息内容 (RequestBody)
	Timestamp       time.Time `json:"@timestamp"`       // ES 默认时间戳
}

func init() {
	alertToES := beego.AppConfig.DefaultString("alert_to_es", "0")
	if alertToES == "1" {
		esURL := beego.AppConfig.Strings("to_es_url")
		esUser := beego.AppConfig.DefaultString("to_es_user", "")
		esPwd := beego.AppConfig.DefaultString("to_es_pwd", "")

		var err error
		// skip es https ca checks
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		cfg := elasticsearch.Config{
			Addresses: esURL,
			Username:  esUser,
			Password:  esPwd,
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		}
		esClient, err = elasticsearch.NewClient(cfg)

		if err != nil {
			logs.Error("[elasticsearch] Error creating the client: %s", err)
		}

		// 如果 ES 不可用或连接异常等问题，获取 ES 集群信息会 panic 导致程序崩溃，因此使用 return 提前退出
		res, err := esClient.Info()
		if err != nil {
			logs.Critical("[elasticsearch] Error getting response: %s from cluster %s", err, esURL)
			return
		}
		// 避免 res 为空（连接错误）时 panic 导致崩溃
		if res != nil {
			defer res.Body.Close()
			if res.IsError() {
				logs.Error("[elasticsearch] Connection error: %s", res.String())
				return
			}
			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
				logs.Error("[elasticsearch] Error parsing the response body: %s", err)
			}
			logs.Info("[elasticsearch] Successfully connected to ES Server Version: %s", r["version"].(map[string]interface{})["number"])
		}
	}
}

func Insert(index string, alert AlertES) {
	doc, err := json.Marshal(alert)
	if err != nil {
		logs.Error("[elasticsearch] error marshaling document: %w", err)
		return
	}

	req := esapi.IndexRequest{
		Index:   index,
		Body:    bytes.NewReader(doc),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		logs.Critical("[elasticsearch] Error getting response: %s", err)
		return
	}
	if res != nil {
		defer res.Body.Close()
		if res.IsError() {
			logs.Error("[elasticsearch] Error indexing alert document: %s", res.String())
			return
		}
	}

	logs.Info("[elasticsearch] alert document indexed successfully in index %s", index)
}
