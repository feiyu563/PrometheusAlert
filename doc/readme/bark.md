PrometheusAlert全家桶Bark配置说明

-----------------

```
#---------------------↓Bark接口-----------------------
#是否开启telegram告警通道,可同时开始多个通道0为关闭,1为开启
open-bark=0
#bark默认地址, 建议自行部署bark-server
BARK_URL=https://api.day.app
#bark key, 多个key使用分割
BARK_KEYS=xxxxx
# 复制, 推荐开启
BARK_COPY=1
# 历史记录保存,推荐开启
BARK_ARCHIVE=1
# 消息分组
BARK_GROUP=PrometheusAlert
```