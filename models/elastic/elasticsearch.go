package elastic

import (
	"context"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	es7 "github.com/olivere/elastic/v7"
)

var esCli *es7.Client

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
}

// es mapping field datatypes
const mapping = `{
	"mappings": {
		"prometheusalert": {
			"properties": {
				"alertname":    {"type": "keyword"},
				"status":       {"type": "keyword"},
				"instance":     {"type": "keyword"},
				"level":        {"type": "keyword"},
				"labels":        {"type": "keyword"},
				"startsAt":      {"type": "date"},
				"endsAt":       {"type": "date"},
				"created":      {"type": "date"},
				"summary":      {"type": "text", "store": true, "fielddata": true},
				"description":  {"type": "text", "store": true, "fielddata": true}
			}
		}
	}
}`

func init() {
	alertToES := beego.AppConfig.DefaultString("alert_to_es", "0")
	if alertToES == "1" {
		esURL := beego.AppConfig.Strings("to_es_url")
		esUser := beego.AppConfig.DefaultString("to_es_user", "")
		esPwd := beego.AppConfig.DefaultString("to_es_pwd", "")

		err := NewESClient(esURL, esUser, esPwd)
		if err != nil {
			logs.Error("[elasticsearch] Connecting to es %s error: %s", esURL, err)
		}
	}
}

// NewESClient creates a new es client.
func NewESClient(url []string, user, pwd string) error {
	ctx := context.Background()
	var err error
	esCli, err = es7.NewClient(
		es7.SetURL(url...),
		es7.SetBasicAuth(user, pwd),
		es7.SetSniff(false),
		es7.SetHealthcheck(false),
	)
	if err != nil {
		return err
	}

	info, code, err := esCli.Ping(url[0]).Do(ctx)
	if err != nil {
		return err
	}
	logs.Info("[elasticsearch] ES returned with code: %d and version: %s", code, info.Version.Number)

	return nil
}

// Insert writes prometheus alert to es.
func Insert(index string, value interface{}) {
	ctx := context.Background()
	/*
		exists, err := esCli.IndexExists(index).Do(ctx)
		if err != nil {
			logs.Error("[elasticsearch] ES index: %s is not exist: %s", index, err)
		}
		if !exists {
			createIndex, err := esCli.CreateIndex(index).Do(ctx)
			if err != nil {
				logs.Error("[elasticsearch] Create es index: %s error: %s", index, err)
			}
			if !createIndex.Acknowledged {
				// Not acknowledge
			}
		}
	*/

	res, err := esCli.Index().
		Index(index).
		Type("prometheusalert").
		BodyJson(value).
		Do(ctx)
	if err != nil {
		logs.Error("[elasticsearch] Index a prometheusalert alert to es error: %s", err)
	} else {
		logs.Info("[elasticsearch] Index a prometheusalert alert id: %s to index: %s succesful.", res.Id, res.Index)
	}
}
