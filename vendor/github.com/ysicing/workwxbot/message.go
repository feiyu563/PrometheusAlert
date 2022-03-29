package workwxbot

import (
	"time"
)

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

type WxBotMessage struct {
	MsgType  string      `json:"msgtype"`
	BotText  BotText     `json:"text"`
	MarkDown BotMarkDown `json:"markdown"`
	Image    BotImage    `json:"image"`
	News     News        `json:"news"`
	File     Media       `json:"file"`
}

type BotText struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list,omitempty"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

type BotMarkDown struct {
	Content string `json:"content"`
}

type BotImage struct {
	Base64 string `json:"base64"`
	Md5    string `json:"md5"`
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

//Media 媒体内容
type Media struct {
	MediaID     string `json:"media_id"`
	Title       string `json:"title,omitempty"`       // 视频参数
	Description string `json:"description,omitempty"` // 视频参数
}

//Card 卡片
type TextCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Btntxt      string `json:"btntxt"`
}

//news 图文
type News struct {
	Articles []Article `json:"articles"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Picurl      string `json:"picurl"`
}

//mpnews 图文
type MpNews struct {
	Articles []MpArticle `json:"articles"`
}

type MpArticle struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	ContentSourceUrl string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
}

// 任务卡片
type TaskCard struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	TaskID      string    `json:"task_id"`
	Btn         []TaskBtn `json:"btn"`
}

type TaskBtn struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	ReplaceName string `json:"replace_name"`
	Color       string `json:"color"`
	IsBold      bool   `json:"is_bold"`
}

//Message 消息主体参数
type Message struct {
	ToUser  string `json:"touser"`
	ToParty string `json:"toparty"`
	ToTag   string `json:"totag"`
	MsgType string `json:"msgtype"`
	AgentID int64  `json:"agentid"`

	Text     Content  `json:"text"`
	Image    Media    `json:"image"`
	Voice    Media    `json:"voice"`
	Video    Media    `json:"video"`
	File     Media    `json:"file"`
	Textcard TextCard `json:"textcard"`
	News     News     `json:"news"`
	MpNews   MpNews   `json:"mpnews"`
	Markdown Content  `json:"markdown"`
	Taskcard TaskCard `json:"taskcard"`
	// EnableDuplicateCheck bool `json:"enable_duplicate_check"`  // 表示是否开启重复消息检查，0表示否，1表示是，默认0
	// DuplicateCheckInterval int `json:"duplicate_check_interval"` // 表示是否重复消息检查的时间间隔，默认1800s，最大不超过4小时
}
