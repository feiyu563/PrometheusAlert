## 测试自定义告警消息模版

--------------------------------------

接上篇添加自定义告警消息模版

- 打开PrometheusAlert Dashboard的模版管理页面`AlertTemplate`

  * 找到刚刚创建的自定义模版，点击右侧的模版测试按钮，进入模版测试页面

![tpladd1](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/tpltest1.png)

- 将之前从PrometheusAlert日志中提取的JSON填入`消息协议JSON内容`文本框中，且输入钉钉机器人地址(如模版的类型不是钉钉，模版测试页面的地址输入框显示会不同名称，如微信机器人地址等)

- 继续点击模版测试按钮即可对新添加的模版进行测试，如模版没有错误，将会收到对应的钉钉消息，如无法收到钉钉消息，请检查模版是否有什么地方配置错误
