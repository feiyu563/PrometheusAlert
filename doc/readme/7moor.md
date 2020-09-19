# 7moor短信和语音通知配置

地址:

- 官网: www.7moor.com
- 文档: developer.7moor.com
- 用户后台: kf.7moor.com

<br>

使用7moor短信或webcall语音通知需要以下配置:

- 七陌账户ID
- 七陌账户APISecret
- 短信模板编号
- webcall虚拟服务号
- webcall语音通知，文本节点里语音信息替换的变量

<br>

**注意:**

- 如果不会配置，请联系七陌客服咨询。
- 七陌短信，默认配置一个变量(`var1`)来代表告警内容。
- 七陌webcall语音通知，文本节点语音信息替换的变量，我默认配置为`text`。
- 七陌webcall语音通知内容里，不支持英文逗号(`,`)句号(`.`)等符号(会响应408错误)，但支持中文符号。请注意！
- 七陌webcall语音通知，语速、重复次数、语调、音声等请自己在用户后台配置。


<br>
<br>


## 相关配置

```
#---------------------↓七陌云接口-----------------------
#是否开启七陌短信告警通道,可同时开始多个通道0为关闭,1为开启
open-7moordx=0
#七陌账户ID
7MOOR_ACCOUNT_ID=Nxxx
#七陌账户APISecret
7MOOR_ACCOUNT_APISECRET=xxx
#七陌账户短信模板编号
7MOOR_DX_TEMPLATENUM=n
#注意：七陌短信变量这里只用一个var1，在代码里写死了。
#-----------
#是否开启七陌webcall语音通知告警通道,可同时开始多个通道0为关闭,1为开启
open-7moordh=0
#请在七陌平台添加虚拟服务号、文本节点
#七陌账户webcall的虚拟服务号
7MOOR_WEBCALL_SERVICENO=xxx
# 文本节点里被替换的变量，我配置的是text。如果被替换的变量不是text，请修改此配置
7MOOR_WEBCALL_VOICE_VAR=text
```

