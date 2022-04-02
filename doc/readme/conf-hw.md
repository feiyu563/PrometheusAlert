## 华为云短信配置

目前华为云仅支持短信告警渠道，开启华为短信告警渠道需要提前注册好华为云平台的帐号。并开通华为云短信相关服务。

开通服务可参考华为云官方文档：

* 短信：https://support.huaweicloud.com/qs-msgsms/sms_02_0001.html

华为云短信模版配置可参考：

* `模版类型:通知类`

* `模版内容:运维告警:${TXT_66}`

华为云短信相关配置：

```
#---------------------↓华为云接口-----------------------
#是否开启华为云短信告警通道,可同时开始多个通道0为关闭,1为开启
open-hwdx=1
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
```