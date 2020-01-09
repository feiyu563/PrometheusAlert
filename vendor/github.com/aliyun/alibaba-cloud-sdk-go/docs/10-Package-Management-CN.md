[← 异步调用](9-Asynchronous-CN.md) | 包管理[(English)](10-Package-Management-EN.md) | [首页 →](../README-CN.md)
***
## 包管理

Alibaba Cloud SDK for Go 支持两种方式的包管理.

### dep

在 alibaba-cloud-sdk-go 目录下执行以下命令：
```bash
# 当存在 Gopkg.lock 及 Gopkg.toml 时, 该指令会去拉取依赖包并放入 vendor 目录下.
dep ensure
```

### go modules

在 alibaba-cloud-sdk-go 目录下执行以下命令：
```bash
# 当存在 go.mod 及 go.sum 时, 该指令会去拉取依赖包并放入 $GOPATH/pkg/mod 目录下.
go mod tidy
```

***
[← 异步调用](9-Asynchronous-CN.md) | 包管理[(English)](10-Package-Management-EN.md) | [首页 →](../README-CN.md)
