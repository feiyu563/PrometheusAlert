PrometheusAlert全家桶飞书配置说明

-----------------

 **开启飞书机器人**

打开飞书,进入飞书群中,选择群设置-->群机器人-->添加机器人-->Custom Bot，可参下图：

![fei](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/feishu1.png)

![fei2](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/feishu2.png)

![fei3](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/feishu3.png)

复制图中的Webhook地址，并填入PrometheusAlert配置文件app.conf中对应配置项即可。

飞书相关配置：

```
#---------------------↓webhook-----------------------
#是否开启飞书告警通道,可同时开始多个通道0为关闭,1为开启
open-feishu=1
#默认飞书机器人地址
fsurl=https://open.feishu.cn/open-apis/bot/hook/xxxxxxxxx
```