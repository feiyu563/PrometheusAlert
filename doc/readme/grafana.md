 **Grafana 接入配置**

首先使用管理员或者具有告警配置权限的帐号登录进Grafana管理页面，登录后进入notification channels配置。

![grafana1](../addchannel.png)
##### 注意此处的地址需要去PrometheusAlert的模版页面获取，如发送给钉钉的grafana模版`http://[prometheusalert_url]:8080/prometheusalert?type=dd&tpl=grafana-dd&ddurl=钉钉机器人地址&at=18888888888`
![grafana2](../addchannel2.png)
PrometheusAlert的模版页面
![grafana4](../grafanaalert3.png)
配置完成后保存即可.继续进行告警消息配置,选择任意一个折线图,点击编辑,进入aler配置,配置参考下图:

![grafana3](../grafanaalert1.png)

![grafana4](../grafanaalert2.png)

最终告警效果:

![grafana5](../grafana.png)