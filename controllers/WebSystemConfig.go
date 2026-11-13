package controllers

import (
	"PrometheusAlert/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func (c *MainController) SystemConfig() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsSystemConfig"] = true
	c.Data["IsTemplateMenu"] = false
	c.Data["IsAlertManageMenu"] = false
	c.TplName = "system_config.html"

	// Gather config keys
	keys := []string{
		"title", "defaultphone", "proxy",
		"open-dingding", "ddurl", "dd_isatall", "open-dingding-secret",
		"open-weixin", "wxurl",
		"open-feishu", "fsurl",
		"open-alydx", "ALY_DX_AccessKeyId", "ALY_DX_AccessSecret", "ALY_DX_SignName", "ALY_DX_Template",
		"open-alydh", "ALY_DH_AccessKeyId", "ALY_DH_AccessSecret", "ALY_DX_CalledShowNumber", "ALY_DH_TtsCode",
		"open-txdx", "TXY_DX_appkey", "TXY_DX_tpl_id", "TXY_DX_sdkappid", "TXY_DX_sign",
		"open-txdh", "TXY_DH_phonecallappkey", "TXY_DH_phonecalltpl_id", "TXY_DH_phonecallsdkappid",
		"open-email", "Email_host", "Email_port", "Email_user", "Email_password", "Email_title", "Default_emails",
		"open-rlydh", "RLY_URL", "RLY_ACCOUNT_SID", "RLY_ACCOUNT_TOKEN", "RLY_APP_ID",
		"open-7moordx", "7MOOR_ACCOUNT_ID", "7MOOR_ACCOUNT_APISECRET", "7MOOR_DX_TEMPLATENUM",
		"open-7moordh", "7MOOR_WEBCALL_SERVICENO", "7MOOR_WEBCALL_VOICE_VAR",
		"open-tg", "TG_TOKEN", "TG_MODE_CHAN", "TG_USERID", "TG_CHANNAME", "TG_PARSE_MODE",
		"open-workwechat", "WorkWechat_CropID", "WorkWechat_AgentID", "WorkWechat_AgentSecret", "WorkWechat_ToUser", "WorkWechat_ToParty", "WorkWechat_ToTag",
		"open-baidudx", "BDY_DX_AK", "BDY_DX_SK", "BDY_DX_ENDPOINT", "BDY_DX_TEMPLATE_ID", "TXY_DX_SIGNATURE_ID",
		"open-ruliu", "BDRL_URL", "BDRL_ID",
		"open-bark", "BARK_URL", "BARK_KEYS", "BARK_COPY", "BARK_ARCHIVE", "BARK_GROUP",
		"open-voice", "VOICE_IP", "VOICE_PORT",
		"open-feishuapp", "FEISHU_APPID", "FEISHU_APPSECRET", "AT_USER_ID",
		"open-kafka", "kafka_server", "kafka_topic", "kafka_key",
		"AlertRecord", "RecordLive", "RecordLiveDay",
		"alert_to_es", "to_es_url", "to_es_user", "to_es_pwd",
		"maxIdleConns",
	}

	configs := make(map[string]string)
	for _, key := range keys {
		// This automatically falls back to base config because we hooked beego.AppConfig!
		configs[key] = beego.AppConfig.String(key)
	}

	c.Data["Configs"] = configs
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

func (c *MainController) SaveSystemConfig() {
	if !CheckAccount(c.Ctx) {
		c.Data["json"] = "unauthorized"
		c.ServeJSON()
		return
	}

	var req map[string]string
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	for k, v := range req {
		models.SetCacheConfig(k, v)
	}

	c.Data["json"] = "ok"
	c.ServeJSON()
}
