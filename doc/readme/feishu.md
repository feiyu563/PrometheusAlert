PrometheusAlert全家桶飞书配置说明

-----------------

飞书机器人目前支持V1和V2两个版本，区别在于地址的不同

`V1地址：https://open.feishu.cn/open-apis/bot/hook​/​xxxxxxxxxxxxxxxxxx`
`V2地址：https://open.feishu.cn/open-apis/bot/v2/hook/xxxxxxxxxxxxxxxxx`

 **开启飞书机器人v1**

打开飞书,进入飞书群中,选择群设置-->群机器人-->添加机器人-->Custom Bot，可参下图：

![fei](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/feishu1.png)

![fei2](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/feishu2.png)

![fei3](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/feishu3.png)

复制图中的Webhook地址，并填入PrometheusAlert配置文件app.conf中对应配置项即可。


 **开启飞书机器人v2**

进入你的目标群组，打开会话设置，找到群机器人，并点击添加机器人。选择通知机器人，添加Custom Bot（自定义机器人）加入群聊。
![fei4](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/feishu4.png)
为你的机器人输入一个合适的名字和描述，并选择添加。同时，你会获取该群组的 webhook 地址，格式如下。请妥善保存好此地址，避免泄露，恶意发送信息。
![fei5](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/feishu5.gif)
飞书V2接口的告警消息采用`消息卡片`消息模式，使用`lark_md`消息格式。
消息格式参考：
https://open.feishu.cn/document/ukTMukTMukTM/uADOwUjLwgDM14CM4ATN

飞书相关配置：

```
#---------------------↓webhook-----------------------
#是否开启飞书告警通道,可同时开始多个通道0为关闭,1为开启
open-feishu=1
#默认飞书机器人地址
fsurl=https://open.feishu.cn/open-apis/bot/v2/hook/xxxxxxxxxxxxxxxxx
```
