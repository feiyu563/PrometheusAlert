# 自定义告警消息模版使用说明

--------------------------------------

### 自定义告警消息模版可以支持任意带有WebHook服务的系统接入到PrometheusAlert上。

### PrometheusAlert 原理：

#### PrometheusAlert 的自定义模板接口 `/prometheusalert` 专门用来接收外部服务传入的 HTTP POST 的 JSON 数据（通常带有WebHook的软件都可以接入PrometheusAlert），`/prometheusalert` 接口在接收到JSON数据后，会根据Url中的`tpl`参数去查找对应的自定义模板，通过找到的模板去渲染JSON以便输出渲染后的文本，最后将渲染后的文本转发给指定的目标，如钉钉机器人，从而实现钉钉告警。

#### 流程参考：

`XXX-WebHook` --> `POST-JSON` --> `/prometheusalert?type=dd&tpl=prometheus-dingding&ddurl=https://oapi.dingtalk.com/robot/send?access_token=xxxx` --> `PrometheusAlert通过tpl模版prometheus-dingding渲染收到的JSON` --> `将渲染后的消息文本发送给https://oapi.dingtalk.com/robot/send?access_token=xxxx` --> `钉钉机器人完成告警`

----------------------------------------------------------------------

### 下面以Prometheus为例演示下完整的使用自定义模板流程

----------------------------------------------------------------------
## 一、配置WebHook将消息转发给PrometheusAlert
----------------------------------------------------------------------

### 1.更改Alertmanager的配置，将所有告警信息都转发到PrometheusAlert自定义接口,参考如下：

  * 这么配置主要是为了方便调试和获取到Prometheus告警消息接口的JSON协议内容，其他的如grafana、graylog、sonarqube等支持WebHook的系统可以直接在控制台页面配置上`http://[YOUR-PrometheusAlert-URL]/prometheusalert`自定义消息模版的接口即可

```
global:
  resolve_timeout: 5m
route:
  group_by: ['instance']
  group_wait: 10m
  group_interval: 10s
  repeat_interval: 10m
  receiver: 'PrometheusAlert'
receivers:
- name: 'PrometheusAlert'
  webhook_configs:
  - url: 'http://[YOUR-PrometheusAlert-URL]/prometheusalert?type=dd&tpl=prometheus-dd'
```

配置完成后，重启或者reload Alertmanager，使配置生效。

- 可手动或等待Prometheus告警触发后，去PrometheusAlert中查看收到的日志消息。找到类似下面的内容（日志文件在log目录下）,json格式的消息一般都是`{`开头：

```
2020/05/21 10:58:17.850 [D] [value.go:460]  [1590029897850034963] {"receiver":"prometheus-alert-center","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"TargetDown","index":"1","instance":"example-1","job":"example","level":"2","service":"example"},"annotations":{"description":"target was down! example dev /example-1 was down for more than 120s.","level":"2","timestamp":"2020-05-21 02:58:07.829 +0000 UTC"},"startsAt":"2020-05-21T02:58:07.830216179Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"https://prometheus-alert-center/graph?g0.expr=up%7Bjob%21%3D%22kubernetes-pods%22%2Cjob%21%3D%22kubernetes-service-endpoints%22%7D+%21%3D+1\u0026g0.tab=1","fingerprint":"e2a5025853d4da64"}],"groupLabels":{"instance":"example-1"},"commonLabels":{"alertname":"TargetDown","index":"1","instance":"example-1","job":"example","level":"2","service":"example"},"commonAnnotations":{"description":"target was down! example dev /example-1 was down for more than 120s.","level":"2","timestamp":"2020-05-21 02:58:07.829 +0000 UTC"},"externalURL":"https://prometheus-alert-center","version":"4","groupKey":"{}/{job=~\"^(?:.*)$\"}:{instance=\"example-1\"}"}
```

- 截取日志中的JSON内容（主要是`{.....}`这部分），通过任意[json格式化工具](https://www.bejson.com/)进行格式化如下：

```
{
	"receiver": "prometheus-alert-center",
	"status": "firing",
	"alerts": [{
		"status": "firing",
		"labels": {
			"alertname": "TargetDown",
			"index": "1",
			"instance": "example-1",
			"job": "example",
			"level": "2",
			"service": "example"
		},
		"annotations": {
			"description": "target was down! example dev /example-1 was down for more than 120s.",
			"level": "2",
			"timestamp": "2020-05-21 02:58:07.829 +0000 UTC"
		},
		"startsAt": "2020-05-21T02:58:07.830216179Z",
		"endsAt": "0001-01-01T00:00:00Z",
		"generatorURL": "https://prometheus-alert-center/graph?g0.expr=up%7Bjob%21%3D%22kubernetes-pods%22%2Cjob%21%3D%22kubernetes-service-endpoints%22%7D+%21%3D+1\u0026g0.tab=1",
		"fingerprint": "e2a5025853d4da64"
	}],
	"groupLabels": {
		"instance": "example-1"
	},
	"commonLabels": {
		"alertname": "TargetDown",
		"index": "1",
		"instance": "example-1",
		"job": "example",
		"level": "2",
		"service": "example"
	},
	"commonAnnotations": {
		"description": "target was down! example dev /example-1 was down for more than 120s.",
		"level": "2",
		"timestamp": "2020-05-21 02:58:07.829 +0000 UTC"
	},
	"externalURL": "https://prometheus-alert-center",
	"version": "4",
	"groupKey": "{}/{job=~\"^(?:.*)$\"}:{instance=\"example-1\"}"
}
```

----------------------------------------------------------------------
## 二、对照JSON新建自定义模板(为方便调试，编写模版可以在web页面的MarkDownTest进行)
----------------------------------------------------------------------

* 然后对照该JSON开始编写模版,并在Dashboard上进行添加,示例模版如下：

```
{{ $var := .externalURL}}{{ range $k,$v:=.alerts }}
## [Prometheus告警信息]({{$v.generatorURL}})
#### [{{$v.labels.alertname}}]({{$var}})
###### 告警级别：{{$v.labels.level}}
###### 开始时间：{{$v.startsAt}}
###### 结束时间：{{$v.endsAt}}
###### 故障主机IP：{{$v.labels.instance}}
##### {{$v.annotations.description}}
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)
{{ end }}
```

* 添加到Dashboard中，并选择对应模版类型和模版用途，注意模版名称一定不要重复,且一定要是英文字符。

![tpladd1](../images/tpladd1.png)


* 添加完自定义模板后，一定要点击保存。

----------------------------------------------------------------------
## 三、测试自定义模板
----------------------------------------------------------------------

### 1.对新添加的模版进行测试

- 打开PrometheusAlert Dashboard的模版管理页面`AlertTemplate`

  * 在表格中找到刚刚创建的自定义模版，点击右侧的模版测试按钮，进入模版测试页面

![tpladd1](../images/tpltest1.png)

- 将之前从PrometheusAlert日志中提取的JSON填入`消息协议JSON内容`文本框中，且输入钉钉机器人地址(如模版的类型不是钉钉，模版测试页面的地址输入框显示会不同名称，如微信机器人地址等)

- 继续点击模版测试按钮即可对新添加的模版进行测试，如模版没有错误，将会收到对应的钉钉消息，如无法收到钉钉消息，请检查模版是否有什么地方配置错误

----------------------------------------------------------------------
## 四、使用自定义模板
----------------------------------------------------------------------

### 1.自定义告警消息模版接口使用非常简单

- 打开PrometheusAlert Dashboard的模版管理页面`AlertTemplate`

  * 找到需要使用的自定义消息模版，复制表格中`路径`一列的地址内容，并将地址中`[xxxxx]`中的地址或手机号替换成你实际的配置，将其粘贴到对应的WebHook地址配置中即可。(注意事项：自定义模版中的手机号是可以忽略的，如果不在url中配置手机号参数，则会优先读取user.csv中的手机号，如未读取到，则会取app.conf中的默认手机号)

  * 如prometheus alertmanager配置如下：
```
- name: 'prometheusalert-all'
  webhook_configs:
  - url: 'http://[prometheusalert_url]:8080/prometheusalert?type=dd&tpl=prometheus-dd&ddurl=钉钉机器人地址'
  #注意：url参数中 ddurl、wxurl、fsurl、phone、email、wxuser、wxparty、wxtag、groupid等可不写，如不写这些参数，则会默认去读取配置文件中的对应参数发送消息
```

![dashboard-tpl-list](../images/template.png)


至此完成！其他的一些支持WebHook的系统也可以通过同样的方式和流程来添加。