## 腾讯云短信和电话告警配置

开启腾讯云短信和语音告警渠道均需要提前注册好腾讯云平台的帐号。并开通腾讯云短信和语音相关服务。

开通服务可参考腾讯云官方文档：

* 短信：https://cloud.tencent.com/document/product/382/18071

* 语音（电话）：https://cloud.tencent.com/document/product/1128/37343

* 注意事项：开通腾讯云短信和语音需要配置的模版请填写类似如下内容：`prometheus告警:{1}`

![tengxun1](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/tengxun1.png)

启用电话告警失败重试功能需要在腾讯云上配置:

云产品--->短信--->事件回调配置--->语音短信回调(回调接口 http://prometheusalertcenter_url/tengxun/status)

ps:回调接口需对公网开放,否则云平台无法访问到接口.开启回调之后请务必创建user.csv文件

![tengxun2](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/tengxun2.png)

腾讯云短信和语音相关配置：

```
#---------------------↓腾讯云接口-----------------------
#是否开启腾讯云短信告警通道,可同时开始多个通道0为关闭,1为开启
open-txdx=1
#腾讯云短信接口key
TXY_DX_appkey=xxxxx
#腾讯云短信模版ID 腾讯云短信模版配置可参考 prometheus告警:{1}
TXY_DX_tpl_id=xxxxx
#腾讯云短信sdk app id
TXY_DX_sdkappid=xxxxx
#腾讯云短信签名 根据自己审核通过的签名来填写
TXY_DX_sign=腾讯云

#是否开启腾讯云电话告警通道,可同时开始多个通道0为关闭,1为开启
TXY_DH_open-txdh=1
#腾讯云电话接口key
TXY_DH_phonecallappkey=xxxxx
#腾讯云电话模版ID
TXY_DH_phonecalltpl_id=xxxxx
#腾讯云电话sdk app id
TXY_DH_phonecallsdkappid=xxxxx
```