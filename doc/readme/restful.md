## PrometheusAlert接口说明

--------------------------------------

PrometheusAlert 目前提供以下几类接口，分别对应各自接入端，负责解析各自接口传入或者传出的消息。

- `prometheusalert自定义模版接口`

```
/prometheusalert?type=${type}&tpl=${template}[&rr=true][&ddurl=${ddur}][&wxurl=${wxurl}][&fsurl=${fsurl}][&phone=${phonenumber}]   自定义模版接口，可通过Dashboard自定义模版后，支持任意WebHook接入
```

- `metrics接口`

```
/metrics           展示PrometheusAlert指标信息
```

- `zabbix接口`

```
/zabbix/alert  处理Zabbix告警消息转发默认接口
```

- `语音短信回调接口`

```
/tengxun/status     处理腾讯云语音短信回调接口，负责失败后重试
```
