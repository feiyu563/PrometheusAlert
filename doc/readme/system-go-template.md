# 进阶用法--go模版语法

--------------------------------------

PrometheusAlert的自定义模板基于go template而来，go template主要将json数据渲染成用户需要的文本

比如下面例子：
```
{
  "name": "张三",
  "age": 18,
  "teacher": ["TC1","TC2","TC3"],
  "like": {
    "balls": "football",
    "other": ["swim","climb"]
  }
}
```
比如我们现在需要在输出的文本（或者发送的消息文本）中要获取到json消息内的`name`、`age`、`teacher`，那么我们可以在模版内通过
```
用户名：{{.name}}
年龄：{{.age}}
老师：{{.teacher}}
```
模版中`{{}}`是模版取值的固定标记，`.`表示当前变量，渲染后输出的文本效果：
```
用户名：张三 年龄：18 老师：[TC1 TC2 TC3]
```
输出的结果似乎并不太美观，我们试着增加下换行，以及把各个老师的名字单独输出：
```
用户名：{{.name}}

年龄：{{.age}}

{{ range $k,$v:=.teacher }}
老师_{{$k}}：{{$v}}
{{end}}
```
渲染后的效果：
```
用户名：张三

年龄：18

老师_0：TC1

老师_1：TC2

老师_2：TC3
```
这里我们用到了比较简单的循环`{{range}}`(range可以遍历map`{}`、slice`[]`)

继续我们在试试只输出一位老师，比如仅输出TC2:
```
用户名：{{.name}}

年龄：{{.age}}

{{ range $k,$v:=.teacher }}
{{if eq $v "TC2"}}
老师_{{$k}}：{{$v}}
{{end}}
{{end}}
```
渲染后的效果：
```
用户名：张三

年龄：18

老师_1：TC2
```
这里我们使用了简单的判断，判断老师的名字是否是TC2，只有TC2的情况下，才会输出（一般可用于判断告警的状态）。

接着我们尝试把`like`做下处理：
```
用户名：{{.name}}

年龄：{{.age}}

{{ range $k,$v:=.teacher }}
{{if eq $v "TC2"}}
老师_{{$k}}：{{$v}}
{{end}}{{end}}
爱好：{{.like.balls}}
```
输出结果：
```
用户名：张三

年龄：18

老师_1：TC2

爱好：football
```
`{{.like.balls}}`这样的用法是取`like`下面的`balls`的值

如果要取`other`里面的`swim`，我们可以使用`{{index}}`函数，比如：
```
用户名：{{.name}}

年龄：{{.age}}

{{ range $k,$v:=.teacher }}
{{if eq $v "TC2"}}
老师_{{$k}}：{{$v}}
{{end}}{{end}}
爱好：{{index .like.other 0}}
```
输出结果：
```
用户名：张三

年龄：18

老师_1：TC2

爱好：swim
```

备注：原始的json通常都是各种告警系统或者WebHook系统主动发送给PrometheusAlert自定义模板接口，提取json可以去PrometheusAlert的日志中查看
