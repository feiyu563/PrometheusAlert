PrometheusAlert 简介

![logo](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/logo.png)
![dashboard](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/dashboard.png)

-----------------

PrometheusAlert是开源的运维告警中心消息转发系统,支持主流的监控系统Prometheus,日志系统Graylog2,Graylog3和数据可视化系统Grafana发出的预警消息,支持将收到的这些消息发送到钉钉,微信,腾讯短信,腾讯电话,阿里云短信,阿里云电话,华为短信等

![it](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/it.png)

--------------------------------------

2.2.1版本更新说明:
---------------------
 - 1.增加配置项dd_isatall,功能是否开启钉钉 @所有人(0为关闭,1为开启)
 - 2.美化前端页面CSS
 - 3.新增接口 /prometheus/router,专门用于处理Prometheus AlertManager的route路由定义的消息(即通过路由指定发送目标),readme中有详细使用方法

--------------------------------------

2.2版本更新说明:
---------------------
 - 1.增加阿里云短信接口支持
 - 2.增加阿里云电话接口支持
 - 3.增加grafana/graylog2/graylog3/prometheus等阿里云电话和短信的发送接口
 - 4.增加dashboard阿里云短信和电话测试告警
 - 5.配置文件新增阿里云短信和电话的配置
 - 6.配置文件更新部分配置名称,主要是短信和电话配置部分名称改变

--------------------------------------

2.1.3版本更新说明:
---------------------
 - 1.修复电话拨打失败回调接口极端情况下会出现死循环的逻辑错误

--------------------------------------

2.1.2版本更新说明:
---------------------
 - 1.修复PrometheusAlert频繁报EOF的BUG

--------------------------------------

2.1.1版本更新说明:
---------------------
 - 1.更新readme文件关于回调接口和user.csv文件的说明
 - 2.更新prometheus对user.csv文件的支持.

--------------------------------------

2.1版本更新说明:
---------------------
 - 1.修复phonecallresolved开关的bug
 - 2.增加告警抑制开关,自动告警抑制是默认同一个告警源的告警信息只发送告警级别最高的第一条告警信息,其他消息默认屏蔽,这么做的目的是为了减少相同告警来源的消息数量,防止告警炸弹,默认开启

--------------------------------------

2.0.2版本更新说明:
---------------------
 - 1.增加配置项故障恢复是否启用电话通知0为关闭,1为开启,phonecallresolved,
 - 2.更改源码包管理工具为godep
 - 3.更新Dockerfile,取消Dockerfile目录

--------------------------------------

2.0版本更新说明:
---------------------
 - 1.prometheus rules中告警级别配置已经由早期的版本中的Annotations字段下迁移到Labels字段下,为方便alertmanager也可以匹配level字段
 - 2.prometheus rules中告警配置summary字段已经删除,现只保留description字段
 - 3.近期新版本的钉钉加了安全设置,只需选择安全设置中的 自定义关键词 即可,并将关键词设置为 Prometheus或者app.conf中设置的title值均可,参考下图
![ding2](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/dingding3.png)
 - 4.配置文件全面更新,支持对各个配置开关,不再限制告警发送的目标,本次配置文件更新较大,建议查看下配置文件中的说明
 - 5.增加新的告警源支持graylog3
 - 6.增加新的目标支持华为云短信
 - 7.告警接口的路由改变较大,主要针对grafana,graylog2,graylog3
 - 8.prometheus rules中告警配置新增字段wxurl,可以指定多个企业微信机器人地址
 - 9.web页面增加华为短信测试按钮
 - 10.grafana,graylog2,graylog3已经开启支持user.csv文件循环告警支持
--------------------------------------

PrometheusAlert具备如下特性
---------------------
 - 支持多种消息来源,目前主要有prometheus,graylog2,graylog3,grafana
 - 支持多种类型的发送目标,支持钉钉,微信,腾讯短信,腾讯语音,华为短信
 - 针对Prometheus增加了告警级别,并且支持按照不同级别发送消息到不同目标对象
 - 简化Prometheus分组配置,支持按照具体消息发送到单个或多个接收方
 - 增加手机号码配置项,和号码自动轮询配置,可固定发送给单一个人告警信息,也可以通过自动轮询的方式发送到多个人员且支持按照不同日期发送到不同人员
 - 增加dashboard,暂时支持测试配置是否正确
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
```
启动后可使用浏览器打开测试地址: http://127.0.0.1:8080
```

配置说明
----
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

--------------------------------------
 **0. 开启钉钉机器人**

打开钉钉,进入钉钉群中,选择群设置-->群机器人-->自定义，可参下图：

![ding](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/dingding1.png)
![ding2](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/dingding2.png)

新版本的钉钉加了安全设置,只需选择安全设置中的 自定义关键词 即可,并将关键词设置为 Prometheus或者app.conf中设置的title值均可,参考下图
![ding2](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/dingding3.png)


--------------------------------------
 **1. 开启企业微信机器人**

打开企业微信,进入企业微信群中,选择群设置-->群机器人-->添加，可参下图：

![wx1](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/wx1.png)
![wx2](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/wx2.png)


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

针对 `/prometheus/router  AlertManager router指定接收端接口`,该接口可在url中直接指定告警的接收端,目前之前三个参数,分别是:wxurl,ddurl,phone(phone用于短信和电话告警)

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
  - url: 'http://[prometheusalert_url]:8080/prometheus/router?wxurl=https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxx&ddurl=https://oapi.dingtalk.com/robot/send?access_token=xxxxx&phone=15395105573'
```

Prometheus Server 的告警rules配置，可参考如下模板：

```
groups:
 1. name: node_alert
  rules:
 2. alert: 主机CPU告警
    expr: node_load1 > 1
    labels:
      name: prometheusalertcenter
      level: 3   #告警级别,告警级别定义 0 信息,1 警告,2 一般严重,3 严重,4 灾难
    annotations:
      description: "{{ $labels.instance }} CPU load占用过高"  #告警信息
      mobile: 15888888881,15888888882,15888888883  #告警发送目标手机号(需要设置电话和短信告警级别)
      ddurl: "https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" #支持添加多个钉钉机器人告警,用,号分割即可,如果留空或者未填写,则默认发送到配置文件中填写的钉钉器人地址
      wxurl: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxxxx-xxxxxx-xxxxxx-xxxxxx,https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxx-xxxx-xxxxxxx-xxxxx" #支持添加多个企业微信机器人告警,用,号分割即可,如果留空或者未填写,则默认发送到配置文件中填写的企业微信机器人地址
```
最终告警效果:

![prometheus1](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/prometheus.png)

--------------------------------------
 **2. Grafana 接入配置**

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

--------------------------------------

 **3. Graylog 接入配置**
打开Graylog管理页面,登录后进入Alerts配置

![graylog1](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/graylog1.png)
点击```Add new notification```创建新的告警通道,选择如下图配置:
![graylog2](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/graylog2.png)
在弹出的窗口中填入名称和对应的PrometheusAlert的接口地址即可,配置参考下图:

 - `graylog2接口`

```
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
/graylog3/phone  腾讯云电话接口(v3.0版本将废弃)
/graylog3/dingding  钉钉接口
/graylog3/weixin  微信接口
/graylog3/txdx  腾讯云短信接口
/graylog3/txdh  腾讯云电话接口
/graylog3/hwdx  华为云短信接口
/graylog3/alydx  阿里云短信接口
/graylog3/alydh  阿里云电话接口
```

![graylog3](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/graylog3.png)
配置完成后,点击```Test```测试下是否能够正常接收告警消息即可

最终告警效果:

![graylog4](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/graylog4.png)

--------------------------------------

**4. 配置文件解析**
--------------------------------------

1.腾讯云短信告警和语音告警

具体参数获取可去腾讯云开通相关服务即可,并配置相关参数https://cloud.tencent.com/document/product/382/18071
腾讯云模版配置可参考

`prometheus告警:{1}`

![tengxun1](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/tengxun1.png)

--------------------------------------

2.华为云短信告警

具体参数获取可去华为云开通相关服务即可,并配置相关参数
https://support.huaweicloud.com/qs-msgsms/sms_02_0001.html
华为云短信模版配置可参考

`模版类型:通知类`

`模版内容:运维告警:${TXT_66}`

--------------------------------------

3.阿里云短信告警和语音告警

具体参数获取可去阿里云开通相关服务即可,并配置相关参数https://help.aliyun.com/document_detail/55288.html?spm=5176.8195934.1283918.11.73af30c917Jfti&aly_as=m1ouSihk

阿里云模版配置可参考,code为变量名称,不推荐更改

`prometheus告警:${code}`

--------------------------------------

```
#---------------------↓全局配置-----------------------
appname = PrometheusAlert
#监听端口
httpport = 8080
runmode = dev
#开启JSON请求
copyrequestbody = true
#告警消息标题
title=NX-TEST
#链接到告警平台地址
GraylogAlerturl=http://graylog.org
#logo图标地址
logourl=https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png
#短信告警级别(等于3就进行短信告警) 告警级别定义 0 信息,1 警告,2 一般严重,3 严重,4 灾难
messagelevel=3
#电话告警级别(等于4就进行语音告警) 告警级别定义 0 信息,1 警告,2 一般严重,3 严重,4 灾难
phonecalllevel=4
#默认拨打号码
defaultphone=15395105573
#故障恢复是否启用电话通知0为关闭,1为开启
phonecallresolved=1
#自动告警抑制(自动告警抑制是默认同一个告警源的告警信息只发送告警级别最高的第一条告警信息,其他消息默认屏蔽,这么做的目的是为了减少相同告警来源的消息数量,防止告警炸弹,0为关闭,1为开启)
silent=1

#---------------------↓webhook-----------------------
#是否开启钉钉告警通道,可同时开始多个通道0为关闭,1为开启
open-dingding=1
#默认钉钉机器人地址
ddurl=https://oapi.dingtalk.com/robot/send?access_token=xxxxx

#是否开启微信告警通道,可同时开始多个通道0为关闭,1为开启
open-weixin=1
#默认企业微信机器人地址
wxurl=https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxx

#---------------------↓腾讯云接口-----------------------
#是否开启腾讯云短信告警通道,可同时开始多个通道0为关闭,1为开启
open-txdx=0
#腾讯云短信接口key
TXY_DX_appkey=xxxxx
#腾讯云短信模版ID 腾讯云短信模版配置可参考 prometheus告警:{1}
TXY_DX_tpl_id=xxxxx
#腾讯云短信sdk app id
TXY_DX_sdkappid=xxxxx
#腾讯云短信签名 根据自己审核通过的签名来填写
TXY_DX_sign=腾讯云

#是否开启腾讯云电话告警通道,可同时开始多个通道0为关闭,1为开启
TXY_DH_open-txdh=0
#腾讯云电话接口key
TXY_DH_phonecallappkey=xxxxx
#腾讯云电话模版ID
TXY_DH_phonecalltpl_id=xxxxx
#腾讯云电话sdk app id
TXY_DH_phonecallsdkappid=xxxxx

#---------------------↓华为云接口-----------------------
#是否开启华为云短信告警通道,可同时开始多个通道0为关闭,1为开启
open-hwdx=0
#华为云短信接口key
HWY_DX_APP_Key=xxxxxxxxxxxxxxxxxxxxxx
#华为云短信接口Secret
HWY_DX_APP_Secret=xxxxxxxxxxxxxxxxxxxxxx
#华为云APP接入地址(端口接口地址)
HWY_DX_APP_Url=https://rtcsms.cn-north-1.myhuaweicloud.com:10743
#华为云短信模板ID
HWY_DX_Templateid=xxxxxxxxxxxxxxxxxxxxxx
#华为云签名名称，必须是已审核通过的，与模板类型一致的签名名称,按照自己的实际签名填写
HWY_DX_Signature=华为云
#华为云签名通道号
HWY_DX_Sender=xxxxxxxxxx

#---------------------↓阿里云接口-----------------------
#是否开启阿里云短信告警通道,可同时开始多个通道0为关闭,1为开启
open-alydx=0
#阿里云短信主账号AccessKey的ID
ALY_DX_AccessKeyId=xxxxxxxxxxxxxxxxxxxxxx
#阿里云短信接口密钥
ALY_DX_AccessSecret=xxxxxxxxxxxxxxxxxxxxxx
#阿里云短信签名名称
ALY_DX_SignName=阿里云
#阿里云短信模板ID
ALY_DX_Template=xxxxxxxxxxxxxxxxxxxxxx

#是否开启阿里云电话告警通道,可同时开始多个通道0为关闭,1为开启
open-alydx=0
#阿里云电话主账号AccessKey的ID
ALY_DH_AccessKeyId=xxxxxxxxxxxxxxxxxxxxxx
#阿里云电话接口密钥
ALY_DH_AccessSecret=xxxxxxxxxxxxxxxxxxxxxx
#阿里云电话被叫显号，必须是已购买的号码
ALY_DX_CalledShowNumber=xxxxxxxxx
#阿里云电话文本转语音（TTS）模板ID
ALY_DH_TtsCode=xxxxxxxx

```

另外 PrometheusAlert 同时支持按照日期发送告警到不同号码,并且已经加入告警失败或者被告警人未接听电话后转联系默认联系人(defaultphone),只需新建user.csv文件,并将文件放到程序运行目录下即可自动加载,同时该文件也是电话回调接口所必需文件(如回调接口返回非0状态,则会在此文件中寻找下一号码进行拨打,如开启回调功能,请务必创建该文件)
ps:目前grafana/graylog的电话和短信告警依赖于该文件,prometheus电话和短信告警优先从rules的Mobile字段读取,如未配置号码,则会从user.csv中读取,如user.csv中也未配置,则会直接发送到defaultphone
关于user.csv中值班时间切换问题,默认每日上午10点进行切换
csv文件格式如下:
```
2019年4月10日,15888888881,小张,15999999999,备用联系人小陈,15999999998,备用联系人小赵
2019年4月11日,15888888882,小李,15999999999,备用联系人小陈,15999999998,备用联系人小赵
2019年4月12日,15888888883,小王,15999999999,备用联系人小陈,15999999998,备用联系人小赵
2019年4月13日,15888888884,小宋,15999999999,备用联系人小陈,15999999998,备用联系人小赵
```

启用电话告警失败重试功能需要在腾讯云上配置:
云产品--->短信--->事件回调配置--->语音短信回调(回调接口 http://prometheusalertcenter_url/tengxun/status)
ps:回调接口需对公网开放,否则云平台无法访问到接口.开启回调之后请务必创建user.csv文件

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
