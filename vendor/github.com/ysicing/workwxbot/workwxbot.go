package workwxbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Roboter is the interface implemented by Robot that can send multiple types of messages.
type Roboter interface {
	SendMarkdown(MsgType string, ConfigID string, Content string, MentionedList string) error
}

// Robot represents a workwxbot custom robot that can send messages to groups.
type Robot struct {
	Webhook string
}

// NewRobot returns a roboter that can send messages.
func NewRobot(webhook string) Roboter {
	return Robot{Webhook: webhook}
}

// SendMarkdown send a markdown type message.
func (r Robot) SendMarkdown(MsgType string, ConfigID string, Content string, MentionedList string) error {
	return r.send(&BotMessage{
		MsgType:       MsgType,
		ProgramType:   programType,
		IsSendNow:     isSendNow,
		ConfigID:      ConfigID,
		Content:       Content,
		MentionedList: MentionedList,
	})
}

type workRsp struct {
	Errcode int
	Errmsg  string
}

func (r Robot) send(msg interface{}) error {
	m, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	resp, err := http.Post(r.Webhook, "application/json", bytes.NewReader(m))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var wsp workRsp
	err = json.Unmarshal(data, &wsp)
	if err != nil {
		return err
	}
	if wsp.Errcode != 0 {
		return fmt.Errorf("wechatrobot send failed: %v", wsp.Errmsg)
	}

	return nil
}
