# alertgroup告警组配置文档

[issue: 增加告警组的功能](https://github.com/feiyu563/PrometheusAlert/issues/250)

<br/>
<br/>

## 告警组介绍

由于之前是使用配置文件里默认的通知媒介地址(`wxurl, ddurl...`)，或自定义的写在 prometheus rule annotations 的地址。修改这个自定义告警地址很麻烦，虽然也可以使用 `vim` 或 `sed` 命令来批量操作。

借鉴了云平台监控的告警通知组这个功能，将自定义地址都配置到告警组，然后配置不同的告警组(如 sa, dev...)。通过 beego 框架配置文件的 section 功能，将不同的告警组配置为不同的 section。

编写一个函数从配置文件中取出这些地址，并将其去重和汇总，然后以一个包含了特定类型地址的变量传递给发送消息的函数。

目前告警组功能仅支持 `/prometheus/alert` 这个旧控制器里面的接口，暂不涉及到自定义模板的接口 `/prometheusalert`。

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

# demo 告警组，用于测试案例测试
[ag-demo]
wxurl=wxurl1,wxurl2
ddurl=ddurl1,ddurl1,
fsurl=fsurl1
email=email1,
phone=phone1,phone2
groupid=groupid1

# sa 组
[ag-sa]
wxurl=wxurl1,wxurl2
ddurl=ddurl1
phone=13x,15x

# 自定义的告警组配置
include "alertgroup.conf"
```

`alertgroup.conf`:

```conf
# ops 组
[ops]
ddurl=ddurl1,ddurl2
fsurl=fsurl1
phone=17x,18x
groupid=groupid1

# dev 组
[dev]
wxurl=wxurl3
ddurl=ddurl3
fsurl=fsurl3
phone=13x,17x,18x
```

## 告警组使用

注意:

- 如果 annotations 配置了告警组，但 `app.conf` 配置里未配置告警组，则会使用配置文件里默认的那个地址(wxurl, ddurl...)。
- 如果 annotations 配置了告警组，并且 `app.conf` 配置里有配置告警组，则会使用告警组里面的地址。
- 每次修改通知地址，就只需要修改告警组里面的地址，而不用去修改 promtheus rules 的地址了。

<br/>

旧的 prometheus rule annotation 配置示例:

```yml
annotations:
  summary: "xxx"
  description: "xxx"
  wxurl: wxurl1,wxurl2
  ddurl: ddurl1
  fsurl: fsurl1
  mobile: 13x,15x
```

使用告警组的 prometheus rule annotations 配置示例:

```yml
annotations:
  summary: "xxx"
  description: "xxx"
  alertgroup: "ag-sa,dev"
```

<br/>
<br/>

## 调试告警组

文件 `conf/prometheus-demo.json` 是一个 prometheus 告警信息的示例 json 内容。我们可以通过修改此 json，然后通过 postman 将自定义的告警消息发送到接口 (`http://127.0.0.1:8080/prometheus/alert`)，来调试告警组和通知消息等。
