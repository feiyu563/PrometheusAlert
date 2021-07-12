## PrometheusAlert接口说明

--------------------------------------

PrometheusAlert 目前提供以下几类接口，分别对应各自接入端，负责解析各自接口传入或者传出的消息。

- `prometheusalert自定义模版接口`

```
/prometheusalert?type=${type}&tpl=${template}&[ddurl=${ddur}][wxurl=${wxurl}][fsurl=${fsurl}][phone=${phonenumber}]   自定义模版接口，可通过Dashboard自定义模版后，支持任意WebHook接入
```

- `metrics接口`

```
/metrics           展示PrometheusAlert指标信息
```

- `prometheus固定模版接口`

```
/prometheus/alert   处理Prometheus告警消息转发到默认接口
/prometheus/router  处理Prometheus AlertManager router消息指定接收端接口
```

- `zabbix接口`

```
/zabbix/alert  处理Zabbix告警消息转发默认接口
```

- `grafana固定模版接口`

```
/grafana/phone     处理Grafana告警消息转发到腾讯云电话接口(v3.0版本将废弃)
/grafana/dingding  处理Grafana告警消息转发到钉钉接口
/grafana/weixin    处理Grafana告警消息转发到微信接口
/grafana/feishu    处理Grafana告警消息转发到飞书接口
/grafana/txdx      处理Grafana告警消息转发到腾讯云短信接口
/grafana/txdh      处理Grafana告警消息转发到腾讯云电话接口
/grafana/hwdx      处理Grafana告警消息转发到华为云短信接口
/grafana/alydx     处理Grafana告警消息转发到阿里云短信接口
/grafana/alydh     处理Grafana告警消息转发到阿里云电话接口
/grafana/rlydh     处理Grafana告警消息转发到容联云电话接口
/grafana/email     处理Grafana告警消息转发到Email接口
/grafana/bddx      处理Grafana告警消息转发到百度云短信接口
/grafana/tg        处理Grafana告警消息转发到telegram接口
/grafana/workwechat处理Grafana告警消息转发到企业微信应用接口
/grafana/ruliu     处理Grafana告警消息转发到百度Hi(如流)接口
```

- `graylog2固定模版接口`

```
特别说明: graylog2接口针对 graylog版本 <= 3.0.x

/graylog2/phone     处理Graylog2告警消息转发到腾讯云电话接口(v3.0版本将废弃)
/graylog2/dingding  处理Graylog2告警消息转发到钉钉接口
/graylog2/weixin    处理Graylog2告警消息转发到微信接口
/graylog2/feishu    处理Graylog2告警消息转发到飞书接口
/graylog2/txdx      处理Graylog2告警消息转发到腾讯云短信接口
/graylog2/txdh      处理Graylog2告警消息转发到腾讯云电话接口
/graylog2/hwdx      处理Graylog2告警消息转发到华为云短信接口
/graylog2/alydx     处理Graylog2告警消息转发到阿里云短信接口
/graylog2/alydh     处理Graylog2告警消息转发到阿里云电话接口
/graylog2/rlydh     处理Graylog2告警消息转发到容联云电话接口
/graylog2/email     处理Graylog2告警消息转发到Email接口
/graylog2/bddx       处理Graylog2告警消息转发到百度云短信接口
/graylog2/tg         处理Graylog2告警消息转发到telegram接口
/graylog2/workwechat 处理Graylog2告警消息转发到企业微信应用接口
/graylog2/ruliu      处理Graylog2告警消息转发到百度Hi(如流)接口
```

- `graylog3固定模版接口`

```
特别说明: graylog3接口针对 graylog版本 >= 3.1.x

/graylog3/phone     处理Graylog3告警消息转发到腾讯云电话接口(v3.0版本将废弃)
/graylog3/dingding  处理Graylog3告警消息转发到钉钉接口
/graylog3/weixin    处理Graylog3告警消息转发到微信接口
/graylog3/feishu    处理Graylog3告警消息转发到微信接口
/graylog3/txdx      处理Graylog3告警消息转发到腾讯云短信接口
/graylog3/txdh      处理Graylog3告警消息转发到腾讯云电话接口
/graylog3/hwdx      处理Graylog3告警消息转发到华为云短信接口
/graylog3/alydx     处理Graylog3告警消息转发到阿里云短信接口
/graylog3/alydh     处理Graylog3告警消息转发到阿里云电话接口
/graylog3/rlydh     处理Graylog3告警消息转发到容联云电话接口
/graylog3/email     处理Graylog3告警消息转发到Email接口
/graylog3/bddx      处理Graylog3告警消息转发到百度云短信接口
/graylog3/tg        处理Graylog3告警消息转发到telegram接口
/graylog3/workwechat处理Graylog3告警消息转发到企业微信应用接口
/graylog3/ruliu     处理Graylog3告警消息转发到百度Hi(如流)接口
```

- `语音短信回调接口`

```
/tengxun/status     处理腾讯云语音短信回调接口，负责失败后重试
```
