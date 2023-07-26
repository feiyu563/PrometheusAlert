# PrometheusAlert接口说明

--------------------------------------

PrometheusAlert 目前提供以下几类接口，分别对应各自接入端，负责解析各自接口传入或者传出的消息。

## prometheusalert自定义模版接口

```
/prometheusalert   #自定义模版接口，可通过Dashboard自定义模版后，支持任意WebHook接入
```

### Url参数解释：

- `type=?`：指定消息转发的目标类型，如钉钉、企业微信、飞书等；`该参数为必选参数`

```
目前支持的值：
dd 钉钉
wx 企业微信
workwechat 企业微信应用
fs 飞书
webhook WebHook
txdx 腾讯云短信
txdh 腾讯云电话
alydx 阿里云短信
alydh 阿里云电话
hwdx 华为云短信
bddx 百度云短信
rlydh 容联云电话
7moordx 七陌短信
7moordh 七陌语音电话
email Email
tg Telegram
rl 百度Hi(如流)
```

- `tpl=?`: 指定消息所使用的模版，如`prometheus-dd`(Prometheus针对钉钉的模板)；模版可以去PrometheusAlert 页面的模版管理-->自定义模板页面查看或新建;`该参数为必选参数`

- `ddurl=?`：指定PrometheusAlert发送消息的钉钉机器人地址，如需要多个地址可以通过`,`分割，该参数需要配合`type=dd`的模版使用;`该参数为可选参数，如未填写，则默认从app.conf中获取默认配置`

- `wxurl=?`：指定PrometheusAlert发送消息的企业微信机器人地址，如需要多个地址可以通过`,`分割，该参数需要配合`type=wx`的模版使用;`该参数为可选参数，如未填写，则默认从app.conf中获取默认配置`

- `fsurl=?`：指定PrometheusAlert发送消息的飞书机器人地址，如需要多个地址可以通过`,`分割，该参数需要配合`type=fs`的模版使用;`该参数为可选参数，如未填写，则默认从app.conf中获取默认配置`

- `phone=?`：指定PrometheusAlert发送消息的手机号，如需要多个号码可以通过`,`分割，该参数需要配合`type=txdx | hwdx | bddx | alydx | txdh | alydh | rlydh | 7moordx | 7moordh`的模版使用;`该参数为可选参数，如未填写，则默认从app.conf中获取默认配置`

- `email=?`：指定PrometheusAlert发送消息的email地址，如需要多个email可以通过`,`分割，该参数需要配合`type=email`的模版使用;`该参数为可选参数，如未填写，则默认从app.conf中获取默认配置`

- `groupid=?`：指定PrometheusAlert发送消息的groupid，该参数需要配合`type=rl`的模版使用;`该参数为可选参数，如未填写，则默认从app.conf中获取默认配置`

- `webhook=?`：指定PrometheusAlert发送消息的webhook，该参数需要配合`type=webhook`的模版使用;`该参数为可选参数`

- `at=?`：钉钉机器人、企业微信机器人开启@某人的功能，如需添加多个@目标，用`,`号分割即可。此处需注意：钉钉@使用的是手机号码，企业微信机器人@使用的是用户帐号。;`该参数为可选参数`

- `rr=?`：该参数为开启随机轮询，目前仅针对ddurl，fsurl，wxurl有效，默认情况下如果上述Url配置的是多个地址，则多个地址全部发送，如开启该选项，则从多个地址中随机取一个地址发送，主要是为了避免消息发送频率过高导致触发部分机器人拦截消息。;`该参数为可选参数`

- `split=?`：该参数仅针对Prometheus告警消息有效，作用是将Prometheus分组消息拆分成单条发送。默认开启，如果Prometheus一次告警附带的同分组的告警消息条数过多，可能会导致告警消息体过大。如需关闭请在url中加入split=false;`该参数为可选参数`

    注意：此参数如设置为`split=false`，则PrometheusAlert web页面的路由和告警记录等功能将自动关闭，请谨慎。


### metrics接口

```
/metrics       展示PrometheusAlert指标信息
```

### zabbix接口

```
/zabbix/alert  处理Zabbix告警消息转发接口
```

### 语音短信回调接口

```
/tengxun/status     处理腾讯云语音短信回调接口，负责失败后重试
```
