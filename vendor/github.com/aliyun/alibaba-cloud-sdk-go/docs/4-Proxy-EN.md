[← SSL Verify](3-Verify-EN.md) | Proxy[(中文)](4-Proxy-CN.md) | [Timeout →](5-Timeout-EN.md)
***

# Proxy

## Description
When you need to use proxy to send your request, you can set environment variables or you can set them by client:
`HTTP_PROXY`: Only the HTTP request to take effect.
`HTTPS_PROXY`: Only the HTTPS request to take effect.
`NO_PROXY`: The Ips or domains in it will not use proxy.

## Setting

### Setting by environment variables
You can set environment variables `HTTP_PROXY`, `HTTPS_PROXY` or `NO_PROXY`

### Setting by client
```go
// client proxy has a high priority than environment variables.
client.SetHttpProxy("http://127.0.0.1:8080")   // Set Http Proxy.
client.GetHttpProxy()                          // Get Http Proxy.

client.SetHttpsProxy("https://127.0.0.1:8080")   // Set Https Proxy.
client.GetHttpsProxy()                           // Get Https Proxy.

client.SetNoProxy("127.0.0.1,localhost")     // Set No Proxy.
client.GetNoProxy()                          // Get No Proxy.
```

***
[← SSL Verify](3-Verify-EN.md) | Proxy[(中文)](4-Proxy-CN.md) | [Timeout →](5-Timeout-EN.md)
