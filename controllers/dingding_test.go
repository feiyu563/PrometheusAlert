package controllers

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDDMesage(t *testing.T) {
	// A sample dingding message json payload
	payload := `
	{
		"msgtype": "markdown",
		"markdown": {
			"title": "AlertPrometheus告警测试",
			"text": "**测试告警信息**\n"
		},
		"at": {
			"atMobiles": ["139xxxxxxxx", "159xxxxxxxx"],
			"isAtAll": false
		}
	}`
	message := DDMessage{}
	err := json.Unmarshal([]byte(payload), &message)
	assert.Nil(t, err)
}

func TestDingdingSign(t *testing.T) {
	var want, result string

	// url without parm secret
	withoutURL := "https://oapi.dingtalk.com/robot/send?access_token=XXX"
	want = withoutURL
	result = dingdingSign(withoutURL)
	assert.Equal(t, want, result)

	// url with parm secret
	withURL := "https://oapi.dingtalk.com/robot/send?access_token=XXX&secret=mysecret"
	result = dingdingSign(withURL)
	t.Logf("一个示例的加签的地址：%s", result)
}
