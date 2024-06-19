# PrometheusAlert全家桶Kafka配置说明

-----------------

PrometheusAlert支持将收到的json消息通过模板渲染后转发给Kafka集群，使用前需要在配置文件app.conf中配置好Kafka服务的连接信息

```
#---------------------↓kafka地址-----------------------
# kafka服务器的地址
open-kafka=1
kafka_server = 127.0.0.1:9092
# 写入消息的kafka topic
kafka_topic = devops
# 用户标记该消息是来自PrometheusAlert,一般无需修改
kafka_key = PrometheusAlert
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
  - url: 'http://[prometheusalert_url]:8080/prometheusalert?type=kafka&tpl=prometheus-kafka'
```