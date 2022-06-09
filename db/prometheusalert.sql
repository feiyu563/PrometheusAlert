-- ----------------------------
-- Records of prometheus_alert_d_b
-- ----------------------------
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (1, 'dd', 'Prometheus', 'prometheus-dd', '{{ $var := .externalURL}}{{ range $k,$v:=.alerts }}
{{if eq $v.status "resolved"}}
## [Prometheus恢复信息]({{$v.generatorURL}})
#### [{{$v.labels.alertname}}]({{$var}})
###### 告警级别：{{$v.labels.level}}
###### 开始时间：{{TimeFormat $v.startsAt "2006-01-02 15:04:05"}}
###### 结束时间：{{TimeFormat $v.endsAt "2006-01-02 15:04:05"}}
###### 故障主机IP：{{$v.labels.instance}}
##### {{$v.annotations.description}}
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)
{{else}}
## [Prometheus告警信息]({{$v.generatorURL}})
#### [{{$v.labels.alertname}}]({{$var}})
###### 告警级别：{{$v.labels.level}}
###### 开始时间：{{TimeFormat $v.startsAt "2006-01-02 15:04:05"}}
###### 结束时间：{{TimeFormat $v.endsAt "2006-01-02 15:04:05"}}
###### 故障主机IP：{{$v.labels.instance}}
##### {{$v.annotations.description}}
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)
{{end}}
{{ end }}
{{ $urimsg:=""}}{{ range $key,$value:=.commonLabels }}{{$urimsg =  print $urimsg $key "%3D%22" $value "%22%2C" }}{{end}}[*** 点我屏蔽该告警]({{$var}}/#/silences/new?filter=%7B{{SplitString $urimsg 0 -3}}%7D)', '2022-05-26 10:00:05');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (2, 'wx', 'Prometheus', 'prometheus-wx', '{{ $var := .externalURL}}{{ range $k,$v:=.alerts }}{{if eq $v.status "resolved"}}[PROMETHEUS-恢复信息]({{$v.generatorURL}})
> **[{{$v.labels.alertname}}]({{$var}})**
> <font color="info">告警级别:</font> {{$v.labels.level}}
> <font color="info">开始时间:</font> {{$v.startsAt}}
> <font color="info">结束时间:</font> {{$v.endsAt}}
> <font color="info">故障主机IP:</font> {{$v.labels.instance}}
> <font color="info">**{{$v.annotations.description}}**</font>{{else}}[PROMETHEUS-告警信息]({{$v.generatorURL}})
> **[{{$v.labels.alertname}}]({{$var}})**
> <font color="warning">告警级别:</font> {{$v.labels.level}}
> <font color="warning">开始时间:</font> {{$v.startsAt}}
> <font color="warning">结束时间:</font> {{$v.endsAt}}
> <font color="warning">故障主机IP:</font> {{$v.labels.instance}}
> <font color="warning">**{{$v.annotations.description}}**</font>{{end}}{{ end }}
{{ $urimsg:=""}}{{ range $key,$value:=.commonLabels }}{{$urimsg =  print $urimsg $key "%3D%22" $value "%22%2C" }}{{end}}[*** 点我屏蔽该告警]({{$var}}/#/silences/new?filter=%7B{{SplitString $urimsg 0 -3}}%7D)', '2022-05-26 09:59:49');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (3, 'fs', 'Prometheus', 'prometheus-fs', '{{ $var := .externalURL}}{{ range $k,$v:=.alerts }}{{if eq $v.status "resolved"}}**[Prometheus恢复信息]({{$v.generatorURL}})**
*[{{$v.labels.alertname}}]({{$var}})*
告警级别：{{$v.labels.level}}
开始时间：{{TimeFormat $v.startsAt "2006-01-02 15:04:05"}}
结束时间：{{TimeFormat $v.endsAt "2006-01-02 15:04:05"}}
故障主机IP：{{$v.labels.instance}}
**{{$v.annotations.description}}**{{else}}**[Prometheus告警信息]({{$v.generatorURL}})**
*[{{$v.labels.alertname}}]({{$var}})*
告警级别：{{$v.labels.level}}
开始时间：{{TimeFormat $v.startsAt "2006-01-02 15:04:05"}}
结束时间：{{TimeFormat $v.endsAt "2006-01-02 15:04:05"}}
故障主机IP：{{$v.labels.instance}}
**{{$v.annotations.description}}**{{end}}{{ end }}
{{ $urimsg:=""}}{{ range $key,$value:=.commonLabels }}{{$urimsg =  print $urimsg $key "%3D%22" $value "%22%2C" }}{{end}}[*** 点我屏蔽该告警]({{$var}}/#/silences/new?filter=%7B{{SplitString $urimsg 0 -3}}%7D)', '2022-05-26 10:10:03');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (4, 'txdx', 'Prometheus', 'prometheus-dx', '{{ range $k,$v:=.alerts }}{{if eq $v.status "resolved"}}
[Prometheus恢复信息]
{{$v.labels.alertname}}
告警级别：{{$v.labels.level}}
故障主机IP：{{$v.labels.instance}}
{{$v.annotations.description}}
{{else}}
[Prometheus告警信息]
{{$v.labels.alertname}}
告警级别：{{$v.labels.level}}
故障主机IP：{{$v.labels.instance}}
{{$v.annotations.description}}
{{end}}
{{ end }}', '2021-09-24 08:44:36');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (5, 'txdh', 'Prometheus', 'prometheus-dh', '{{ range $k,$v:=.alerts }}{{if eq $v.status "resolved"}}恢复信息故障主机IP{{$v.labels.instance}}{{$v.annotations.description}}{{else}}告警信息故障主机IP{{$v.labels.instance}}{{$v.annotations.description}}{{end}}{{ end }}', '2020-07-29 09:39:56');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (6, 'email', 'Prometheus', 'prometheus-email', '{{ $var := .externalURL}}{{ range $k,$v:=.alerts }}
{{if eq $v.status "resolved"}}
<h1><a href ={{$v.generatorURL}}>Prometheus恢复信息</a></h1>
<h2><a href ={{$var}}>{{$v.labels.alertname}}</a></h2>
<h5>告警级别：{{$v.labels.level}}</h5>
<h5>开始时间：{{$v.startsAt}}</h5>
<h5>结束时间：{{$v.endsAt}}</h5>
<h5>故障主机IP：{{$v.labels.instance}}</h5>
<h3>{{$v.annotations.description}}</h3>
<img src=https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png />
{{else}}
<h1><a href ={{$v.generatorURL}}>Prometheus告警信息</a></h1>
<h2><a href ={{$var}}>{{$v.labels.alertname}}</a></h2>
<h5>告警级别：{{$v.labels.level}}</h5>
<h5>开始时间：{{$v.startsAt}}</h5>
<h5>结束时间：{{$v.endsAt}}</h5>
<h5>故障主机IP：{{$v.labels.instance}}</h5>
<h3>{{$v.annotations.description}}</h3>
<img src=https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png />
{{end}}
{{ end }}', '2020-09-30 09:46:05');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (7, 'dd', 'Graylog2', 'graylog2-dd', '## [Graylog2告警信息](http://graylog.org)
#### {{.check_result.result_description}}
{{ range $k,$v:=.check_result.matching_messages }}
###### 告警索引：{{$v.index}}
###### 开始时间：{{GetCSTtime $v.timestamp}}
###### 告警主机：{{$v.fields.gl2_remote_ip}}:{{$v.fields.gl2_remote_port}}
##### {{$v.message}}
{{end}}
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)', '2020-07-29 09:45:33');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (8, 'wx', 'Graylog2', 'graylog2-wx', '[Graylog2告警信息](http://graylog.org)
>**{{.check_result.result_description}}**
{{ range $k,$v:=.check_result.matching_messages }}
>告警索引:{{$v.index}}
开始时间:{{GetCSTtime $v.timestamp}}
告警主机:{{$v.fields.gl2_remote_ip}}:{{$v.fields.gl2_remote_port}}
**{{$v.message}}**
{{end}}', '2020-07-29 09:49:58');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (9, 'fs', 'Graylog2', 'graylog2-fs', '**[Graylog2告警信息](http://graylog.org)**
*{{.check_result.result_description}}*
{{ range $k,$v:=.check_result.matching_messages }}
告警索引：{{$v.index}}
开始时间：{{GetCSTtime $v.timestamp}}
告警主机：{{$v.fields.gl2_remote_ip}}:{{$v.fields.gl2_remote_port}}
**{{$v.message}}**
{{end}}', '2020-09-30 10:08:04');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (10, 'txdx', 'Graylog2', 'graylog2-dx', '{{ range $k,$v:=.check_result.matching_messages }}
告警主机 {{$v.fields.gl2_remote_ip}}端口 {{$v.fields.gl2_remote_port}}告警消息 {{$v.message}}
{{end}}', '2020-07-29 09:55:27');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (20, 'txdh', 'Graylog2', 'graylog2-dh', '{{ range $k,$v:=.check_result.matching_messages }}
告警主机 {{$v.fields.gl2_remote_ip}}端口 {{$v.fields.gl2_remote_port}}告警消息 {{$v.message}}
{{end}}', '2020-07-29 09:56:14');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (27, 'email', 'Graylog2', 'graylog2-email', '<h1><a href =http://graylog.org>Graylog2告警信息</a></h1>
<h2>{{.check_result.result_description}}</h2>
{{ range $k,$v:=.check_result.matching_messages }}
<h5>告警索引：{{$v.index}}</h5>
<h5>开始时间：{{GetCSTtime $v.timestamp}}</h5>
<h5>告警主机：{{$v.fields.gl2_remote_ip}}:{{$v.fields.gl2_remote_port}}</h5>
<h3>{{$v.message}}</h3>
{{end}}
<img src=https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png />', '2020-07-29 09:59:08');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (28, 'dd', 'Graylog3', 'graylog3-dd', '## [Graylog3告警信息](.check_result.Event.Source)
#### {{.check_result.event_definition_description}}
{{ range $k,$v:=.backlog }}
###### 告警索引：{{$v.index}}
###### 开始时间：{{GetCSTtime $v.timestamp}}
###### 告警主机：{{$v.fields.gl2_remote_ip}}:{{$v.fields.gl_2_remote_port}}
##### {{$v.message}}
{{end}}
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)', '2020-07-29 10:00:32');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (29, 'wx', 'Graylog3', 'graylog3-wx', '[Graylog3告警信息](.check_result.Event.Source)
>**{{.check_result.event_definition_description}}**
{{ range $k,$v:=.backlog }}
>告警索引:{{$v.index}}
开始时间:{{GetCSTtime $v.timestamp}}
告警主机:{{$v.fields.gl2_remote_ip}}:{{$v.fields.gl_2_remote_port}}
**{{$v.message}}**
{{end}}', '2020-07-29 10:03:29');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (30, 'fs', 'Graylog3', 'graylog3-fs', '**[Graylog3告警信息](.check_result.Event.Source)**
*{{.check_result.event_definition_description}}*
{{ range $k,$v:=.backlog }}
告警索引：{{$v.index}}
开始时间：{{GetCSTtime $v.timestamp}}
告警主机：{{$v.fields.gl2_remote_ip}}:{{$v.fields.gl_2_remote_port}}
**{{$v.message}}**
{{end}}', '2020-09-30 09:52:11');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (31, 'txdx', 'Graylog3', 'graylog3-dx', '{{ range $k,$v:=.backlog }}
告警主机 {{$v.fields.gl2_remote_ip}}端口 {{$v.fields.gl_2_remote_port}}告警消息 {{$v.message}}
{{end}}', '2020-07-29 10:50:06');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (32, 'txdh', 'Graylog3', 'graylog3-dh', '{{ range $k,$v:=.backlog }}
告警主机 {{$v.fields.gl2_remote_ip}}端口 {{$v.fields.gl_2_remote_port}}告警消息 {{$v.message}}
{{end}}', '2020-07-29 10:50:30');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (33, 'email', 'Graylog3', 'graylog3-email', '<h1><a href =.check_result.Event.Source>Graylog3告警信息</a></h1>
<h2>{{.check_result.event_definition_description}}</h2>
{{ range $k,$v:=.backlog }}
<h5>告警索引：{{$v.index}}</h5>
<h5>开始时间：{{GetCSTtime $v.timestamp}}</h5>
<h5>告警主机：{{$v.fields.gl2_remote_ip}}:{{$v.fields.gl_2_remote_port}}</h5>
<h3>{{$v.message}}</h3>
{{end}}
<img src=https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png />', '2020-07-29 10:53:20');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (34, 'dd', 'Grafana', 'grafana-dd', '{{if eq .state "ok"}}
## [Grafana恢复信息]({{.ruleUrl}})
#### {{.ruleName}}
###### 告警级别：严重
###### 开始时间：{{GetCSTtime ""}}
##### {{.message}}
{{else}}
## [Grafana告警信息]({{.ruleUrl}})
#### {{.ruleName}}
###### 告警级别：严重
###### 开始时间：{{GetCSTtime ""}}
##### {{.message}}
{{end}}
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)', '2020-07-29 11:40:33');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (35, 'wx', 'Grafana', 'grafana-wx', '{{if eq .state "ok"}}
[Grafana恢复信息]({{.ruleUrl}})
>**{{.ruleName}}**
>告警级别:严重
开始时间:{{GetCSTtime ""}}
{{.message}}
{{else}}
[Grafana告警信息]({{.ruleUrl}})
>**{{.ruleName}}**
>告警级别:严重
开始时间:{{GetCSTtime ""}}
{{.message}}
{{end}}', '2020-07-29 11:41:07');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (36, 'fs', 'Grafana', 'grafana-fs', '{{if eq .state "ok"}}
**[Grafana恢复信息]({{.ruleUrl}})**
*{{.ruleName}}*
告警级别：严重
开始时间：{{GetCSTtime ""}}
**{{.message}}**
{{else}}
**[Grafana告警信息]({{.ruleUrl}})**
*{{.ruleName}}*
告警级别：严重
开始时间：{{GetCSTtime ""}}
**{{.message}}**
{{end}}', '2020-09-30 09:53:04');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (37, 'txdx', 'Grafana', 'grafana-dx', '{{if eq .state "ok"}}
Grafana恢复信息{{.message}}
{{else}}
Grafana告警信息{{.message}}
{{end}}', '2020-07-29 11:44:28');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (38, 'txdh', 'Grafana', 'grafana-dh', '{{if eq .state "ok"}}
Grafana恢复信息{{.message}}
{{else}}
Grafana告警信息{{.message}}
{{end}}', '2020-07-29 11:44:49');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (39, 'email', 'Grafana', 'grafana-email', '{{if eq .state "ok"}}
<h1><a href ={{.ruleUrl}}>Grafana恢复信息</a></h1>
<h2>{{.ruleName}}</h2>
<h5>告警级别：严重</h5>
<h5>开始时间：{{GetCSTtime ""}}</h5>
<h3>{{.message}}</h3>
{{else}}
<h1><a href ={{.ruleUrl}}>Grafana恢复信息</a></h1>
<h2>{{.ruleName}}</h2>
<h5>告警级别：严重</h5>
<h5>开始时间：{{GetCSTtime ""}}</h5>
<h3>{{.message}}</h3>
{{end}}
<img src=https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png />', '2020-07-29 11:52:25');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (40, 'dd', 'SonarQube', 'sonar-dd-example', '## [Sonar告警信息]({{.serverUrl}})
###### 检测状态：{{.status}}
###### 检测时间：{{.analysedAt}}
###### ---------------------------------
{{ range $k,$v:=.qualityGate.conditions}}
###### metric：{{$v.metric}}
###### errorThreshold：{{$v.errorThreshold}}
###### operator：{{$v.operator}}
###### status：{{$v.status}}
###### -----------------------------------
{{ end }}
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)', '2020-07-29 11:53:35');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (41, 'dd', 'Jenkins', 'jenkins-dd-example', '## [Jenkins构建信息]({{.buildUrl}})
###### Jenkins地址：[{{.buildUrl}}]({{.buildUrl}})
###### 构建项目：{{.projectName}}
###### 构建事件：{{.event}}
###### 构建名称：{{.buildName}}
###### 构建时间：{{GetCSTtime ""}}
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)', '2020-07-29 11:54:02');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (42, 'fs', 'Prometheus', 'prometheus-fsv2', '{{ $var := .externalURL}}{{ range $k,$v:=.alerts }}
{{if eq $v.status "resolved"}}
**[Prometheus恢复信息]({{$v.generatorURL}})**
*[{{$v.labels.alertname}}]({{$var}})*
告警级别：{{$v.labels.level}}
开始时间：{{$v.startsAt}}
结束时间：{{$v.endsAt}}
故障主机IP：{{$v.labels.instance}}
**{{$v.annotations.description}}**
{{else}}
**[Prometheus告警信息]({{$v.generatorURL}})**
*[{{$v.labels.alertname}}]({{$var}})*
告警级别：{{$v.labels.level}}
开始时间：{{$v.startsAt}}
结束时间：{{$v.endsAt}}
故障主机IP：{{$v.labels.instance}}
**{{$v.annotations.description}}**
{{end}}
{{ end }}', '2020-11-20 08:05:04');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (43, 'workwechat', 'Prometheus', 'prometheus-wechatapp', '{{ $var := .externalURL}}{{ range $k,$v:=.alerts }}
{{if eq $v.status "resolved"}}
[Prometheus恢复信息]($v.generatorURL}})
>**[{{$v.labels.alertname}}]({{$var}})**
>告警级别: {{$v.labels.level}}
开始时间: {{$v.startsAt}}
结束时间: {{$v.endsAt}}
故障主机IP: {{$v.labels.instance}}
**{{$v.annotations.description}}**
{{else}}
[Prometheus告警信息]($v.generatorURL}})
>**[{{$v.labels.alertname}}]({{$var}})**
>告警级别: {{$v.labels.level}}
开始时间: {{$v.startsAt}}
结束时间: {{$v.endsAt}}
故障主机IP: {{$v.labels.instance}}
**{{$v.annotations.description}}**
{{end}}
{{ end }}', '2021-01-18 08:13:55');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (44, 'rl', 'Prometheus', 'prometheus-ruliu', '{{ $var := .externalURL}}{{ range $k,$v:=.alerts }}
{{if eq $v.status "resolved"}}
## [Prometheus恢复信息]({{$v.generatorURL}})
#### [{{$v.labels.alertname}}]({{$var}})
###### 告警级别：{{$v.labels.level}}
###### 开始时间：{{$v.startsAt}}
###### 结束时间：{{$v.endsAt}}
###### 故障主机IP：{{$v.labels.instance}}
##### {{$v.annotations.description}}
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)
{{else}}
## [Prometheus告警信息]({{$v.generatorURL}})
#### [{{$v.labels.alertname}}]({{$var}})
###### 告警级别：{{$v.labels.level}}
###### 开始时间：{{$v.startsAt}}
###### 结束时间：{{$v.endsAt}}
###### 故障主机IP：{{$v.labels.instance}}
##### {{$v.annotations.description}}
![Prometheus](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)
{{end}}
{{ end }}', '2021-02-02 07:45:30');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (45, 'webhook', 'WebHook', 'prometheus-webhook', '{
{{ $var := index .alerts 0}}
{{if eq $var.status "resolved"}}
"title": "Prometheus恢复信息",
"prometheus-url": "{{$var.generatorURL}}",
"alert-name": "{{$var.labels.alertname}}",
"alertmanager-url": "{{.externalURL}}",
"alert-level": "{{$var.labels.level}}",
"alert-start-time": "{{$var.startsAt}}",
"alert-end-time": "{{$var.endsAt}}",
"alert-instance": "{{$var.labels.instance}}",
"message": "{{$var.annotations.description}}"
{{else}}
"title": "Prometheus告警信息",
"prometheus-url": "{{$var.generatorURL}}",
"alert-name": "{{$var.labels.alertname}}",
"alertmanager-url": "{{.externalURL}}",
"alert-level": "{{$var.labels.level}}",
"alert-start-time": "{{$var.startsAt}}",
"alert-end-time": "{{$var.endsAt}}",
"alert-instance": "{{$var.labels.instance}}",
"message": "{{$var.annotations.description}}"
{{end}}
}', '2021-04-29 08:04:12');
INSERT INTO prometheus_alert_d_b (id, tpltype, tpluse, tplname, tpl, created) VALUES (46, 'wx', 'ALiYun', 'aliyun', 'ALiYun {{.AlertState}}信息
>**{{.AlertName}}**
>告警级别: {{.TriggerLevel}}
开始时间: {{GetTime .Timestamp}}
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
**当前值：{{.CurValue}}**', '2021-07-14 06:57:31');
