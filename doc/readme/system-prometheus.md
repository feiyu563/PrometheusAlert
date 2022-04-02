## Prometheus接入配置

--------------------------------------

### 配置方法

通过自定义告警消息模版的方式(使用web页面上的自定义模版)

参考alertmanager配置参考：

```
global:
  resolve_timeout: 5m
route:
  group_by: ['instance']
  group_wait: 10m
  group_interval: 10s
  repeat_interval: 10m
  receiver: 'web.hook.prometheusalert'
receivers:
- name: 'web.hook.prometheusalert'
  webhook_configs:
  - url: 'http://[prometheusalert_url]:8080/prometheusalert?type=dd&tpl=prometheus-dd&ddurl=https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxxxxxxxxxxxxxx&at=18888888888'
```

详细配置也可参考：[推荐 任意告警源（自定义消息模版）接入配置](customtpl.md)

最终告警效果:

![prometheus1](../prometheus.png)