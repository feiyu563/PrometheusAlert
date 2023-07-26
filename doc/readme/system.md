# 告警系统接入PrometheusAlert配置

----------------------------------------

PrometheusAlert的原理就是通过自定义模版接口`/prometheusalert`接收各种告警系统或者任何带有WebHook功能的系统发来的消息，然后将收到的消息经过自定义模板渲染成消息文本，最终转发给不同的接收目标。

一般情况下如果使用的是钉钉，企业微信、飞书等机器人作为接收目标的，可以不去配置PrometheusAlert的配置文件app.conf；但是如果需要使用如短信，电话，邮箱等功能，则需要先配置好app.conf中的相关配置项方可使用。

#### 系统接入PrometheusAlert流程参考

- 1.安装好PrometheusAlert 参考：[安装部署PrometheusAlert](base-install.md)
- 2.配置 app.conf [可选] 参考：[【 app.conf 默认参数配置】](conf.md)
- 3.配置告警系统接入PrometheusAlert

#### PrometheusAlert处理告警消息原理

`XXX-WebHook` --> `POST-JSON` --> `/prometheusalert?type=dd&tpl=prometheus-dingding&ddurl=https://oapi.dingtalk.com/robot/send?access_token=xxxx` --> `PrometheusAlert通过tpl模版prometheus-dingding渲染收到的JSON` --> `将渲染后的消息文本发送给https://oapi.dingtalk.com/robot/send?access_token=xxxx` --> `钉钉机器人完成告警`
