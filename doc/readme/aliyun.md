## 阿里云短信和电话告警配置

开启阿里云短信告警和语音告警需要提前注册好阿里云平台的帐号。并开通阿里云短信和语音相关服务。

开通服务可参考阿里云官方文档：

* 短信：https://help.aliyun.com/document_detail/55288.html?spm=5176.8195934.1283918.11.73af30c917Jfti&aly_as=m1ouSihk

* 语音（电话）：https://help.aliyun.com/document_detail/55070.html?spm=a2c4g.11186623.6.547.77e271b9SAbG6p

* 注意事项：开通阿里云短信和语音需要配置的模版请填写类似如下内容，`prometheus告警:${code}`，其中`${code}`为固定内容，请勿替换。


配置阿里云短信和语音参数：

```
#---------------------↓阿里云接口-----------------------
#是否开启阿里云短信告警通道,可同时开始多个通道0为关闭,1为开启
open-alydx=1
#阿里云短信主账号AccessKey的ID
ALY_DX_AccessKeyId=xxxxxxxxxxxxxxxxxxxxxx
#阿里云短信接口密钥
ALY_DX_AccessSecret=xxxxxxxxxxxxxxxxxxxxxx
#阿里云短信签名名称
ALY_DX_SignName=阿里云
#阿里云短信模板ID
ALY_DX_Template=xxxxxxxxxxxxxxxxxxxxxx

#是否开启阿里云电话告警通道,可同时开始多个通道0为关闭,1为开启
open-alydx=1
#阿里云电话主账号AccessKey的ID
ALY_DH_AccessKeyId=xxxxxxxxxxxxxxxxxxxxxx
#阿里云电话接口密钥
ALY_DH_AccessSecret=xxxxxxxxxxxxxxxxxxxxxx
#阿里云电话被叫显号，必须是已购买的号码（为空则为阿里云公共号码池电话号码）
ALY_DX_CalledShowNumber=xxxxxxxxx
#阿里云电话文本转语音（TTS）模板ID
ALY_DH_TtsCode=xxxxxxxx
```
开通阿里云语音服务，并进行资质管理认证，添加语音通知模板（建议模板内容使用中文，阿里云语音服务朗读会比较清晰）

![aliyun01](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/aliyun-01.png)

在AlertTemplate中创建阿里云电话通知模板，内容可参照如图（appname为我自定义标签，可修改为自己的定义的标签）

![aliyun02](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/aliyun-02.png)

在prometheus中添加rules，如下（description为语音朗读内容）

```
  - alert: MemoryUsageAlert
    expr: (node_memory_MemTotal_bytes - (node_memory_MemFree_bytes + node_memory_Buffers_bytes + node_memory_Cached_bytes)) / node_memory_MemTotal_bytes * 100 > 90
    for: 1m
    labels:
      type: memory
      level: 4
    annotations:
      summary: "Instance {{ $labels.instance }} Memory usgae high"
      description: "{{ $labels.appname }} 内存使用率超过90% (current value: {{ $value }})"

```

在alertmanager中配置路由规则，采用的为prometheusalert接口，模板为采用上面创建的ali-phone模板

```
  - receiver: 'prometheusalert-phone-db'
    group_wait: 10s
    match:
      level: '4'
      type: db
receivers:
- name: 'prometheusalert-phone-db'
  webhook_configs:
  - url: 'http://localhost:8080/prometheusalert?type=alydh&tpl=ali-phone&phone=136xxxxxxx'
    send_resolved: false
```
注意：
1、阿里云语音服务现在已不支持专有号码外呼
2、阿里云语音服务对IP信息会Block，语音播报避免播报IP信息
3、阿里云语音服务存在流控，https://help.aliyun.com/document_detail/149826.html?spm=a2c4g.11186623.6.683.6cf54c07wo2zTO

