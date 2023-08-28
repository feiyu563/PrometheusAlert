package controllers

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/astaxie/beego"
	"github.com/stretchr/testify/assert"
)

func TestURLDeduplication(t *testing.T) {
	assert := assert.New(t)
	var want, result string

	want = ""
	result = URLDeduplication("")
	assert.Equal(want, result)

	want = "sa-wxurl0"
	result = URLDeduplication("sa-wxurl0")
	assert.Equal(want, result)

	want = "sa-wxurl0"
	result = URLDeduplication("sa-wxurl0,")
	assert.Equal(want, result)

	want = "sa-wxurl0"
	result = URLDeduplication(",sa-wxurl0")
	assert.Equal(want, result)

	want = "sa-wxurl0"
	result = URLDeduplication(",sa-wxurl0,")
	assert.Equal(want, result)

	want = "sa-wxurl0"
	result = URLDeduplication("sa-wxurl0, ")
	assert.Equal(want, result)

	want = "sa-wxurl0,sa-wxurl1"
	result = URLDeduplication("sa-wxurl0,sa-wxurl1")
	assert.Equal(want, result)

	want = "sa-wxurl0,sa-wxurl1"
	result = URLDeduplication("sa-wxurl0, sa-wxurl1")
	assert.Equal(want, result)

	want = "sa-wxurl0,sa-wxurl1"
	result = URLDeduplication("sa-wxurl0,sa-wxurl1, ")
	assert.Equal(want, result)

	want = "sa-wxurl0,sa-wxurl1"
	result = URLDeduplication("sa-wxurl0,sa-wxurl1,sa-wxurl0,sa-wxurl1")
	assert.Equal(want, result)
}

func TestAlertgroup(t *testing.T) {
	assert := assert.New(t)
	var want, result map[string]string

	// tempFile is /tmp/app*.conf
	tempFile, err := os.CreateTemp("", "app*.conf")
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())
	assert.Nil(err)

	// mock app.conf
	configData := `
		[sa]
		wxurl=wxurl1
		ddurl=ddurl1,
		fsurl=,fsurl1
		email=,email1,
		phone=phone1,phone1
		groupid=groupid1,groupid1,
		[ops]
		wxurl=wxurl1,wxurl2
		ddurl=ddurl1,ddurl2,
		fsurl=,fsurl1,fsurl2
		email=,email1,email2,
		phone=phone2
		groupid=groupid2
		[customtpl]
		wxurl=wxurl1,wxurl2
		ddurl=ddurl1,ddurl2
		fsurl=fsurl1,fsurl2
		email=email1,email2
		phone=phone1,phone2
		groupid=groupid1,groupid2
		webhookurl=webhookurl1,webhookurl2
	`

	// Write mock conf into tempFile
	_, err = tempFile.Write([]byte(configData))
	assert.Nil(err)

	beego.InitBeegoBeforeTest(tempFile.Name())

	result = Alertgroup("")
	assert.Equal(want, result)

	want = map[string]string{
		"wxurl":      "wxurl1",
		"ddurl":      "ddurl1",
		"fsurl":      "fsurl1",
		"phone":      "phone1",
		"email":      "email1",
		"groupid":    "groupid1",
		"webhookurl": "",
	}
	result = Alertgroup("sa")
	assert.Equal(want, result)

	want = map[string]string{
		"wxurl":      "wxurl1,wxurl2",
		"ddurl":      "ddurl1,ddurl2",
		"fsurl":      "fsurl1,fsurl2",
		"email":      "email1,email2",
		"phone":      "phone1,phone2",
		"groupid":    "groupid1,groupid2",
		"webhookurl": "",
	}
	result = Alertgroup("sa,ops")
	assert.Equal(want, result)

	want = map[string]string{
		"wxurl":      "wxurl1,wxurl2",
		"ddurl":      "ddurl1,ddurl2",
		"fsurl":      "fsurl1,fsurl2",
		"email":      "email1,email2",
		"phone":      "phone1,phone2",
		"groupid":    "groupid1,groupid2",
		"webhookurl": "webhookurl1,webhookurl2",
	}
	result = Alertgroup("customtpl")
	assert.Equal(want, result)
}

func TestPrometheusJSON(t *testing.T) {
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
	assert.Nil(t, err)
}

func TestCheckURL(t *testing.T) {
	assert := assert.New(t)
	var want, result string

	want = ""
	result = checkURL("", "")
	assert.Equal(want, result)

	want = "url1"
	result = checkURL("url1", "", "")
	assert.Equal(want, result)

	want = "url2"
	result = checkURL("", "url2", "")
	assert.Equal(want, result)

	want = "url3"
	result = checkURL("", "", "url3")
	assert.Equal(want, result)

	want = "url1"
	result = checkURL("url1", "url2", "url3")
	assert.Equal(want, result)
}
