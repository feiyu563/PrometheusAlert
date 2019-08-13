PrometheusAlert 简介

![logo](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/logo.png)

-----------------

PrometheusAlert是开源的运维告警中心消息转发系统,支持主流的监控系统Prometheus,日志系统Graylog和数据可视化系统Grafana发出的预警消息,支持将收到的这些消息发送到钉钉,短信和语音提醒等

![it](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/it.png)
--------------------------------------

PrometheusAlert具备如下特性
---------------------
 - 支持多种消息来源,目前主要有prometheus,graylog,grafana
 - 支持多种类型的发送目标,支持钉钉,短信,语音
 - 针对Prometheus增加了告警级别,并且支持按照不同级别发送消息到不同目标对象
 - 简化Prometheus分组配置,支持按照具体消息发送到单个或多个接收方
 - 增加手机号码配置项,和号码自动轮询配置,可固定发送给单一个人告警信息,也可以通过自动轮询的方式发送到多个人员且支持按照不同日期发送到不同人员

--------------------------------------
部署方式
----

PrometheusAlert可以部署在本地和云平台上，支持windows、linux、公有云、私有云、混合云、容器和kubernetes。你可以根据实际场景或需求，选择相应的方式来部署PrometheusAlert：

 - 使用容器部署
```
git clone https://github.com/feiyu563/PrometheusAlert.git
mkdir /etc/prometheusalert-center/
cp PrometheusAlert/conf/app.conf /etc/prometheusalert-center/
docker run -d -p 8080:8080 -v /etc/prometheusalert-center:/app/conf --name prometheusalert-center feiyu563/prometheus-alert:latest
```
 - 在linux系统中部署
```
git clone https://github.com/feiyu563/PrometheusAlert.git
cd PrometheusAlert/example/linux/
./PrometheusAlert #后台运行请执行nohup ./PrometheusAlert &
```
- 在windows系统中运行
```
git clone https://github.com/feiyu563/PrometheusAlert.git
cd PrometheusAlert/example/windows/
双击运行 PrometheusAlert.exe即可
```
- 在kubernetes中运行
```
kubectl app -n monitoring -f https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/example/kubernetes/PrometheusAlert-Deployment.yaml
```
- 使用helm部署
```
git clone https://github.com/feiyu563/PrometheusAlert.git
cd PrometheusAlert/example/helm/prometheusalert
#如需修改配置文件,请更新config中的app.conf
helm install -n monitoring .
```
配置说明
----
--------------------------------------

PrometheusAlert 暂提供以下三个接口,分别对应各自接入端

 - `prometheus接口`  `/prometheus/alert`
 - `grafana钉钉接口`     `/grafana/dingding`
 - `grafana电话接口`     `/grafana/phone`
 - `graylog接口`     `/graylog/alert`
 - `腾讯云语音短信回调接口`     `/tengxun/status`
 - `阿里云短信/语音提醒`  `即将接入`
 
--------------------------------------
 **0. 开启钉钉机器人**

打开钉钉,进入钉钉群中,选择群设置-->群机器人-->自定义，可参下图：

![ding](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/dingding1.png)
![ding2](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/dingding2.png)

--------------------------------------

 **1. Prometheus 接入配置**

在 Prometheus Alertmanager 中启用 Webhook，可参考如下模板：

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
  - url: 'http://[prometheusalert_url]:8080/prometheus/alert'
```

Prometheus Server 的告警rules配置，可参考如下模板：

```
groups:
 1. name: node_alert
  rules:
 2. alert: 主机CPU告警
    expr: node_load1 > 3
    labels:
      severity: warning
    annotations:
      description: "{{ $labels.instance }} CPU load占用过高"  #告警信息
      summary: "{{ $labels.instance }} CPU load占用过高已经恢复"  #告警恢复信息
      level: 3   #告警级别,告警级别定义 0 信息,1 警告,2 一般严重,3 严重,4 灾难
      mobile: 15888888881,15888888882,15888888883  #告警发送目标手机号(需要设置电话和短信告警级别)
      ddurl: "https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" #支持添加多个钉钉告警,用,号分割即可,如果留空或者未填写,则默认发送到配置文件中填写的钉钉地址
```
最终告警效果:

![prometheus1](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/prometheus.png)

--------------------------------------
 **2. Grafana 接入配置**
 
打开grafana管理页面,登录后进入notification channels配置

![grafana1](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/addchannel.png)
注意这里的url地址填写上自己部署所在的url
![grafana2](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/addchannel2.png)
配置完成后保存即可.继续进行告警消息配置,选择任意一个折线图,点击编辑,进入aler配置,配置参考下图:
![grafana3](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/grafanaalert1.png)
![grafana4](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/grafanaalert2.png)

Notifications配置格式参考,支持配置多个钉钉机器人url:
```
告警消息内容&&ddurl[钉钉机器人url,钉钉机器人url....]
```

最终告警效果:

![grafana5](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/grafana.png)

--------------------------------------

 **3. Graylog 接入配置**
打开Graylog管理页面,登录后进入Alerts配置

![graylog1](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/graylog1.png)
点击```Add new notification```创建新的告警通道,选择如下图配置:
![graylog2](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/graylog2.png)
在弹出的窗口中填入名称和对应的PrometheusAlert的地址即可,配置参考下图:
![graylog3](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/graylog3.png)
配置完成后,点击```Test```测试下是否能够正常接收告警消息即可

最终告警效果:

![graylog4](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/graylog4.png)

--------------------------------------

**4. 配置文件解析**

短信告警和语音告警均使用的是腾讯云的短信和语音提醒接口,具体参数获取可去腾讯云开通相关服务即可
![tengxun1](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/tengxun1.png)


```
appname = PrometheusAlert
#监听端口
httpport = 8080
runmode = dev
#开启JSON请求
copyrequestbody = true
#钉钉机器人地址
ddurl=https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
#告警消息标题
title=NB云平台
#点击告警消息后链接到告警平台地址
alerturl=http://prometheus.local
#告警消息中显示的logo图标地址
logourl=https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/logo.png
#腾讯短信接口key
appkey=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
#腾讯短信模版ID
tpl_id=xxxxxxxx
#腾讯短信sdk app id
sdkappid=xxxxxxxxxx
#短信告警级别(等于3就进行短信告警) 告警级别定义 0 信息,1 警告,2 一般严重,3 严重,4 灾难
messagelevel=3
#腾讯电话接口key
phonecallappkey=xxxxxxxxxxxxxxxxxx
#腾讯电话模版ID
phonecalltpl_id=xxxxxxxx
#腾讯电话sdk app id
phonecallsdkappid=xxxxxxxx
#电话告警级别(等于4就进行语音告警) 告警级别定义 0 信息,1 警告,2 一般严重,3 严重,4 灾难
phonecalllevel=4
#默认拨打号码,默认不配置,如果配置了此项,那么按照user.csv文件轮询的方式将自动失效
#defaultphone=
#未使用user.csv或者user.csv中的号码轮询结束备用号码
backupphone=xxxxxxxxxxx
```

另外 PrometheusAlert 同时支持按照日期发送告警到不同号码,并且已经加入告警失败或者被告警人未接听电话后转联系备用联系人,只需新建user.csv文件,并将文件放到程序运行目录下即可自动加载,该功能暂时支持grafana发出的告警信息,csv文件格式如下:
```
2019年4月10日,15888888881,小张,15999999999,备用联系人小陈,15999999998,备用联系人小赵
2019年4月11日,15888888882,小李,15999999999,备用联系人小陈,15999999998,备用联系人小赵
2019年4月12日,15888888883,小王,15999999999,备用联系人小陈,15999999998,备用联系人小赵
2019年4月13日,15888888884,小宋,15999999999,备用联系人小陈,15999999998,备用联系人小赵
```

启用告警失败重试功能需要在腾讯云上配置: 
云产品--->短信--->事件回调配置--->语音短信回调

![tengxun2](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/tengxun2.png)


--------------------------------------


项目源码
----

 - [PrometheusAlert][1]


  [1]: https://github.com/feiyu563/PrometheusAlert

FOR HELP
----
Email: 244217140@qq.com

![me](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/wx.png)

