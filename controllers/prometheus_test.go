package controllers

import (
	"encoding/json"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	"github.com/stretchr/testify/assert"
)

// Load app conf
func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestURLDeduplication(t *testing.T) {
	assert := assert.New(t)
	var want, result string

	want = ""
	result = URLDeduplication("")
	assert.Equal(want, result)

	want = "ag0-wxurl0,ag0-wxurl1"
	result = URLDeduplication("ag0-wxurl0,ag0-wxurl1")
	assert.Equal(want, result)

	want = "ag0-wxurl0,ag0-wxurl1"
	result = URLDeduplication("ag0-wxurl0,ag0-wxurl1,")
	assert.Equal(want, result)

	want = "ag0-wxurl0,ag0-wxurl1"
	result = URLDeduplication("ag0-wxurl0,ag0-wxurl1,ag0-wxurl0,ag0-wxurl1")
	assert.Equal(want, result)
}

func TestAlertgroup(t *testing.T) {
	assert := assert.New(t)
	var want, result map[string]string

	result = Alertgroup("")
	assert.Equal(want, result)

	open := beego.AppConfig.String("open-alertgroup")
	if open == "1" {
		result = Alertgroup("ag-demo")
		want = map[string]string{
			"wxurl":   "wxurl1,wxurl2",
			"ddurl":   "ddurl1",
			"fsurl":   "fsurl1",
			"phone":   "phone1,phone2",
			"email":   "email1",
			"groupid": "groupid1",
		}
		assert.Equal(want, result)
	}
}

func TestPrometheus(t *testing.T) {
	// A sample prometheus alertmanager JSON payload
	payload := `
	{
		"status": "firing",
		"alerts": [
			{
				"status": "firing",
				"labels": {
					"alertname": "TestAlert",
					"instance": "localhost",
					"level": "1",
					"severity": "warning",
					"job": "node_exporter",
					"hostgroup": "test",
					"hostname": "ecs01"
				},
				"annotations": {
					"description": "This is a test alert",
					"summary": "Test Alert Summary",
					"alertgroup": "sa,dev"
				},
				"startsAt": "2023-06-25T10:00:00Z",
				"endsAt": "2023-06-25T11:00:00Z",
				"generatorURL": "http://localhost/alerts"
			}
		],
		"externalURL": "http://localhost/prometheus"
	}`

	alert := Prometheus{}
	err := json.Unmarshal([]byte(payload), &alert)
	if err != nil {
		t.Errorf("Prometheus alertmanager json payload parse error: %s", err)
	}
}
