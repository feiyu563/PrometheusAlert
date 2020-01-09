[← Logger](7-Logger-EN.md) | Concurrent[(中文)](8-Concurrent-CN.md) | [Asynchronous Call →](9-Asynchronous-EN.md)
***

## Concurrent Request

* Due to the concurrency nature of the Go language, we recommend that you control the concurrent requests for the SDK at the application level.
* In order to facilitate your use, we also provide a direct use of concurrent invocation mode, the relevant concurrency control by the SDK internal implementation.

### Open SDK Client's concurrent function.

```go
// Maximum Running Vusers
poolSize := 2
// The maximum number of requests that can be cached
maxTaskQueueSize := 5

// Enable asynchronous functionality at creation time
config := sdk.NewConfig()
            .WithEnableAsync(true)
            .WithGoRoutinePoolSize(poolSize)            // Optional，default:5
            .WithMaxTaskQueueSize(maxTaskQueueSize)     // Optional，default:1000
ecsClient, err := ecs.NewClientWithOptions(config)

// It can also be opened after client is initialized
client.EnableAsync(poolSize, maxTaskQueueSize)
```

***
[← Logger](7-Logger-EN.md) | Concurrent[(中文)](8-Concurrent-CN.md) | [Asynchronous Call →](9-Asynchronous-EN.md)