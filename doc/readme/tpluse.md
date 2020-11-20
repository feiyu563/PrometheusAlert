## 使用自定义告警消息模版

--------------------------------------

自定义告警消息模版接口使用非常简单

- 打开PrometheusAlert Dashboard的模版管理页面`AlertTemplate`

  * 找到需要使用的自定义消息模版，复制表格中`路径`一列的地址内容，并将地址中`[xxxxx]`中的地址或手机号替换成你实际的配置，将其粘贴到对应的WebHook地址配置中即可。(注意事项：自定义模版中的手机号是可以忽略的，如果不在url中配置手机号参数，则会优先读取user.csv中的手机号，如未读取到，则会取app.conf中的默认手机号)

![dashboard-tpl-list](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/dashboard-tpl-list.png)
