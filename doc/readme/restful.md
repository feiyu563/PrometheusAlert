## PrometheusAlert接口说明
--------------------------------------

PrometheusAlert 目前提供以下几类接口，分别对应各自接入端，负责解析各自接口传入或者传出的消息。

 - `prometheus接口`

```
/prometheus/alert   处理Prometheus告警消息转发到默认接口
/prometheus/router  处理Prometheus AlertManager router消息指定接收端接口
```

 - `zabbix接口`

```
/zabbix/alert  处理Zabbix告警消息转发默认接口
```

 - `grafana接口`

```
/grafana/phone     处理Grafana告警消息转发到腾讯云电话接口(v3.0版本将废弃)
/grafana/dingding  处理Grafana告警消息转发到钉钉接口
/grafana/weixin    处理Grafana告警消息转发到微信接口
/grafana/txdx      处理Grafana告警消息转发到腾讯云短信接口
/grafana/txdh      处理Grafana告警消息转发到腾讯云电话接口
/grafana/hwdx      处理Grafana告警消息转发到华为云短信接口
/grafana/alydx     处理Grafana告警消息转发到阿里云短信接口
/grafana/alydh     处理Grafana告警消息转发到阿里云电话接口
```

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

 - `语音短信回调接口`

```
/tengxun/status     处理腾讯云语音短信回调接口，负责失败后重试
```