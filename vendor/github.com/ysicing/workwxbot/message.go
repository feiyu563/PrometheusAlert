package workwxbot

import "time"

const (
	//msgTypeMarkdown = "markdown"
	programType = "OA"
	isSendNow   = true
)

// BotMessage 机器人消息
type BotMessage struct {
	MsgType       string `json:"msgtype"`
	ProgramType   string `json:"program"`
	IsSendNow     bool   `json:"issendimmediately"`
	ConfigID      string `json:"configid"`
	Content       string `json:"content"`
	MentionedList string `json:"mentioned_list"`
}

// Err 微信返回错误
type Err struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

//AccessToken 微信企业号请求Token
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Err
	ExpiresInTime time.Time
}

//Client 微信企业号应用配置信息
type Client struct {
	CropID      string
	AgentID     int64
	AgentSecret string
	Token       AccessToken
}

//Result 发送消息返回结果
type Result struct {
	Err
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"infvalidparty"`
	InvalidTag   string `json:"invalidtag"`
}

//Content 文本消息内容
type Content struct {
	Content string `json:"content"`
}

//Message 消息主体参数
type Message struct {
	ToUser   string  `json:"touser"`
	ToParty  string  `json:"toparty"`
	ToTag    string  `json:"totag"`
	MsgType  string  `json:"msgtype"`
	AgentID  int64     `json:"agentid"`
	Markdown Content `json:"markdown"`
}
