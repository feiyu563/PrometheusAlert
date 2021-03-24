## PrometheusAlert全家桶

![author](https://img.shields.io/badge/author-jikun.zhang-blueviolet.svg)
![license](https://img.shields.io/github/license/feiyu563/PrometheusAlert.svg)
![last commit](https://img.shields.io/github/last-commit/feiyu563/PrometheusAlert.svg)
![issues](https://img.shields.io/github/issues/feiyu563/PrometheusAlert.svg)
![stars](https://img.shields.io/github/stars/feiyu563/PrometheusAlert.svg)
![forks](https://img.shields.io/github/forks/feiyu563/PrometheusAlert.svg)
[![docker](https://img.shields.io/docker/pulls/feiyu563/prometheus-alert.svg)](https://hub.docker.com/r/feiyu563/prometheus-alert)

[![GitHub Release](https://img.shields.io/github/release/feiyu563/PrometheusAlert.svg)](https://github.com/feiyu563/PrometheusAlert/releases)

[![Travis](https://img.shields.io/badge/language-Go-green.svg)]()
[![Travis](https://img.shields.io/badge/language-JavaScript-yellow.svg)]()
[![Travis](https://img.shields.io/badge/language-Shell-blue.svg)]()

-----------------

PrometheusAlert是开源的运维告警中心消息转发系统，支持主流的监控系统Prometheus、Zabbix，日志系统Graylog2，Graylog3、数据可视化系统Grafana、SonarQube，以及所有支持WebHook接口的系统发出的预警消息，支持将收到的这些消息发送到钉钉，微信，email，飞书，腾讯短信，腾讯电话，阿里云短信，阿里云电话，华为短信，百度云短信，容联云电话，七陌短信，七陌语音，TeleGram，百度Hi(如流)等。

![it](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/it.png)

-----------------
* 如果这个项目能为您带来帮助，我将感到非常荣幸！同时非常欢迎您加入PrometheusAlert全家桶的微信群一起探讨关于PrometheusAlert新一期的开发计划。

* PrometheusAlert的成长离不开所有爱好运维和开发的朋友们的支持！如果您也对这个项目感兴趣，请点击一下 Star 以示鼓励，谢谢


<br>
<br>

## 启动

```
#clone项目源代码
git clone https://github.com/feiyu563/PrometheusAlert.git

#进入程序目录并运行PrometheusAlert
cd PrometheusAlert/example/linux/

#运行PrometheusAlert
./PrometheusAlert (#后台运行请执行 nohup ./PrometheusAlert &)

#启动后可使用浏览器打开以下地址查看：http://xxx.xxx.xxx.xxx:8080
```


<br>
<br>

-----------------

![dashboard](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/dashboard.png)
![dashboard-test](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/dashboard-test.png)
![dashboard-tpl-list](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/dashboard-tpl-list.png)

--------------------------------------

PrometheusAlert具备如下特性
---------------------

 - 支持任意包含WebHook接口系统的消息作为消息来源，常见的如Prometheus，Zabbix，Graylog2，Graylog3，Grafana，SonarQube和其他WebHook接口的系统。
 - 支持多种类型的发送目标，支持钉钉，企业微信，email，飞书，腾讯短信，腾讯语音，阿里云短信，阿里云语音，华为短信，百度云短信，容联云电话，七陌短信，七陌语音，TeleGram，百度Hi(如流)。
 - 针对Prometheus增加了告警级别，并且支持按照不同级别发送消息到不同目标对象。
 - 简化Prometheus分组配置，支持按照具体消息发送到单个或多个接收方。
 - 增加手机号码配置项和号码自动轮询配置，可固定发送给单一个人告警信息，也可以通过自动轮询的方式发送到多个人员且支持按照不同日期发送到不同人员。
 - 增加Dashboard，提供配置测试，告警消息模版自定义，告警模版测试等功能
 - 增加自定义告警消息模版支持，可以通过自定义模版对任意服务WebHook进行支持
 - 增加代理支持
 - 增加支持mysql、sqlite3(默认)、postgres作为模版存储，便于集群化部署
 
---------------------

致谢以下伙伴对PrometheusAlert的贡献
---------------------
 - [@minibear2333](https://github.com/minibear2333)
 - [@bigdimple](https://github.com/bigdimple)
 - [@ysicing](https://github.com/ysicing)
 - [@Zhang21](https://github.com/Zhang21)
 - [@jicki](https://github.com/jicki)


--------------------------------------

## PrometheusAlert全家桶新手入门

文档阅读提示


* [【PrometheusAlert安装及配置说明】](doc/readme/install.md)
  * [安装部署PrometheusAlert](doc/readme/install.md)
  * [接口说明](doc/readme/restful.md)
  * [Metrics指标](doc/readme/metrics.md)
  * [配置文件说明](doc/readme/setup.md)
  * [其他配置](doc/readme/other.md)
  * [使用Dashboard对配置进行测试](doc/readme/dashboard.md)
* [【 app.conf 告警接收目标配置】](doc/readme/dingding.md)
  * [钉钉告警配置](doc/readme/dingding.md)
  * [飞书告警配置](doc/readme/feishu.md)
  * [企业微信告警配置](doc/readme/wechat.md)
  * [企业微信应用告警配置](doc/readme/workwechat.md)
  * [阿里云短信和电话告警配置](doc/readme/aliyun.md)
  * [腾讯云短信和电话告警配置](doc/readme/tx.md)
  * [容联云电话告警配置](doc/readme/rly.md)
  * [华为云短信配置](doc/readme/hw.md)
  * [百度云短信配置](doc/readme/bd.md)
  * [Email配置](doc/readme/email.md)
  * [七陌短信和电话告警配置](doc/readme/7moor.md)
  * [TeleGram告警配置](doc/readme/tg.md)
  * [百度Hi(如流)告警配置](doc/readme/ruliu.md)
* [【告警源配置】](doc/readme/prometheus.md)
  * [★★★★★ 推荐 任意告警源（自定义消息模版）接入配置](doc/readme/customtpl.md)
  * [Prometheus 接入配置](doc/readme/prometheus.md)
  * [graylog（内置固定消息模版）接入配置](doc/readme/graylog.md)
  * [grafana（内置固定消息模版）接入配置](doc/readme/grafana.md)
  * [zabbix（内置固定消息模版）接入配置](doc/readme/zabbix.md)
  * [sonarqube（自定义消息模版）接入配置](doc/readme/sonarqube.md)
  * [jenkins（自定义消息模版）接入配置](doc/readme/jenkins.md)

* [【版本更新说明】](https://github.com/feiyu563/PrometheusAlert/releases)
* [【如何对PrometheusAlert进行二次开发】](https://github.com/feiyu563/PrometheusAlert/wiki/%E5%A6%82%E4%BD%95%E8%BF%9B%E8%A1%8C%E3%80%90PrometheusAlert%E3%80%91%E4%BA%8C%E6%AC%A1%E5%BC%80%E5%8F%91)


--------------------------------------

项目源码
----

 - [【GitHub】: PrometheusAlert](https://github.com/feiyu563/PrometheusAlert)
 - [【码云】: PrometheusAlert](https://gitee.com/feiyu563/PrometheusAlert)

【FOR HELP】
----
【Email: 244217140@qq.com】

![me](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/wx.png)
