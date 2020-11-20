## 百度云短信配置

目前百度云仅支持短信告警渠道，开启百度云短信告警渠道需要提前注册好百度云平台的帐号。并开通百度云短信相关服务。

开通服务可参考百度云官方文档：

* 短信：https://cloud.baidu.com/doc/SMS/index.html

百度云短信模版配置可参考：

* `模版内容:运维告警:{code}`

百度云短信相关配置：

```
#---------------------↓百度云接口-----------------------
#是否开启百度云短信告警通道,可同时开始多个通道0为关闭,1为开启
open-baidudx=0
#百度云短信接口AK(ACCESS_KEY_ID)
BDY_DX_AK=xxxxx
#百度云短信接口SK(SECRET_ACCESS_KEY)
BDY_DX_SK=xxxxx
#百度云短信ENDPOINT（ENDPOINT参数需要用指定区域的域名来进行定义，如服务所在区域为北京，则为）
BDY_DX_ENDPOINT=http://smsv3.bj.baidubce.com
#百度云短信模版ID,根据自己审核通过的模版来填写(模版支持一个参数code：如prometheus告警:{code})
BDY_DX_TEMPLATE_ID=xxxxx
#百度云短信签名ID，根据自己审核通过的签名来填写
TXY_DX_SIGNATURE_ID=xxxxx
```