# 自定义模版参数说明
----------------------------------------------------------------------

#### 1.钉钉机器人、企业微信机器人均已经支持@某人的功能。使用时，需要在Url中加入`&at= 1539510xxxx`；如需添加多个@目标，用,号分割即可。此处需注意：钉钉@使用的是手机号码，企业微信机器人@使用的是用户帐号。

`示例：http://[prometheusalert_url]:8080/prometheusalert?type=dd&tpl=prometheus-dd&ddurl=https://oapi.dingtalk.com/robot/send?access_token=xxxx&at= 1539510xxxx`


#### 2.url参数中 `ddurl、wxurl、fsurl、phone、email、wxuser、wxparty、wxtag、groupid `等可不写，如不写这些参数，则会默认去读取配置文件中的对应参数发送消息。

`示例：http://[prometheusalert_url]:8080/prometheusalert?type=dd&tpl=prometheus-dd`


#### 3.url参数中支持参数 `rr=true`， 该参数为开启随机轮询，目前仅针对ddurl，fsurl，wxurl有效，默认情况下如果上述Url配置的是多个地址，则多个地址全部发送，如开启该选项，则从多个地址中随机取一个地址发送，主要是为了避免消息发送频率过高导致触发部分机器人拦截消息。

`示例：http://[prometheusalert_url]:8080/prometheusalert?type=dd&tpl=prometheus-dd&ddurl=https://oapi.dingtalk.com/robot/send?access_token=xxxx,https://oapi.dingtalk.com/robot/send?access_token=xxxxxx,https://oapi.dingtalk.com/robot/send?access_token=xxxxxx&rr=true`


#### 4.url参数新增 `split=true`，该参数仅针对Prometheus告警消息有效，作用是将Prometheus分组消息拆分成单条发送。默认开启，如果Prometheus一次告警附带的同分组的告警消息条数过多，可能会导致告警消息体过大。如需关闭请在url中加入split=false

注意：此参数如设置为`split=false`，则PrometheusAlert web页面的路由和告警记录等功能将自动关闭，请谨慎。

`示例：http://[prometheusalert_url]:8080/prometheusalert?type=dd&tpl=prometheus-dd&ddurl=https://oapi.dingtalk.com/robot/send?access_token=xxxx,https://oapi.dingtalk.com/robot/send?access_token=xxxxxx,https://oapi.dingtalk.com/robot/send?access_token=xxxxxx&rr=true&split=false`


#### 5.自定义模板使用的是go语言的template模版，可以参考默认模版的一些语法来进行自定义。

#### 6.模版数据等信息均存储在程序目录的下的`db/PrometheusAlertDB.db`中。

#### 7.关于优先级问题：路由功能 > URL参数 > app.conf