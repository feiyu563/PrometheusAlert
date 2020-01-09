[← SSL 验证](3-Verify-CN.md) | 代理[(English)](4-Proxy-EN.md) | [超时 →](5-Timeout-CN.md)
***

# 代理

## 描述
当你需要使用代理来发送你的请求时，你可以通过设置环境变量或者通过客户端来设置代理。
`HTTP_PROXY`: 仅对 http 请求有效。
`HTTPS_PROXY`: 仅对 https 请求有效。
`NO_PROXY`: NO_PROXY 中的 ip 或者域名不使用代理。

## 设置

### 通过环境变量设置
你可以设置环境变量 `HTTP_PROXY`, `HTTPS_PROXY` 或者 `NO_PROXY` 。

### 通过客户端设置
```go
// 客户端设置代理优先级比环境变量高
client.SetHttpProxy("http://127.0.0.1:8080")   // 设置 Http 代理
client.GetHttpProxy()                          // 获取 Http 代理.

client.SetHttpsProxy("https://127.0.0.1:8080")   // 设置 Https 代理.
client.GetHttpsProxy()                           // 获取 Https 代理.

client.SetNoProxy("127.0.0.1,localhost")     // 设置代理白名单.
client.GetNoProxy()                          // 获取代理白名单
```

***
[← SSL 验证](3-Verify-CN.md) | 代理[(English)](4-Proxy-EN.md) | [超时 →](5-Timeout-CN.md)
