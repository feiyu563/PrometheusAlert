PrometheusAlert全家桶Telegram配置说明

-----------------

TG机器人TOKEN获取方式: 登录Telegram联系`@BotFather`创建机器人

```
#---------------------↓telegram接口-----------------------
#是否开启telegram告警通道,可同时开始多个通道0为关闭,1为开启
open-tg=0
#tg机器人token
TG_TOKEN=xxxxx
#tg消息模式 个人消息或者频道消息 0为关闭(推送给个人)，1为开启(推送给频道)
TG_MODE_CHAN=0
#tg用户ID 数字
TG_USERID=xxxxx
#tg频道name
TG_CHANNAME=xxxxx
#tg api地址, 可以配置为代理地址
#TG_API_PROXY="https://api.telegram.org/bot%s/%s"
```