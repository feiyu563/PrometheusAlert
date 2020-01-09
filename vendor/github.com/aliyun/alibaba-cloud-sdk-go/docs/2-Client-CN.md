[← 安装](1-Installation-CN.md) | 客户端[(English)](2-Client-EN.md) | [SSL 验证 →](3-Verify-CN.md)
***

# 客户端
您可以同时创建多个不同的客户端，每个客户端都可以有独立的配置，每一个请求都可以指定发送的客户端，如果不指定则使用默认客户端。客户端可以通过配置文件自动加载创建，也可以手动创建、管理。不同类型的客户端需要不同的凭证 `Credential`，内部也选取不同的签名算法 `Signature`，您也可以自定义客户端：即传入自定义的凭证和签名。
## 客户端类型

### AccessKey 客户端
通过[用户信息管理][ak]设置AccessKey，它们具有该账户完全的权限，请妥善保管。有时出于安全考虑，您不能把具有完全访问权限的主账户 AccessKey 交于一个项目的开发者使用，您可以[创建RAM子账户][ram]并为子账户[授权][permissions]，使用RAM子用户的 AccessKey 来进行API调用。 

```go
client, err := sdk.NewClientWithAccessKey("regionId", "accessKeyId", "accessKeySecret")

```

### STS 客户端
通过安全令牌服务（Security Token Service，简称 STS），申请临时安全凭证（Temporary Security Credentials，简称 TSC），创建临时安全客户端。

```go
client, err := sdk.NewClientWithStsToken("regionId", "subaccessKeyId", "subaccessKeySecret", "stsToken")
```


### RamRoleArn 客户端
通过指定[RAM角色][RAM Role]，让客户端在发起请求前自动申请维护 STS Token，自动转变为一个有时限性的STS客户端。您也可以自行申请维护 STS Token，再创建 `STS客户端`。  
> 示例代码：创建一个 RamRoleArn 方式认证的客户端。

```go
client, err := sdk.NewClientWithRamRoleArn("regionId", "subaccessKeyId", "subaccessKeySecret", "roleArn", "roleSession")
```

如果你想限制生成的 STS Token 的权限([构建Policy][policy]), 你可以使用如下方式创建客户端:
```go
client, err := sdk.NewClientWithRamRoleArnAndPolicy("regionId", "subaccessKeyId", "subaccessKeySecret", "roleArn", "roleSession", "policy")
```


### EcsRamRole 客户端
通过指定角色名称，让客户端在发起请求前自动申请维护 STS Token，自动转变为一个有时限性的STS客户端。您也可以自行申请维护 STS Token，再创建 `STS客户端`。  
> 示例代码：创建一个 EcsRamRole 方式认证的客户端。

```go
client, err := NewClientWithEcsRamRole("regionid", "roleName")
```


### Bearer Token 客户端
如呼叫中心(CCC)需用此类认证方式的客户端，请自行申请维护 Bearer Token。  
> 示例代码：创建一个 Bearer Token 方式认证的客户端。

```go
client, err := NewClientWithBearerToken("regionId", "bearerToken")
```


### RsaKeyPair 客户端
通过指定公钥ID和私钥文件，让客户端在发起请求前自动申请维护 AccessKey，自动转变成为一个有时限性的AccessKey客户端，仅支持日本站。  
> 示例代码：创建一个 RsaKeyPair 方式认证的客户端。


```go
client, err := NewClientWithRsaKeyPair("regionid", "publicKey", "privateKey", 3600)
```

## 自动创建客户端
在发送请求前，如果没有创建任何客户端，将使用默认凭证提供程序链创建客户端，也可以自定义程序链。

### 默认凭证提供程序链
默认凭证提供程序链查找可用的客户端，寻找顺序如下：

#### 1. 环境凭证
程序首先会在环境变量里寻找环境凭证，如果定义了 `ALIBABA_CLOUD_ACCESS_KEY_ID`  和 `ALIBABA_CLOUD_ACCESS_KEY_SECRET` 环境变量且不为空，程序将使用他们创建默认客户端。如果请求指定的客户端不是默认客户端，程序会在配置文件中加载和寻找客户端。

#### 2. 配置文件
> 如果用户主目录存在默认文件 `~/.alibabacloud/credentials` （Windows 为 `C:\Users\USER_NAME\.alibabacloud\credentials`），程序会自动创建指定类型和名称的客户端。默认文件可以不存在，但解析错误会抛出异常。  客户端名称不分大小写，若客户端同名，后者会覆盖前者。也可以手动加载指定文件： `AlibabaCloud::load('/data/credentials', 'vfs://AlibabaCloud/credentials', ...);` 不同的项目、工具之间可以共用这个配置文件，因为超出项目之外，也不会被意外提交到版本控制。Windows 上可以使用环境变量引用到主目录 %UserProfile%。类 Unix 的系统可以使用环境变量 $HOME 或 ~ (tilde)。 可以通过定义 `ALIBABA_CLOUD_CREDENTIALS_FILE` 环境变量修改默认文件的路径。

```ini
[default]                          # 默认客户端
type = access_key                  # 认证方式为 access_key
access_key_id = foo                # Key
access_key_secret = bar            # Secret

[client1]                          # 命名为 `client1` 的客户端
type = ecs_ram_role                # 认证方式为 ecs_ram_role
role_name = EcsRamRoleTest         # Role Name

[client2]                          # 命名为 `client2` 的客户端
type = ram_role_arn                # 认证方式为 ram_role_arn
access_key_id = foo
access_key_secret = bar
role_arn = role_arn
role_session_name = session_name

[client3]                          # 命名为 `client3` 的客户端
type = rsa_key_pair                # 认证方式为 rsa_key_pair
public_key_id = publicKeyId        # Public Key ID
private_key_file = /your/pk.pem    # Private Key 文件

```

#### 3. 实例 RAM 角色
如果定义了环境变量 `ALIBABA_CLOUD_ECS_METADATA` 且不为空，程序会将该环境变量的值作为角色名称，请求 `http://100.100.100.200/latest/meta-data/ram/security-credentials/` 获取临时安全凭证，再创建一个默认客户端。

### 自定义凭证提供程序链
可通过自定义程序链代替默认程序链的寻找顺序，也可以自行编写闭包传入提供者。
```go
client, err := sdk.NewClientWithProvider("regionId", ProviderInstance, ProviderProfile, ProviderEnv)
```

***
[← 安装](1-Installation-CN.md) | 客户端[(English)](2-Client-EN.md) | [SSL 验证 →](3-Verify-CN.md)

[ak]: https://usercenter.console.aliyun.com/#/manage/ak
[ram]: https://ram.console.aliyun.com/users
[permissions]: https://ram.console.aliyun.com/permissions
[RAM Role]: https://ram.console.aliyun.com/#/role/list