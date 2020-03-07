PrometheusAlert全家桶接口配置说明
--------------------------------------

PrometheusAlert 暂提供以下几类接口,分别对应各自接入端

 - `prometheus接口`

```
/prometheus/alert  默认接口
/prometheus/router  AlertManager router指定接收端接口
```

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

 - `graylog2接口`

```
特别说明: graylog2接口针对 graylog版本 <= 3.0.x

/graylog2/phone  腾讯云电话接口(v3.0版本将废弃)
/graylog2/dingding  钉钉接口
/graylog2/weixin  微信接口
/graylog2/txdx  腾讯云短信接口
/graylog2/txdh  腾讯云电话接口
/graylog2/hwdx  华为云短信接口
/graylog2/alydx  阿里云短信接口
/graylog2/alydh  阿里云电话接口
```

 - `graylog3接口`

```
特别说明: graylog3接口针对 graylog版本 >= 3.1.x

/graylog3/phone  腾讯云电话接口(v3.0版本将废弃)
/graylog3/dingding  钉钉接口
/graylog3/weixin  微信接口
/graylog3/txdx  腾讯云短信接口
/graylog3/txdh  腾讯云电话接口
/graylog3/hwdx  华为云短信接口
/graylog3/alydx  阿里云短信接口
/graylog3/alydh  阿里云电话接口
```

 - `语音短信回调接口`

```
/tengxun/status  腾讯语音短信回调接口
```