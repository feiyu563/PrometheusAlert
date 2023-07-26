# 自定义模版函数

--------------------------------------

## 自定义模版函数和使用（兼容alertmanager模板函数`toUpper、toLower、title、join、match、safeHtml、reReplaceAll、stringSlice`）
----------------------------------------------------------------------

### 1 `GetCSTtime` 函数仅支持在PrometheusAlert的自定义模版中使用，该函数主要用于强制将时间字段时区从UTC转换到CST

目前支持两种使用方式：

- 取的当前时间 `{{GetCSTtime ""}}` ,如：

```
{{ $var := .externalURL}}{{ range $k,$v:=.alerts }}
{{if eq $v.status "resolved"}}
## [Prometheus恢复信息]({{$v.generatorURL}})
#### [{{$v.labels.alertname}}]({{$var}})
###### 告警级别：{{$v.labels.level}}
###### 开始时间：{{GetCSTtime $v.startsAt}}
###### 结束时间：{{GetCSTtime $v.endsAt}}
###### 故障主机IP：{{$v.labels.instance}}
##### 当前时间 {{GetCSTtime ""}} {{$v.annotations.description}}  #{{GetCSTtime ""}} 即会自动获取当前的时间嵌入到消息中
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)
{{else}}
## [Prometheus告警信息]({{$v.generatorURL}})
#### [{{$v.labels.alertname}}]({{$var}})
###### 告警级别：{{$v.labels.level}}
###### 开始时间：{{GetCSTtime $v.startsAt}}
###### 结束时间：{{GetCSTtime $v.endsAt}}
###### 故障主机IP：{{$v.labels.instance}}
##### {{$v.annotations.description}}
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)
{{end}}
{{ end }}
```

- 转换UTC时间到CST时间 `{{GetCSTtime $v.startsAt}}` ,如

```
{{ $var := .externalURL}}{{ range $k,$v:=.alerts }}
{{if eq $v.status "resolved"}}
## [Prometheus恢复信息]({{$v.generatorURL}})
#### [{{$v.labels.alertname}}]({{$var}})
###### 告警级别：{{$v.labels.level}}
###### 开始时间：{{GetCSTtime $v.startsAt}}   #{{GetCSTtime $v.startsAt}} 中传入Prometheus告警消息的时间字段即可将该传入的时间转换为CST时间
###### 结束时间：{{GetCSTtime $v.endsAt}}
###### 故障主机IP：{{$v.labels.instance}}
##### {{$v.annotations.description}}
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)
{{else}}
## [Prometheus告警信息]({{$v.generatorURL}})
#### [{{$v.labels.alertname}}]({{$var}})
###### 告警级别：{{$v.labels.level}}
###### 开始时间：{{GetCSTtime $v.startsAt}}
###### 结束时间：{{GetCSTtime $v.endsAt}}
###### 故障主机IP：{{$v.labels.instance}}
##### {{$v.annotations.description}}
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)
{{end}}
{{ end }}
```



### 2 `TimeFormat` 函数仅支持在PrometheusAlert的自定义模版中使用，该函数主要用于格式化时间显示

如下示例将prmetheus的告警时间格式改为：2006-01-02T15:04:05

```
{{ $var := .externalURL}}{{ range $k,$v:=.alerts }}
{{if eq $v.status "resolved"}}
## [Prometheus恢复信息]({{$v.generatorURL}})
#### [{{$v.labels.alertname}}]({{$var}})
###### 告警级别：{{$v.labels.level}}
###### 开始时间：{{TimeFormat $v.startsAt "2006-01-02T15:04:05"}}  
###### 结束时间：{{TimeFormat $v.endsAt "2006-01-02T15:04:05"}}
###### 故障主机IP：{{$v.labels.instance}}
##### 当前时间 {{GetCSTtime ""}} {{$v.annotations.description}}  #{{GetCSTtime ""}} 即会自动获取当前的时间嵌入到消息中
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)
{{else}}
## [Prometheus告警信息]({{$v.generatorURL}})
#### [{{$v.labels.alertname}}]({{$var}})
###### 告警级别：{{$v.labels.level}}
###### 开始时间：{{TimeFormat $v.startsAt "2006-01-02T15:04:05"}}
###### 结束时间：{{TimeFormat $v.endsAt "2006-01-02T15:04:05"}}
###### 故障主机IP：{{$v.labels.instance}}
##### {{$v.annotations.description}}
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)
{{end}}
{{ end }}
```



### 3 `GetTime` 函数仅支持在PrometheusAlert的自定义模版中使用，该函数主要用于将`毫秒或秒`级时间戳转换为时间字符

特别说明：`GetTime`函数支持字符和数值类型参数，字符型支持秒级和毫秒级时间戳的处理，数值类型暂时只支持秒级时间戳处理。

目前支持两种使用方式：

- 使用默认时间字符串格式输出 `{{GetTime .Timestamp}}` ,如：

```
ALiYun {{.AlertState}}信息
>**{{.AlertName}}**
>告警级别: {{.TriggerLevel}}
开始时间: {{GetTime .Timestamp}} //输出时间格式：2006-01-02T15:04:05
故障主机: {{.InstanceName}}
------------详细信息--------------
metricName: {{.MetricName}}
expression: {{.Expression}}
signature: {{.Signature}}
metricProject: {{.MetricProject}}
userId: {{.UserId}}
namespace: {{.Namespace}}
preTriggerLevel: {{.PreTriggerLevel}}
ruleId: {{.RuleId}}
dimensions: {{.Dimensions}}
**当前值：{{.CurValue}}**
```

- 指定输出时间格式输出 `{{GetTime .Timestamp "2006/01/02 15:04:05"}}` ,如


```
ALiYun {{.AlertState}}信息
>**{{.AlertName}}**
>告警级别: {{.TriggerLevel}}
开始时间: {{GetTime .Timestamp}} //输出时间格式：2006/01/02 15:04:05
故障主机: {{.InstanceName}}
------------详细信息--------------
metricName: {{.MetricName}}
expression: {{.Expression}}
signature: {{.Signature}}
metricProject: {{.MetricProject}}
userId: {{.UserId}}
namespace: {{.Namespace}}
preTriggerLevel: {{.PreTriggerLevel}}
ruleId: {{.RuleId}}
dimensions: {{.Dimensions}}
**当前值：{{.CurValue}}**
```

### 4 `SplitString` 函数仅支持在PrometheusAlert的自定义模版中使用，该函数主要用于截取字符串

用法示例: 

原始json为：{"Instance": "192.168.1.110:9100"}

需求：需要截取IP值 `192.168.1.110`

示例：`{{SplitString .Instance 0 -5}}` 或 `{{SplitString .Instance 0 13}}`

参数解释：SplitString需要三个参数，原始字符串、截取开始位置、截取结束位置；截取结束位置如果为负数，则相当于取字符串总长度+负数得出的位置。

特别说明：`SplitString`函数

目前支持两种使用方式：

- 使用默认时间字符串格式输出 `{{GetTime .Timestamp}}` ,如：

```
ALiYun {{.AlertState}}信息
>**{{.AlertName}}**
>告警级别: {{.TriggerLevel}}
开始时间: {{.Timestamp}}
故障主机: {{SplitString .Instance 0 -5}}
------------详细信息--------------
metricName: {{.MetricName}}
expression: {{.Expression}}
signature: {{.Signature}}
metricProject: {{.MetricProject}}
userId: {{.UserId}}
namespace: {{.Namespace}}
preTriggerLevel: {{.PreTriggerLevel}}
ruleId: {{.RuleId}}
dimensions: {{.Dimensions}}
**当前值：{{.CurValue}}**
```