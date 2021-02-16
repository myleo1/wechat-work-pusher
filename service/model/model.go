package model

import "github.com/mizuki1412/go-core-kit/class"

type TextMessage struct {
	ToUser  class.String `json:"touser"`
	MsgType class.String `json:"msgtype"`
	AgentId class.String `json:"agentid"`
	Text    struct {
		Content class.String `json:"content"`
	} `json:"text"`
}

type TextCardMessage struct {
	ToUser   class.String `json:"touser"`
	MsgType  class.String `json:"msgtype"`
	AgentId  class.String `json:"agentid"`
	TextCard struct {
		Title       class.String `json:"title"`
		Description class.String `json:"description"`
		Url         class.String `json:"url"`
		Detail      class.String `json:"btntxt"`
	} `json:"textcard"`
}
