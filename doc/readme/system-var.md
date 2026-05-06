# 自定义模版参数说明
----------------------------------------------------------------------

#### 1.钉钉机器人、企业微信机器人均已经支持@某人的功能。使用时，需要在Url中加入`&at=1539510xxxx`；如需添加多个 `@` 目标，用`,`号分隔即可。
> 此处需注意：钉钉@使用的是**手机号码**，企业微信机器人@使用的是**用户帐号**。

示例：
```shell
http://[prometheusalert_url]:8080/prometheusalert?type=dd&tpl=prometheus-dd&ddurl=https://oapi.dingtalk.com/robot/send?access_token=xxxx&at=1539510xxxx
```


#### 2.url参数中 `ddurl、wxurl、fsurl、phone、email、wxuser、wxparty、wxtag、groupid `等可不写，如不写这些参数，则会默认去读取配置文件中的对应参数发送消息。

示例：
```shell
http://[prometheusalert_url]:8080/prometheusalert?type=dd&tpl=prometheus-dd
```


#### 3.url参数中支持参数 `rr=true`， 该参数为开启随机轮询，目前仅针对 `ddurl`，`fsurl`，`wxurl` 有效，默认情况下如果上述Url配置的是多个地址，则多个地址全部发送，如开启该选项，则从多个地址中随机取一个地址发送，主要是为了避免消息发送频率过高导致触发部分机器人拦截消息。

示例：
```shell
http://[prometheusalert_url]:8080/prometheusalert?type=dd&tpl=prometheus-dd&ddurl=https://oapi.dingtalk.com/robot/send?access_token=xxxx,https://oapi.dingtalk.com/robot/send?access_token=xxxxxx,https://oapi.dingtalk.com/robot/send?access_token=xxxxxx&rr=true
```


#### 4.url参数新增 `split=true`，该参数仅针对Prometheus告警消息有效，作用是将Prometheus分组消息拆分成单条发送。默认开启，如果Prometheus一次告警附带的同分组的告警消息条数过多，可能会导致告警消息体过大。如需关闭请在url中加入 `split=false`

> 注意：此参数如设置为`split=false`，则PrometheusAlert web页面的**路由**和**告警记录**等功能将自动关闭，请谨慎。

示例：
```shell
http://[prometheusalert_url]:8080/prometheusalert?type=dd&tpl=prometheus-dd&ddurl=https://oapi.dingtalk.com/robot/send?access_token=xxxx,https://oapi.dingtalk.com/robot/send?access_token=xxxxxx,https://oapi.dingtalk.com/robot/send?access_token=xxxxxx&rr=true&split=false
```


#### 5.url参数新增 `emailtitle=运维监控中心`，该参数仅针对email有效，作用是替换app.conf中的 `Email_title` 邮件标题配置，实现动态定义邮件标题

> 注意：此参数如设置为`split=false`，则PrometheusAlert web页面的**路由**和**告警记录**等功能将自动关闭，请谨慎。

示例：
```shell
http://[prometheusalert_url]:8080/prometheusalert?type=email&tpl=prometheus-email&email=xxxx@xxx.com&emailtitle=运维监控中心
```


#### 6.自定义模板使用的是go语言的 `template` 模版，可以参考默认模版的一些语法来进行自定义。
Issue： [通用自定义模版专用：欢迎共享大家的自定义模版，方便其他人也可以直接使用 #30](https://github.com/feiyu563/PrometheusAlert/issues/30) 
 有大家共享的自定义模板配置，可以参考。

#### 7.模版数据等信息均存储在程序目录的下的`db/PrometheusAlertDB.db`中。

#### 8.关于优先级问题：路由功能 > URL参数 > `app.conf`

#### 9. url 参数中新增有 `usealertname=true` 参数，用于使用 `AlertName` 作为渠道通知的消息预览文字
默认情况下告警消息的预览文字为 `PrometheusAlert` 固定值，利用此参数可以将**告警名称**作为预览文字显示在未读消息中。
```shell
http://[prometheusalert_url]:8080/prometheusalert?type=dd&tpl=prometheus-dd&usealertname=true
```

该特性适用于如下渠道：
- [x] dingding(钉钉)
- [x] feishu(飞书)
