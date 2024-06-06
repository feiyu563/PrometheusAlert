package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var (
	r        map[string]interface{}
	esClient *elasticsearch.Client
)

// AlertES is a alert structure used for serializing data in ES.
type AlertES struct {
	Alertname   string    `json:"alertname"`
	Status      string    `json:"status"`
	Instance    string    `json:"instance"`
	Level       string    `json:"level"`
	Labels      string    `json:"labels"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	StartsAt    string    `json:"startsAt"`
	EndsAt      string    `json:"endsAt"`
	Created     time.Time `json:"created"`
	Cloud       string    `json:"cloud"`
	Hostgroup   string    `json:"hostgroup"`
	Hostnmae    string    `json:"hostname"`
}

func init() {
	alertToES := beego.AppConfig.DefaultString("alert_to_es", "0")
	if alertToES == "1" {
		esURL := beego.AppConfig.Strings("to_es_url")
		esUser := beego.AppConfig.DefaultString("to_es_user", "")
		esPwd := beego.AppConfig.DefaultString("to_es_pwd", "")

		var err error
		cfg := elasticsearch.Config{
			Addresses: esURL,
			Username:  esUser,
			Password:  esPwd,
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
			logs.Error("[elasticsearch] Error indexing alert document: %s", err)
			return
		}
	}

	logs.Info("[elasticsearch] alert document indexed successfully in index %s", index)
}
