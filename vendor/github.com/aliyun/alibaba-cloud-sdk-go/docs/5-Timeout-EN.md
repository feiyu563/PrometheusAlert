[← Proxy](4-Proxy-EN.md) | Timeout[(中文)](5-Timeout-CN.md) | [Debug →](6-Debug-EN.md)
***

# Timeout

## Description
When you want to limit the time of request costing, you can set `ConnectTimeout` and `ReadTimeout` by request or client:

## Default
- `defaultConnectTimeout`: 5 * time.Second
- `defaultReadTimeout`: 10 * time.Second

## Setting
### Setting on Request
```go
// Set request Timeout(Only the request is effected.)
request.SetReadTimeout(10 * time.Second)             // Set request ReadTimeout to 10 second.
readTimeout := request.GetReadTimeout()              // Get request ReadTimeout.
request.SetConnectTimeout(5 * time.Second)           // Set request ConnectTimeout to 5 second.
connectTimeout := request.GetConnectTimeout()        // Get request ConnectTimeout.
```

### Setting on Client
> When the request is not set, the client settings are used.

```go
// Set client Timeout(For all requests which is sent by the client.)
client.SetReadTimeout(10 * time.Second)              // Set client ReadTimeout to 10 second.
readTimeout := client.GetReadTimeout()               // Get client ReadTimeout.
client.SetConnectTimeout(5 * time.Second)            // Set client ConnectTimeout to 5 second.
connectTimeout := client.GetConnectTimeout()         // Get client ConnectTimeout.
```

***
[← Proxy](4-Proxy-EN.md) | Timeout[(中文)](5-Timeout-CN.md) | [Debug →](6-Debug-EN.md)
