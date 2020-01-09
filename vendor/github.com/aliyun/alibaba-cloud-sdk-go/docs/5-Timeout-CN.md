[← 代理](4-Proxy-CN.md) | 超时[(English)](5-Timeout-EN.md) | [调试 →](6-Debug-CN.md)
***

# 超时

## 描述
如果你想限制请求花费的时间，你可以通过请求或者客户端设置 `ConnectTimeout` 和 `ReadTimeout`。

## 默认值
- `defaultConnectTimeout`: 5 秒
- `defaultReadTimeout`: 10 秒

## 设置
### 通过请求设置
```go
// 设置请求超时（仅对当前请求有效）
request.SetReadTimeout(10 * time.Second)             // 设置请求读超时为10秒
readTimeout := request.GetReadTimeout()              // 获取请求读超时
request.SetConnectTimeout(5 * time.Second)           // 设置请求连接超时为5秒
connectTimeout := request.GetConnectTimeout()        // 获取请求连接超时
```

### 通过客户端设置
> 当请求未设置超时时，客户端设置的超时才会生效。

```go
// 设置客户端超时（对所有通过该客户端发送的请求生效）
client.SetReadTimeout(10 * time.Second)             // 设置客户端读超时为10秒
readTimeout := client.GetReadTimeout()              // 获取客户端读超时
client.SetConnectTimeout(5 * time.Second)           // 设置客户端连接超时为5秒
connectTimeout := client.GetConnectTimeout()        // 获取客户端连接超时
```

***
[← 代理](4-Proxy-CN.md) | 超时[(English)](5-Timeout-EN.md) | [调试 →](6-Debug-CN.md)
