 ## graylog接入配置
 
首先使用管理员或者具有告警配置权限的帐号登录进Graylog日志系统后台，打开Graylog管理页面并进入Alerts配置。

![graylog1](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/graylog1.png)

点击```Add new notification```创建新的告警通道,选择如下图配置:

![graylog2](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/graylog2.png)

在弹出的窗口中填入名称和对应的PrometheusAlert的接口地址即可,配置参考下图:

- `graylog2固定模版接口`

```
特别说明: graylog2接口针对 graylog版本 <= 3.0.x

/graylog2/dingding  处理Graylog2告警消息转发到钉钉接口，可选参数(ddurl)
/graylog2/weixin    处理Graylog2告警消息转发到微信接口，可选参数(wxurl)
/graylog2/feishu    处理Graylog2告警消息转发到飞书接口，可选参数(fsurl)
/graylog2/txdx      处理Graylog2告警消息转发到腾讯云短信接口，可选参数(phone)
/graylog2/txdh      处理Graylog2告警消息转发到腾讯云电话接口，可选参数(phone)
/graylog2/hwdx      处理Graylog2告警消息转发到华为云短信接口，可选参数(phone)
/graylog2/bddx      处理Graylog2告警消息转发到百度云短信接口，可选参数(phone)
/graylog2/alydx     处理Graylog2告警消息转发到阿里云短信接口，可选参数(phone)
/graylog2/alydh     处理Graylog2告警消息转发到阿里云电话接口，可选参数(phone)
/graylog2/rlydh     处理Graylog2告警消息转发到容联云电话接口，可选参数(phone)
/graylog2/email     处理Graylog2告警消息转发到Email接口，可选参数(email)
/graylog2/tg        处理Graylog2告警消息转发到telegram接口
/graylog2/workwechat处理Graylog2告警消息转发到企业微信应用接口
```

- `graylog3固定模版接口`

```
特别说明: graylog3接口针对 graylog版本 >= 3.1.x

/graylog3/dingding  处理Graylog3告警消息转发到钉钉接口，可选参数(ddurl)
/graylog3/weixin    处理Graylog3告警消息转发到微信接口，可选参数(wxurl)
/graylog3/feishu    处理Graylog3告警消息转发到飞书接口，可选参数(fsurl)
/graylog3/txdx      处理Graylog3告警消息转发到腾讯云短信接口，可选参数(phone)
/graylog3/txdh      处理Graylog3告警消息转发到腾讯云电话接口，可选参数(phone)
/graylog3/hwdx      处理Graylog3告警消息转发到华为云短信接口，可选参数(phone)
/graylog3/bddx      处理Graylog3告警消息转发到百度云短信接口，可选参数(phone)
/graylog3/alydx     处理Graylog3告警消息转发到阿里云短信接口，可选参数(phone)
/graylog3/alydh     处理Graylog3告警消息转发到阿里云电话接口，可选参数(phone)
/graylog3/rlydh     处理Graylog3告警消息转发到容联云电话接口，可选参数(phone)
/graylog3/email     处理Graylog3告警消息转发到Email口，可选参数(email)
/graylog3/tg        处理Graylog3告警消息转发到telegram口
/graylog3/workwechat处理Graylog3告警消息转发到企业微信应用接口
```

关于接口说明：graylog的所有接口均支持传参,如直接使用接口，未在接口后加入参数，默认会优先使用配置文件中的参数作为告警渠道的配置。如果接口中加入了参数，将默认使用url中的参数作为告警渠道的配置。如下：

```
/graylog3/dingding?ddurl=https://oapi.dingtalk.com/robot/send?access_token=xxxxx
/graylog3/weixin?wxurl=https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxx
/graylog3/feishu?fsurl=https://open.feishu.cn/open-apis/bot/hook/xxxxxxxxx
/graylog3/txdx?phone=15395105573
/graylog3/txdh?phone=15395105573
/graylog3/hwdx?phone=15395105573
/graylog3/alydx?phone=15395105573
/graylog3/alydh?phone=15395105573
/graylog3/rlydh?phone=15395105573
/graylog3/email?email=123@qq.com
/graylog3/tg
/graylog3/workwechat
```

![graylog3](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/graylog3.png)

配置完成后,点击```Test```测试下是否能够正常接收告警消息即可

最终告警效果:

![graylog4](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/graylog4.png)
