腾讯云短信告警和语音告警

具体参数获取可去腾讯云开通相关服务即可,并配置相关参数：
短信：https://cloud.tencent.com/document/product/382/18071
语音（电话）：https://cloud.tencent.com/document/product/1128/37343
腾讯云模版配置可参考

`prometheus告警:{1}`

![tengxun1](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/tengxun1.png)

启用电话告警失败重试功能需要在腾讯云上配置:
云产品--->短信--->事件回调配置--->语音短信回调(回调接口 http://prometheusalertcenter_url/tengxun/status)
ps:回调接口需对公网开放,否则云平台无法访问到接口.开启回调之后请务必创建user.csv文件

![tengxun2](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/tengxun2.png)