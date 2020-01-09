[← 并发](8-Concurrent-CN.md) | 异步调用[(English)](9-Asynchronous-EN.md) | [包管理 →](10-Package-Management-CN.md)
***
## 异步调用

### 发起异步调用
Alibaba Cloud SDK for Go 支持两种方式的异步调用：

1. 使用channel作为返回值
    ```go
    responseChannel, errChannel := client.FooWithChan(request)

    // this will block
    response := <-responseChannel
    err = <-errChannel
    ```

2. 使用 callback 控制回调

    ```go
    blocker := client.FooWithCallback(request, func(response *FooResponse, err error) {
        // handle the response and err
    })

    // blocker 为(chan int)，用于控制同步，返回1为成功，0为失败
    // 在<-blocker返回失败时，err依然会被传入的callback处理
    result := <-blocker
    ```

***
[← 并发](8-Concurrent-CN.md) | 异步调用[(English)](9-Asynchronous-EN.md) | [包管理 →](10-Package-Management-CN.md)