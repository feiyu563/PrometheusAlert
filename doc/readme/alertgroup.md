# alertgroup告警组配置文档

[issue: 增加告警组的功能](https://github.com/feiyu563/PrometheusAlert/issues/250)

<br/>
<br/>

## 告警组介绍

> **注意**：<br/>
> 告警组功能适合以逗号隔开的地址，组装这一类的地址。<br/>
> 自定义模板中的 模板类型(type)、模板(tpl)、webhookContentType、企微应用、at、rr、split等，便不太适合写入告警组，因此这里并没有处理这些值，还是按照原来写在URL参数上。

<br/>

借鉴了云平台监控的告警通知组这个功能，将自定义地址都配置到告警组，然后配置不同的告警组(如 sa, dev...)。通过 beego 框架配置文件的 section 功能，将不同的告警组配置为不同的 section。

编写一个函数从配置文件中取出这些地址，并将其去重和汇总，然后以一个包含了特定类型地址的变量传递给发送消息的函数。

旧接口，之前是使用配置文件里默认的通知媒介地址(`wxurl, ddurl...`)，或自定义的写在 prometheus rule annotations 的地址。修改这个自定义告警地址很麻烦，虽然也可以使用 `vim` 或 `sed` 命令来批量操作。

自定义模板接口，不必将很长的多个地址写到 URL，而只需要修改和配置 `app.conf` 文件，要方便一些。

旧的接口 `/prometheus/alert` 和自定义模板接口 `/prometheusalert` 有一些不同，下面会介绍各自的用法。

<br/>
<br/>

## 告警组配置

示例配置文件 `app-example.conf` 底部新增了了告警组配置。
由于涉及到 beego conf section，建议把告警组相关配置放置到最底部，或者通过 `include` 包含到另一个配置文件，便于修改。

下面的示例配置定义了几个告警组，每个告警组有自己的通知地址。

`app.conf`:

```conf
#---------------------↓告警组-----------------------
# 是否启用告警组功能
open-alertgroup=1

# sa 组
[ag-sa]
wxurl=wxurl1,wxurl2
ddurl=ddurl1
phone=13x,15x

# 自定义的告警组配置
include "alertgroup.conf"
```

扩展的告警组配置文件 `alertgroup.conf`:

```conf
# ops 示例组
[ops]
ddurl=ddurl1,ddurl2
fsurl=fsurl1
phone=17x,18x
groupid=groupid1

# dev 示例组
[dev]
wxurl=wxurl3
ddurl=ddurl3
fsurl=fsurl3
phone=13x,17x,18x

# 自定义模板告警组示例，目前仅处理了以下这些参数
[customtpl]
wxurl=wxurl1,wxurl2
ddurl=ddurl1,ddurl2
fsurl=fsurl1,fsurl2
email=email1,email2
phone=phone1,phone2
groupid=groupid1,groupid2
webhookurl=webhookurl1,webhookurl2
```

<br/>
<br/>

## 告警组使用

<br/>

### 旧接口 /prometheus/alert 的使用

注意:

- 告警组并不影响原来的 annotations 或默认配置的使用。
- 如果 annotations 配置了告警组，但 `app.conf` 配置里未配置告警组，则会使用配置文件里默认的那个地址(wxurl, ddurl...)。
- 如果 annotations 配置了告警组，并且 `app.conf` 配置里有配置告警组，则会使用告警组里面的地址。
- 地址的判断先后顺序：告警组-annotations-配置文件，哪个有值使用哪个。
- 每次修改通知地址，就只需要修改告警组里面的地址，而不用去修改 annotations 里的地址了。

<br/>

annotations 示例：

```yml
# 旧的 annotation 配置示例
annotations:
  summary: "xxx"
  description: "xxx"
  wxurl: wxurl1,wxurl2
  ddurl: ddurl1
  fsurl: fsurl1
  mobile: 13x,15x
```

```yml
# 使用告警组的 prometheus rule annotations 配置示例
annotations:
  summary: "xxx"
  description: "xxx"
  alertgroup: "ag-sa,dev"
```

<br/>
<br/>

### 自定义模板接口 /prometheusalert 的使用

注意：

- 告警组并不影响原来的 URL 传递参数的使用。
- 如果配置了告警组，且告警组中的地址不为空，则使用告警组中配置的地址。
- 地址的判断先后顺序：告警组-URL参数-配置文件，哪个不为空使用哪个。
- 自定义模板中的 模板类型(type)、模板(tpl)、webhookContentType、企微应用、at、rr、split等，便不太适合写入告警组，因此这里并没有处理这些值，还是按照原来写在URL参数上。
- 在 URL 参数上配置 `alertgroup=xxx`，然后将具体的地址写到配置文件的告警组下面。

自定义接口示例，这样在 alertmanager 或其它软件中接入地址的时候，就不需要带上很长的具体各个媒介的地址，而只需要写告警组就可以了。

```txt
# prometheus-dd
http://127.0.0.1:8080/prometheusalert?type=dd&tpl=prometheus-dd&alertgroup=customtpl
http://127.0.0.1:8080/prometheusalert?type=dd&tpl=prometheus-dd&alertgroup=customtpl&at=188xxx

# prometheus-wx
http://127.0.0.1:8080/prometheusalert?type=wx&tpl=prometheus-wx&alertgroup=sa,dev
http://127.0.0.1:8080/prometheusalert?type=wx&tpl=prometheus-wx&alertgroup=sa&at=zhangsan

# prometheus-fs
http://127.0.0.1:8080/prometheusalert?type=fs&tpl=prometheus-fs&&alertgroup=sa&at=zhangsan@xxx.com

# 其它效果类似
# 可以通过下面的测试告警组文档先测试效果
```

<br/>
<br/>

## 调试告警组

通过下面的示例 JSON 内容，在 postman 将示例告警内容发送到接口进行测试。

旧接口 (`http://127.0.0.1:8080/prometheus/alert`) 的示例告警 JSON 内容：

```json
{
  "receiver": "prometheus-alert-center",
	"status": "firing",
	"alerts": [
		{
			"status": "firing",
			"labels": {
				"alertname": "TestAlert",
				"instance": "localhost",
				"level": "1",
				"severity": "warning",
				"job": "node_exporter",
				"hostgroup": "test",
				"hostname": "ecs01"
			},
			"annotations": {
				"description": "This is a test alert",
				"summary": "Test Alert Summary",
				"alertgroup": "sa,dev"
			},
			"startsAt": "2023-06-25T10:00:00Z",
			"endsAt": "2023-06-25T11:00:00Z",
			"generatorURL": "http://localhost/alerts"
		}
	],
	"externalURL": "http://localhost/prometheus"
}
```

<br/>

自定义接口（`http://127.0.0.1:8080/prometheusalert?type=xx&tpl=xxx&alertgroup=xxx`）的示例告警 JSON 内容：

```json
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
