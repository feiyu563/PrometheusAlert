[← 日志](7-Logger-CN.md) | 并发[(English)](8-Concurrent-EN.md) | [异步调用 →](9-Asynchronous-CN.md)
***

## 并发请求

* 因 Go 语言的并发特性，我们建议您在应用层面控制 SDK 的并发请求。
* 为了方便您的使用，我们也提供了可直接使用的并发调用方式，相关的并发控制由 SDK 内部实现。

### 开启 SDK Client 的并发功能

```go
// 最大并发数
poolSize := 2
// 可缓存的最大请求数
maxTaskQueueSize := 5

// 在创建时开启异步功能
config := sdk.NewConfig()
            .WithEnableAsync(true)
            .WithGoRoutinePoolSize(poolSize)            // 可选，默认5
            .WithMaxTaskQueueSize(maxTaskQueueSize)     // 可选，默认1000
ecsClient, err := ecs.NewClientWithOptions(config)

// 也可以在client初始化后再开启
client.EnableAsync(poolSize, maxTaskQueueSize)
```

***
[← 日志](7-Logger-CN.md) | 并发[(English)](8-Concurrent-EN.md) | [异步调用 →](9-Asynchronous-CN.md)