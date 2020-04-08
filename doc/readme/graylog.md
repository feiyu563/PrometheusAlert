 ## graylog接入配置
 
首先使用管理员或者具有告警配置权限的帐号登录进Graylog日志系统后台，打开Graylog管理页面并进入Alerts配置。

![graylog1](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/graylog1.png)
点击```Add new notification```创建新的告警通道,选择如下图配置:
![graylog2](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/graylog2.png)
在弹出的窗口中填入名称和对应的PrometheusAlert的接口地址即可,配置参考下图:

 - `graylog2接口`

```
特别说明: graylog2接口针对 graylog版本 <= 3.0.x

/graylog2/phone     处理Graylog2告警消息转发到腾讯云电话接口(v3.0版本将废弃)
/graylog2/dingding  处理Graylog2告警消息转发到钉钉接口
/graylog2/weixin    处理Graylog2告警消息转发到微信接口
/graylog2/txdx      处理Graylog2告警消息转发到腾讯云短信接口
/graylog2/txdh      处理Graylog2告警消息转发到腾讯云电话接口
/graylog2/hwdx      处理Graylog2告警消息转发到华为云短信接口
/graylog2/alydx     处理Graylog2告警消息转发到阿里云短信接口
/graylog2/alydh     处理Graylog2告警消息转发到阿里云电话接口
```

 - `graylog3接口`

```
特别说明: graylog3接口针对 graylog版本 >= 3.1.x

/graylog3/phone     处理Graylog3告警消息转发到腾讯云电话接口(v3.0版本将废弃)
/graylog3/dingding  处理Graylog3告警消息转发到钉钉接口
/graylog3/weixin    处理Graylog3告警消息转发到微信接口
/graylog3/txdx      处理Graylog3告警消息转发到腾讯云短信接口
/graylog3/txdh      处理Graylog3告警消息转发到腾讯云电话接口
/graylog3/hwdx      处理Graylog3告警消息转发到华为云短信接口
/graylog3/alydx     处理Graylog3告警消息转发到阿里云短信接口
/graylog3/alydh     处理Graylog3告警消息转发到阿里云电话接口
```

![graylog3](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/graylog3.png)
配置完成后,点击```Test```测试下是否能够正常接收告警消息即可

最终告警效果:

![graylog4](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/graylog4.png)
