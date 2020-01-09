[← 调试](6-Debug-CN.md) | 日志[(English)](7-Logger-EN.md) | [并发 →](8-Concurrent-CN.md)
***

# 日志

## 描述

logger 主要用于提供支持审计的能力，用于记录每次的调用情况，类似服务端的 access log。

## 使用

### 初始化日志

如果您想要使用日志功能，您需要先初始化一个日志对象，您可以在初始化日志对象的时候设置日志等级，日志模版， 日志的输出路径以及 channel。
```go
// level: 默认为 info
// channel: 默认为 AlibabaCloud
// file: 一个实现了 io.writer 接口的对象
// templete: 日志的模板， 若不输入，则默认为 `{time} {channel}: "{method} {uri} HTTP/{version}" {code} {cost} {hostname}`
client.SetLogger("level", "channel", file, templete)      // 设置客户端的日志， 当您调用该方法，默认为您开启日志功能
```

### 相关操作

```go
logger := client.GetLogger()    // 获取客户端的 logger 
client.OpenLogger()            // 开启日志功能，若此时客户端的 logger 不存在， 则创建一个配置一个默认的 logger
client.CloseLogger()           // 关闭日志功能
client.GetLoggerMsg()          // 获取上一条日志信息，若此时客户端的 logger 不存在， 则创建一个配置一个默认的 logger 
client.SetTemplate(templete)   // 设置日志模板，若此时客户端的 logger 不存在， 则创建一个配置一个默认的 logger
client.GetTemplate()           // 获取当前的日志模板，若此时客户端的 logger 不存在， 则创建一个配置一个默认的 logger
```

### 变量

|    变量    |   描述    |
|----------|-------------|
| {channel}     | 日志的对象 |
| {host}     | 请求主机 |
| {ts}     | GMT中的 ISO 8601日期 |
| {method}     | 请求方法 |
| {uri}     | 请求的URI |
| {version}     | 协议版本 |
| {target}     | 请求目标 (path + query) |
| {hostname}     | 发送请求的计算机的主机名 |
| {code}     | 响应的状态代码（如果可用） |
| {error}     | 任何错误消息（如果有） |
| {req_headers}     | 请求头 |
| {res_headers}     | 响应头 |
| {pid}     | PID |
| {cost}     | 耗时 |
| {start_time}  | 开始时间 |
| {res_body}  | 响应主体 |

***
[← 调试](6-Debug-CN.md) | 日志[(English)](7-Logger-EN.md) | [并发 →](8-Concurrent-CN.md)
