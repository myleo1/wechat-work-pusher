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
	"time"
	"wechat-work-pusher/constant"
)

type getToken struct {
	cropId string
	secret string
	once   sync.Once
}

type respToken struct {
	val    string
	expire time.Time
	lock   sync.RWMutex
}

var getTokenParams getToken
var respTokenParams respToken

func (th *respToken) set(val string) {
	th.lock.Lock()
	defer th.lock.Unlock()
	th.val = val
	th.expire = time.Now().Add(time.Minute * constant.RespTokenExpireMin)
}

func (th *respToken) isValid() bool {
	if !th.expire.IsZero() && time.Now().Before(th.expire) {
		return true
	}
	return false
}

func GetTokenFromWechat() string {
	if respTokenParams.isValid() {
		return respTokenParams.val
	}
	getTokenParams.once.Do(func() {
		getTokenParams.cropId = configkit.GetStringD(constant.ConfigKeyWorkCorpId)
		getTokenParams.secret = configkit.GetStringD(constant.ConfigKeyWorkCorpSecret)
	})
	resp, code := httpkit.Request(httpkit.Req{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", getTokenParams.cropId, getTokenParams.secret),
	})
	if code != http.StatusOK {
		panic(exception.New(fmt.Sprintf("get wechat token req err,code:%s", cast.ToString(code))))
	}
	if err := gjson.Get(resp, "errcode").Int(); err != 0 {
		panic(exception.New(fmt.Sprintf("get wechat token resp error,code:%s", cast.ToString(err))))
	}
	token := gjson.Get(resp, "access_token").String()
	respTokenParams.set(token)
	return token
}
