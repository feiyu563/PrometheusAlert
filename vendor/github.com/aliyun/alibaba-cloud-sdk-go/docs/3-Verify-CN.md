[← 客户端](2-Client-CN.md) | SSL 验证[(English)](3-Verify-EN.md) | [代理 →](4-Proxy-CN.md)
***

# SSL 验证

## 摘要
请求时验证SSL证书行为。
- 设置成 `false` 禁用证书验证，(这是不安全的，请设置证书！)。
- 设置成 `true` 启用SSL证书验证，默认使用操作系统提供的CA包。

## 默认值
- `false` 

## 设置
### 通过请求设置
```go
// 设置请求 HTTPSInsecure (只影响当前)
request.SetHTTPSInsecure(true)                           // 设置请求 HTTPSInsecure 为 true
isInsecure := request.GetHTTPSInsecure()                 // 获取请求 HTTPSInsecure
```

### 通过客户端设置
> 当请求未设置时，客户端设置才能生效.

```go
// 设置客户端 HTTPSInsecure (用于客户端发送的所有请求)。
client.SetHTTPSInsecure(true)                         // 设置客户端 HTTPSInsecure 为 true
isInsecure := client.GetHTTPSInsecure()               // 获取客户端 HTTPSInsecure
```

***
[← 客户端](2-Client-CN.md) | SSL 验证[(English)](3-Verify-EN.md) | [代理 →](4-Proxy-CN.md)
