PrometheusAlert全家桶百度Hi(如流)配置说明

-----------------

 **开启百度Hi(如流)机器人**

在企业群右上角点击【机器人图标】-【添加机器人】-【创建机器人】，添加机器人进群后，可获取机器人的Webhook地址,可参下图：

![ruliu1](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/ruliu1.png)

百度Hi(如流)目前支持Markdown类型消息

```
Markdown支持如下语法子集：

1、标题：1-6级标题，# 号和文本之间要有一个空格

# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题

2、粗体：

**加粗文本**

3、斜体：

*斜体文本*

4、引用：

>引用文本

5、字体颜色：只支持三种内置颜色：green（绿色）、gray（灰色）、red（红色）

<font color="green">绿色</font>
<font color="gray">灰色</font>
<font color="red">红色</font>

6、链接:不支持自定义链接文字,会直接显示原链接

1. [百度一下]（https://www.baidu.com/）
2. https://www.baidu.com/

7、有序列表：1. 和文本之间要加一个空格

1. 有序列表
2. 有序列表
3. 有序列表

8、无序列表：- 和文本之间要加一个空格

- 无序列表
- 无序列表

9、行内代码

`code`

消息发送频率限制
为了保障群成员使用体验，以防收到大量消息的打扰，机器人发消息限制每个机器人每群20条/分钟，超出后接口将返回警告错误码，将限流5分钟。限流期间的消息将会被丢弃。
```

百度Hi(如流)相关配置：

```
#---------------------↓全局配置-----------------------
#告警消息标题
title=PrometheusAlert
#钉钉告警 告警logo图标地址
logourl=https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png
#钉钉告警 恢复logo图标地址
rlogourl=https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png

#---------------------↓百度Hi(如流)-----------------------
#是否开启百度Hi(如流)告警通道,可同时开始多个通道0为关闭,1为开启
open-ruliu=0
#默认百度Hi(如流)机器人地址
BDRL_URL=https://api.im.baidu.com/api/msg/groupmsgsend?access_token=xxxxxxxxxxxxxx
#百度Hi(如流)群ID
BDRL_ID=123456
```
