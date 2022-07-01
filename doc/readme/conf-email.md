# PrometheusAlert全家桶Email配置说明

-----------------

Email相关配置：

```
#---------------------↓邮件配置-----------------------
#是否开启邮件
open-email=1
#邮件发件服务器地址
Email_host=smtp.qq.com
#邮件发件服务器端口
Email_port=465
#邮件帐号
Email_user=123456789@qq.com
#邮件密码
Email_password=xxxxxxx
#邮件标题
Email_title=运维告警
#默认发送邮箱
Default_emails=123456@qq.com,123456@baidu.com
```


**如何使用**

以Prometheus配合自定义模板为例：

Prometheus配置参考：

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
  - url: 'http://[prometheusalert_url]:8080/prometheusalert?type=email&tpl=prometheus-email&email=Email地址,Email地址2'
```