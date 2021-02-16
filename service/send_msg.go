package service

import (
	"fmt"
	"github.com/mizuki1412/go-core-kit/class/exception"
	"github.com/mizuki1412/go-core-kit/library/httpkit"
	"github.com/mizuki1412/go-core-kit/service/configkit"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"net/http"
	"sync"
	"wechat-work-pusher/constant"
	"wechat-work-pusher/service/model"
)

type agent struct {
	val  string
	once sync.Once
}

var agentParams agent

func SendMsg(to, content string) {
	agentParams.once.Do(func() {
		agentParams.val = configkit.GetStringD(constant.ConfigKeyWorkAgentId)
	})
	msg := &model.TextMessage{}
	msg.AgentId.Set(agentParams.val)
	msg.MsgType.Set(constant.MessageText)
	msg.Text.Content.Set(content)
	if to != "" {
		msg.ToUser.Set(to)
	} else {
		msg.ToUser.Set(configkit.GetStringD(constant.ConfigKeyDefaultReceiver))
	}
	resp, code := httpkit.Request(httpkit.Req{
		Method:      http.MethodPost,
		Url:         fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", GetTokenFromWechat()),
		ContentType: httpkit.ContentTypeJSON,
		JsonData:    msg,
	})
	if code != http.StatusOK {
		panic(exception.New(fmt.Sprintf("send text to wechat req err,code:%s", cast.ToString(code))))
	}
	if err := gjson.Get(resp, "errcode").Int(); err != 0 {
		panic(exception.New(fmt.Sprintf("send text to wechat resp error,code:%s", cast.ToString(err))))
	}
}

func SendCardMsg(to, title, des, url string) {
	agentParams.once.Do(func() {
		agentParams.val = configkit.GetStringD(constant.ConfigKeyWorkAgentId)
	})
	msg := &model.TextCardMessage{}
	msg.AgentId.Set(agentParams.val)
	msg.MsgType.Set(constant.MessageCard)
	msg.TextCard.Title.Set(title)
	msg.TextCard.Description.Set(des)
	msg.TextCard.Url.Set(url)
	msg.TextCard.Detail.Set("详情")
	if to != "" {
		msg.ToUser.Set(to)
	} else {
		msg.ToUser.Set(configkit.GetStringD(constant.ConfigKeyDefaultReceiver))
	}
	resp, code := httpkit.Request(httpkit.Req{
		Method:      http.MethodPost,
		Url:         fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", GetTokenFromWechat()),
		ContentType: httpkit.ContentTypeJSON,
		JsonData:    msg,
	})
	if code != http.StatusOK {
		panic(exception.New(fmt.Sprintf("send textcard to wechat req err,code:%s", cast.ToString(code))))
	}
	if err := gjson.Get(resp, "errcode").Int(); err != 0 {
		panic(exception.New(fmt.Sprintf("send textcard to wechat resp error,code:%s", cast.ToString(err))))
	}
}
