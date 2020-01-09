[← Debug](6-Debug-EN.md) | Logger[(中文)](7-Logger-CN.md) | [Concurrent →](8-Concurrent-EN.md)
***


# Logger

## Description

The logger is mainly used to provide support for auditing, to record each call, similar to the server's access log.

## Using

### Init logger

If you want to use the log function, you need to initialize a log object first. You can set the log level, log template, log output path and channel when initializing the log object.
```go
// level: default value is info
// channel: default value is AlibabaCloud
// file: should be an object that implements the io.writer interface
// templete: logger template, If not entered, the default value is `{time} {channel}: "{method} {uri} HTTP/{version}" {code} {cost} {hostname}`
client.SetLogger("level", "channel", file, templete)      // Set the client's log. When you call this method, the log function is enabled by default.
```

### Related operations

```go
logger := client.GetLogger()    // Get client logger 
client.OpenLogger()            // Open logger, if clien logger is not exist, there will create a default logger for client
client.CloseLogger()           // Close logger
client.GetLoggerMsg()          // Get last logger message，if clien logger is not exist, there will create a default logger for client
client.SetTemplate(templete)   // Set client logger template，if clien logger is not exist, there will create a default logger for client
client.GetTemplate()           // Get client logger template，if clien logger is not exist, there will create a default logger for client
```

### Variables

| Variables |  Description  |
|----------|-------------|
| {channel}     | name of the log |
| {host}     | Host of the request |
| {ts}     | GMT中的 ISO 8601日期 |
| {method}     | Method of the request |
| {uri}     | URI of the request |
| {version}     | Protocol version |
| {target}     | Request target of the request (path + query) |
| {hostname}     | Hostname of the machine that sent the request |
| {code}     | Status code of the response (if available) |
| {error}     | Any error messages (if available) |
| {req_headers}     | Request headers |
| {res_headers}     | Response headers |
| {pid}     | PID |
| {cost}     | Cost Time |
| {start_time}     | start Time |
| {res_body}  | Response body |

***
[← Debug](6-Debug-EN.md) | Logger[(中文)](7-Logger-CN.md) | [Concurrent →](8-Concurrent-EN.md)