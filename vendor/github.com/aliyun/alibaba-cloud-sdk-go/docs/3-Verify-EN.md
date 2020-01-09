[← Client](2-Client-EN.md) | SSL Verify[(中文)](3-Verify-CN.md) | [Proxy →](4-Proxy-EN.md)
***

# SSL Verify

## Summary
Describes the SSL certificate verification behavior of a request.
- Set `false` to disable certificate validation, (This is not safe, please set certificates! )
- Set to `true` to enable SSL certificate verification and use the default CA bundle provided by operating system.

## Default
- `false` 

## Setting
### Setting on Request
```go
// Set request HTTPSInsecure(Only the request is effected.)
request.SetHTTPSInsecure(true)                           // Set request HTTPSInsecure to true.
isInsecure := request.GetHTTPSInsecure()                 // Get request HTTPSInsecure.
```

### Setting on Client
> When the request is not set, the client settings are used.

```go
// Set client HTTPSInsecure(For all requests which is sent by the client.)
client.SetHTTPSInsecure(true)                         // Set client HTTPSInsecure to true.
isInsecure := client.GetHTTPSInsecure()               // Get client HTTPSInsecure.
```

***
[← Client](2-Client-EN.md) | SSL Verify[(中文)](3-Verify-CN.md) | [Proxy →](4-Proxy-EN.md)
