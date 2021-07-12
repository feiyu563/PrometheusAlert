PrometheusAlert全家桶企业微信应用配置说明

-----------------

企业微信应用目前支持的markdown语法是如下的子集：

```
标题 （支持1至6级标题，注意#与文字中间要有空格）
# 标题一
## 标题二
### 标题三
#### 标题四
##### 标题五
###### 标题六

加粗
**bold**

链接
[这是一个链接](http://work.weixin.qq.com/api/doc)

行内代码段（暂不支持跨行）
`code`

引用
> 引用文字

字体颜色(只支持3种内置颜色)
<font color="info">绿色</font>
<font color="comment">灰色</font>
<font color="warning">橙红色</font>
```

企业微信应用相关配置：

```
#---------------------↓workwechat接口-----------------------
#是否开启workwechat告警通道,可同时开始多个通道0为关闭,1为开启
open-workwechat=0
# 企业ID
WorkWechat_CropID=xxxxx
# 应用ID
WorkWechat_AgentID=xxxx
# 应用secret
WorkWechat_AgentSecret=xxxx
# 接受用户
WorkWechat_ToUser="zhangsan|lisi"
# 接受部门
WorkWechat_ToParty="ops|dev"
# 接受标签
WorkWechat_ToTag=""
# 消息类型, 暂时只支持markdown
# WorkWechat_Msgtype = "markdown"
```

特别提醒，目前仅自定义模板接口（`/prometheusalert`）支持动态定义 `接受用户`,`接受部门`,`接受标签`.其他接口均默认使用配置文件中的固定配置。

![workwechat1](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/workwechat1.png)