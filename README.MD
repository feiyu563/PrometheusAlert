## PrometheusAlert全家桶

[![Travis](https://img.shields.io/badge/language-Go-green.svg)]()
[![Travis](https://img.shields.io/badge/language-JavaScript-yellow.svg)]()
[![Travis](https://img.shields.io/badge/language-Shell-blue.svg)]()

![历史共访问次数](https://visitor-count-badge.herokuapp.com/total.svg?repo_id=PrometheusAlert.PrometheusAlert)
![今天被访问次数](https://visitor-count-badge.herokuapp.com/today.svg?repo_id=PrometheusAlert.PrometheusAlert)

-----------------

![dashboard](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/dashboard.png)

-----------------

PrometheusAlert是开源的运维告警中心消息转发系统，支持主流的监控系统Prometheus、Zabbix，日志系统Graylog2，Graylog3和数据可视化系统Grafana发出的预警消息，支持将收到的这些消息发送到钉钉，微信，腾讯短信，腾讯电话，阿里云短信，阿里云电话，华为短信等。

![it](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/it.png)

--------------------------------------

PrometheusAlert具备如下特性
---------------------
 - 支持多种消息来源，目前主要有Prometheus，Zabbix，Graylog2，Graylog3，Grafana。
 - 支持多种类型的发送目标，支持钉钉，微信，腾讯短信，腾讯语音，阿里云短信，阿里云语音，华为短信。
 - 针对Prometheus增加了告警级别，并且支持按照不同级别发送消息到不同目标对象。
 - 简化Prometheus分组配置，支持按照具体消息发送到单个或多个接收方。
 - 增加手机号码配置项和号码自动轮询配置，可固定发送给单一个人告警信息，也可以通过自动轮询的方式发送到多个人员且支持按照不同日期发送到不同人员。
 - 增加Dashboard
 - 增加代理支持，感谢[@pzqu](https://github.com/pzqu)提交
--------------------------------------

## PrometheusAlert全家桶新手入门

* [PrometheusAlert配置]
  * [安装部署PrometheusAlert](doc/readme/install.md)
  * [PrometheusAlert接口说明](doc/readme/restful.md)
  * [PrometheusAlert配置文件说明](doc/readme/setup.md)
    * [其他配置](doc/readme/other.md)
* [告警源配置]
  * [Prometheus配置](doc/readme/prometheus.md)
  * [graylog配置](doc/readme/graylog.md)
  * [grafana配置](doc/readme/grafana.md)
  * [zabbix配置](doc/readme/zabbix.md)
* [告警接收目标配置]
  * [钉钉告警配置](doc/readme/dingding.md)
  * [企业微信告警配置](doc/readme/wechat.md)
  * [阿里云短信和电话告警配置](doc/readme/aliyun.md)
  * [腾讯云短信和电话告警配置](doc/readme/tx.md)
  * [华为云短信配置](doc/readme/hw.md)
* [PrometheusAlert版本更新说明](https://github.com/feiyu563/PrometheusAlert/releases)
* [如何对PrometheusAlert进行二次开发](https://github.com/feiyu563/PrometheusAlert/wiki/%E5%A6%82%E4%BD%95%E8%BF%9B%E8%A1%8CPrometheusAlert%E4%BA%8C%E6%AC%A1%E5%BC%80%E5%8F%91)


--------------------------------------

项目源码
----

 - [PrometheusAlert][1]


  [1]: https://github.com/feiyu563/PrometheusAlert

FOR HELP
----
Email: 244217140@qq.com

![me](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/wx.png)
