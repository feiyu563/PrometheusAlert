 **Grafana 接入配置**

首先使用管理员或者具有告警配置权限的帐号登录进Grafana管理页面，登录后进入notification channels配置。

![grafana1](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/addchannel.png)

注意这里的url地址填写上自己部署所在的url

- `grafana固定模版接口`

```
/grafana/dingding  处理Grafana告警消息转发到钉钉接口，可选参数(ddurl)
/grafana/weixin    处理Grafana告警消息转发到微信接口，可选参数(wxurl)
/grafana/feishu    处理Grafana告警消息转发到飞书接口，可选参数(fsurl)
/grafana/txdx      处理Grafana告警消息转发到腾讯云短信接口，可选参数(phone)
/grafana/txdh      处理Grafana告警消息转发到腾讯云电话接口，可选参数(phone)
/grafana/hwdx      处理Grafana告警消息转发到华为云短信接口，可选参数(phone)
/grafana/bddx      处理Grafana告警消息转发到百度云短信接口，可选参数(phone)
/grafana/alydx     处理Grafana告警消息转发到阿里云短信接口，可选参数(phone)
/grafana/rlydh     处理Grafana告警消息转发到容联云电话接口，可选参数(phone)
/grafana/email     处理Grafana告警消息转发到email接口，可选参数(email)
/grafana/tg        处理Grafana告警消息转发到telegram接口
/grafana/workwechat处理Grafana告警消息转发到企业微信应用接口
```

关于接口说明：grafana的所有接口均支持传参,如直接使用接口，未在接口后加入参数，默认会优先使用配置文件中的参数作为告警渠道的配置。如果接口中加入了参数，将默认使用url中的参数作为告警渠道的配置。如下：

```
/grafana/dingding?ddurl=https://oapi.dingtalk.com/robot/send?access_token=xxxxx
/grafana/weixin?wxurl=https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxx
/grafana/feishu?fsurl=https://open.feishu.cn/open-apis/bot/hook/xxxxxxxxx
/grafana/txdx?phone=15395105573
/grafana/txdh?phone=15395105573
/grafana/hwdx?phone=15395105573
/grafana/bddx?phone=15395105573
/grafana/alydx?phone=15395105573
/grafana/alydh?phone=15395105573
/grafana/rlydh?phone=15395105573
/grafana/email?email=123@qq.com
/grafana/tg
/grafana/workwechat
```

![grafana2](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/addchannel2.png)

配置完成后保存即可.继续进行告警消息配置,选择任意一个折线图,点击编辑,进入aler配置,配置参考下图:

![grafana3](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/grafanaalert1.png)

![grafana4](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/grafanaalert2.png)

最终告警效果:

![grafana5](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/grafana.png)