[← Concurrent](8-Concurrent-EN.md) | Asynchronous Call[(中文)](9-Asynchronous-CN.md) | [Package Management →](10-Package-Management-EN.md)
***
## Asynchronous Call

### Make an asynchronous call
Alibaba Cloud Go SDK supports asynchronous calls in two ways：

1. Using channel as return values
    ```go
    responseChannel, errChannel := client.FooWithChan(request)

    // this will block
    response := <-responseChannel
    err = <-errChannel
    ```

2. Use callback to control the callback

    ```go
    blocker := client.FooWithCallback(request, func(response *FooResponse, err error) {
        // handle the response and err
    })

    // blocker which is type of (chan int)，is used to control synchronization，when returning 1 means success，and returning 0 means failure.
    // When <-blocker returns failure，err also will be handled by afferent callback.
    result := <-blocker
    ```

***
[← Concurrent](8-Concurrent-EN.md) | Asynchronous Call[(中文)](9-Asynchronous-CN.md) | [Package Management →](10-Package-Management-EN.md)