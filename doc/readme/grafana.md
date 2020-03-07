 **Grafana 接入配置**

打开grafana管理页面,登录后进入notification channels配置

![grafana1](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/addchannel.png)
注意这里的url地址填写上自己部署所在的url
 - `grafana接口`

```
/grafana/phone  腾讯云电话接口(v3.0版本将废弃)
/grafana/dingding  钉钉接口
/grafana/weixin  微信接口
/grafana/txdx  腾讯云短信接口
/grafana/txdh  腾讯云电话接口
/grafana/hwdx  华为云短信接口
/grafana/alydx  阿里云短信接口
/grafana/alydh  阿里云电话接口
```

![grafana2](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/addchannel2.png)
配置完成后保存即可.继续进行告警消息配置,选择任意一个折线图,点击编辑,进入aler配置,配置参考下图:
![grafana3](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/grafanaalert1.png)
![grafana4](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/grafanaalert2.png)

Notifications配置格式参考,支持配置多个告警机器人url:
```
告警消息内容&&url[钉钉或微信机器人url,钉钉或微信机器人url....]
```

最终告警效果:

![grafana5](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/grafana.png)