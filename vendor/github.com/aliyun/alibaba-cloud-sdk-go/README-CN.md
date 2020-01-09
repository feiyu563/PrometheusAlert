[English](./README.md) | 简体中文

<p align="center">
<a href=" https://www.alibabacloud.com"><img src="https://aliyunsdk-pages.alicdn.com/icons/Aliyun.svg"></a>
</p>

<h1 align="center">Alibaba Cloud SDK for Go</h1>

<p align="center">
<a href="https://badge.fury.io/gh/aliyun%2Falibaba-cloud-sdk-go"><img src="https://badge.fury.io/gh/aliyun%2Falibaba-cloud-sdk-go.svg" alt="Latest Stable Version"></a>
<a href="https://app.fossa.io/projects/git%2Bgithub.com%2Faliyun%2Falibaba-cloud-sdk-go?ref=badge_shield"><img src="https://app.fossa.io/api/projects/git%2Bgithub.com%2Faliyun%2Falibaba-cloud-sdk-go.svg?type=shield" alt="License"></a>
<br/>
<a href="https://codecov.io/gh/aliyun/alibaba-cloud-sdk-go"><img src="https://codecov.io/gh/aliyun/alibaba-cloud-sdk-go/branch/master/graph/badge.svg" alt="codecov"></a>
<a href="https://travis-ci.org/aliyun/alibaba-cloud-sdk-go"><img src="https://travis-ci.org/aliyun/alibaba-cloud-sdk-go.svg?branch=master" alt="Travis Build Status"></a>
<a href="https://ci.appveyor.com/project/aliyun/alibaba-cloud-sdk-go/branch/master"><img src="https://ci.appveyor.com/api/projects/status/gn17u48i53ktblfp/branch/master?svg=true" alt="Appveyor Build Status"></a>
<a href="https://app.codacy.com/app/aliyun/alibaba-cloud-sdk-go?utm_source=github.com&utm_medium=referral&utm_content=aliyun/alibaba-cloud-sdk-go&utm_campaign=Badge_Grade_Dashboard"><img src="https://api.codacy.com/project/badge/Grade/291a39e242364b04ad442f0cce0e30d5" alt="Codacy Badge"></a>
<a href="https://goreportcard.com/report/github.com/aliyun/alibaba-cloud-sdk-go"><img src="https://goreportcard.com/badge/github.com/aliyun/alibaba-cloud-sdk-go" alt="Go Report Card"></a>
</p>

欢迎使用 Alibaba Cloud SDK for Go。Alibaba Cloud SDK for Go 让您不用复杂编程即可访问云服务器、云监控等多个阿里云服务。
这里向您介绍如何获取 [Alibaba Cloud SDK for Go][SDK] 并开始调用。

## 在线示例
[API Explorer][open-api] 提供在线调用阿里云产品，并动态生成 SDK 代码和快速检索接口等能力，能显著降低使用云 API 的难度。


## 环境要求
- 您的系统需要达到 [环境要求][Requirements], 例如，安装了不低于 1.10.x 版本的 Go 环境。

## 安装
使用 `go get` 下载安装 SDK

```sh
$ go get -u github.com/aliyun/alibaba-cloud-sdk-go/sdk
```

如果您使用了 glide 管理依赖，您也可以使用 glide 来安装 Alibaba Cloud SDK for Go

```sh
$ glide get github.com/aliyun/alibaba-cloud-sdk-go
```

另外，Alibaba Cloud SDK for Go 也会发布在 https://develop.aliyun.com/tools/sdk#/go 这个地址。

## 快速使用
在您开始之前，您需要注册阿里云帐户并获取您的[凭证](https://usercenter.console.aliyun.com/#/manage/ak)。

### 创建客户端
```go
package main

import "github.com/aliyun/alibaba-cloud-sdk-go/sdk"

func main() {

	client, err := sdk.NewClientWithAccessKey("REGION_ID", "ACCESS_KEY_ID", "ACCESS_KEY_SECRET")
	if err != nil {
		// Handle exceptions
		panic(err)
	}
}
```

### ROA 请求
```go
package main

import "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

func main() {
	request := requests.NewCommonRequest()        // 构造一个公共请求
	request.Method = "GET"                        // 设置请求方式
	request.Product = "CS"                        // 指定产品
	request.Domain = "cs.aliyuncs.com"            // 指定域名则不会寻址，如认证方式为 Bearer Token 的服务则需要指定
	request.Version = "2015-12-15"                // 指定产品版本
	request.PathPattern = "/clusters/[ClusterId]" // 指定ROA风格路径规则
	request.ApiName = "DescribeCluster"           // 指定接口名
	request.QueryParams["ClusterId"] = "123456"   // 设置参数值
	request.QueryParams["RegionId"] = "region_id" // 指定请求的区域，不指定则使用客户端区域、默认区域
	request.TransToAcsRequest()                   // 把公共请求转化为acs请求
}
```

### RPC 请求
```go
package main

import "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

func main() {
	request := requests.NewCommonRequest()                // 构造一个公共请求
	request.Method = "POST"                               // 设置请求方式
	request.Product = "Ecs"                               // 指定产品
	request.Domain = "ecs.aliyuncs.com"                   // 指定域名则不会寻址，如认证方式为 Bearer Token 的服务则需要指定
	request.Version = "2014-05-26"                        // 指定产品版本
	request.ApiName = "CreateInstance"                    // 指定接口名
	request.QueryParams["InstanceType"] = "ecs.g5.large"  // 设置参数值
	request.QueryParams["RegionId"] = "region_id"         // 指定请求的区域，不指定则使用客户端区域、默认区域
	request.TransToAcsRequest()                           // 把公共请求转化为acs请求
}
```


## 文档
* [Requirements](docs/0-Requirements-CN.md)
* [Installation](docs/1-Installation-CN.md)
* [Client](docs/2-Client-CN.md)
* [SSL Verify](docs/3-Verify-CN.md)
* [Proxy](docs/4-Proxy-CN.md)
* [Timeout](docs/5-Timeout-CN.md)
* [Debug](docs/6-Debug-CN.md)
* [Logger](docs/7-Logger-CN.md)
* [Concurrent](docs/8-Concurrent-CN.md)
* [Asynchronous Call](docs/9-Asynchronous-CN.md)
* [Package Management](docs/10-Package-Management-CN.md)


## 问题
[提交 Issue][issue] 不符合指南的问题可能会立即关闭。


## 贡献
提交 Pull Request 之前请阅读[贡献指南](CONTRIBUTING.md)。

## 相关
* [阿里云服务 Regions & Endpoints][endpoints]
* [OpenAPI Explorer][open-api]
* [Go][go]
* [最新发行版本][latest-release]


## 许可证
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Faliyun%2Falibaba-cloud-sdk-go.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Faliyun%2Falibaba-cloud-sdk-go?ref=badge_large)

[SDK]: https://github.com/aliyun/alibaba-cloud-sdk-go
[apache]: http://www.apache.org/licenses/LICENSE-2.0
[issue]: https://github.com/aliyun/alibaba-cloud-sdk-go/issues/new
[open-api]: https://api.aliyun.com/
[latest-release]: https://github.com/aliyun/alibaba-cloud-sdk-go/releases
[go]: https://golang.org/dl/
[endpoints]: https://developer.aliyun.com/endpoints
[Requirements]: docs/0-Requirements-CN.md
