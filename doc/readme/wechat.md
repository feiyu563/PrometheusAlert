## 企业微信告警配置

打开企业微信,进入企业微信群中,选择群设置-->群机器人-->添加，可参下图：

![wx1](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/wx1.png)

![wx2](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/wx2.png)

复制图中的Webhook地址，并填入PrometheusAlert配置文件app.conf中对应配置项即可。

企业微信机器人目前支持的markdown语法是如下的子集：

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

企业微信机器人相关配置：

```
#---------------------↓webhook-----------------------
#是否开启微信告警通道,可同时开始多个通道0为关闭,1为开启
open-weixin=1
#默认企业微信机器人地址
wxurl=https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxx
```