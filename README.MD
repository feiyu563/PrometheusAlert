该插件基于beego web框架开发，如需编译或运行请先下载beego，安装beego请移步 https://beego.me/quickstart

关于部署，我已经将编译好的文件放到了build目录下，可以直接放到Linux下的任意目录，给予执行权限运行即可。

使用方法很简单，只需要修改下prometheus alert插件配置，可参考如下模板：
--------------------------
```
global:
  resolve_timeout: 5m

route:
  group_by: ['alertname']
  group_wait: 10m
  group_interval: 10s
  repeat_interval: 10m
  receiver: 'web.hook.dingding'
receivers:
- name: 'web.hook.dingding'
  webhook_configs:
  - url: 'http://192.168.50.73:8080/prometheus/alert'
  ```
--------------------------
webhook地址为： http://[程序部署的IP:端口]/prometheus/alert 即可.


关于告警信息：
配置Prometheus server的告警rules即可，可参考如下模板：
--------------------------------------  
```
groups:
- name: node_alert
  rules:
  - alert: Node_CPU_Alert
    expr: node_load1 > 3
    labels:
      node: alert
    annotations:
      description: If this value is high, it shows that CPU is very complex and very high !  #描述告警信息
      summary: "{{ $labels.instance }} node CPU load goes too high !(current value: {{ $value }}s)"  #告警的实际内容
```
--------------------------------------      

显示到钉钉的告警信息如下：
--------------------------------------
![Dashboard UI workloads page](doc/alert.png)

内容还不算很完善，有什么好的建议可以直接与我联系，我抽空多丰富些功能

现在已经支持graylog的alert
router 是 graylog/alert
--------------------------------------
QQ:244217140
