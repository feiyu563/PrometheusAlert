## 自定义模版函数

--------------------------------------

`GetCSTtime` 自定义函数仅支持在PrometheusAlert的自定义模版中使用

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

- 转换UTC时间到CST时间 `{{GetCSTtime ""}}` ,如

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