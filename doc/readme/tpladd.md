## 自定义告警消息模版使用说明

--------------------------------------

自定义告警消息模版可以支持任意带有WebHook服务的系统接入到PrometheusAlert上。

使用该功能需要使用者对go语言的template模版有一些初步了解，可以参考默认模版的一些语法来进行自定义。

模版数据等信息均存储在程序目录的下的`db/PrometheusAlertDB.db`中。

下面以添加Prometheus的自定义告警消息模版为例讲解如何添加自定义模版

- 开始之前，请先临时更改你的Alertmanager的配置，将所有告警信息都转发到PrometheusAlert自定义接口,参考如下：

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
  - url: 'http://[YOUR-PrometheusAlert-URL]/prometheusalert'
```

配置完成后，重启或者reload Alertmanager，是配置生效。

- 可手动或等待Prometheus告警触发后，去PrometheusAlert中查看收到的日志消息。找到类似下面的内容：

```
2020/05/21 10:58:17.850 [D] [value.go:460]  [1590029897850034963] {"receiver":"prometheus-alert-center","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"TargetDown","index":"1","instance":"example-1","job":"example","level":"2","service":"example"},"annotations":{"description":"target was down! example dev /example-1 was down for more than 120s.","level":"2","timestamp":"2020-05-21 02:58:07.829 +0000 UTC"},"startsAt":"2020-05-21T02:58:07.830216179Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"https://prometheus-alert-center/graph?g0.expr=up%7Bjob%21%3D%22kubernetes-pods%22%2Cjob%21%3D%22kubernetes-service-endpoints%22%7D+%21%3D+1\u0026g0.tab=1","fingerprint":"e2a5025853d4da64"}],"groupLabels":{"instance":"example-1"},"commonLabels":{"alertname":"TargetDown","index":"1","instance":"example-1","job":"example","level":"2","service":"example"},"commonAnnotations":{"description":"target was down! example dev /example-1 was down for more than 120s.","level":"2","timestamp":"2020-05-21 02:58:07.829 +0000 UTC"},"externalURL":"https://prometheus-alert-center","version":"4","groupKey":"{}/{job=~\"^(?:.*)$\"}:{instance=\"example-1\"}"}
```

- 继续截取日志中的JSON内容，通过任意json格式化工具进行格式化如下：

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

* 添加到Dashboard中，并选择对应模版类型和用途等信息，注意模版名称一定不要重复,且一定要是英文字符。

![tpladd1](https://gitee.com/feiyu563/PrometheusAlert/raw/master/doc/tpladd1.png)

* 继续对新添加的模版进行测试即可