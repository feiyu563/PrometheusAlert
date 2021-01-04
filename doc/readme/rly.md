## 容联云电话告警配置

登录地址：https://www.yuntongxun.com/member/main

如需开通请联系云通讯商务人员或拨打400-610-1019

* 语音（电话）：https://doc.yuntongxun.com/p/5a5342c73b8496dd00dce139

* 注意事项：开通容联云需要配置的模版请填写类似如下内容：`prometheus告警:{9}`

所需配置信息：

![rly](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/ronglianyun.png)

容联云语音相关配置：

```
#---------------------↓容联云接口-----------------------
#是否开启容联云电话告警通道,可同时开始多个通道0为关闭,1为开启
RLY_DH_open-rlydh=1
#容联云基础接口地址
RLY_URL=https://app.cloopen.com:8883/2013-12-26/Accounts/
#容联云后台SID
RLY_ACCOUNT_SID=xxxxxxxxxxx
#容联云api-token
RLY_ACCOUNT_TOKEN=xxxxxxxxxx
#容联云app_id
RLY_APP_ID=xxxxxxxxxxxxx
```