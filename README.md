企业微信消息推送-Go

> Wechat-Work-Pusher Based On Golang

## 部署

#### 准备工作

1、注册企业微信： https://work.weixin.qq.com/ 注册完成后点击我的企业，获取企业ID(`cropId`)

2、进入应用管理，创建自建应用，获取`agentId`和`cropSecret`

3、进入我的企业-微信插件-邀请关注 扫码关注企业微信

#### config.json配置

1、将准备工作中获取到的cropId、cropSecret、agentId填入相应位置

2、receiver为默认接收者的微信号，对应通讯录中的帐号

3、填写token

#### Docker部署

    docker run -d -v /home/config.json:/root/config.json -p 9000:9000 --restart=always --name wechat-work-pusher myleo1/wechat-work-pusher

> 注意：替换/home/config.json 为config.json路径

## 使用方法

以Go为例，其他语言类似

#### 文本消息

```go
func Push2Wechat(to, msg string) {
   httpkit.Request(httpkit.Req{
      Method: http.MethodPost,
      //将127.0.0.1替换成自己的ip
      Url:    http://127.0.0.1:9000/wechat-work-pusher/msg,
      Header: map[string]string{
         "Cookie": fmt.Sprintf("session=%s", config中填写的token),
      },
      FormData: map[string]string{
         "to":      to,
         "content": msg,
      },
   })
}
```

#### 卡片消息

卡片消息参数具体请参考：https://work.weixin.qq.com/api/doc/90000/90135/90236#%E6%96%87%E6%9C%AC%E5%8D%A1%E7%89%87%E6%B6%88%E6%81%AF

```go
func Push2WechatCard(to, msg string) {
   httpkit.Request(httpkit.Req{
      Method: http.MethodPost,
      //将127.0.0.1替换成自己的ip
      Url:    http://127.0.0.1:9000/wechat-work-pusher/card,
      Header: map[string]string{
         "Cookie": fmt.Sprintf("session=%s", config中填写的token),
      },
      FormData: map[string]string{
         "to":      to,
         "title":		title
	 "description": fmt.Sprintf("<div class=\"gray\">%s\n</div> <div class=\"normal\">恭喜您%s~,学号[%s]打卡成功！\n</div><div class=\"highlight\">点击卡片进入云战役打卡官网查看详情~</div>", time.Now().Format(timekit.TimeLayoutYMD), name, id),
         "url": "https://www.google.com",
      },
   })
}
```
## 效果演示

![image](https://user-images.githubusercontent.com/66349676/111748431-7eadf380-88cb-11eb-8590-73e1414d98e6.png)

