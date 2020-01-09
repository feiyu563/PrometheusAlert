[← Asynchronous Call](9-Asynchronous-EN.md) | Package Management[(中文)](10-Package-Management-CN.md) | [Home →](../README.md)
***
## Package Management

Alibaba Cloud SDK for Go supports two ways for package management.

### dep

Execute the following command in the alibaba-cloud-sdk-go directory：
```bash
# When gopkg.lock and gopkg.toml exist, this instruction will pull the dependency package and put it into the vendor directory.
dep ensure 
```

### go modules

Execute the following command in the alibaba-cloud-sdk-go directory：
```bash
# When go.mod and go.sum exist, the command will pull the dependent package and put it into the $GOPATH/pkg/mod directory.
go mod tidy
```

***
[← Asynchronous Call](9-Asynchronous-EN.md) | Package Management[(中文)](10-Package-Management-CN.md) | [Home →](../README.md)
 