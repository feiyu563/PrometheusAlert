[← Installation](1-Installation-EN.md) | Client[(中文)](2-Client-CN.md) | [SSL Verify →](3-Verify-EN.md)
***

# Client
You may create multiple different clients simultaneously. Each client can have its own configuration, and each request can be sent by specified client. Use the Default Client if it is not specified. The client can be created by auto-loading of the configuration files, or created and managed manually. Different types of clients require different `Credential`，and different `Signature` algorithms that are selected. You may also customize the client: that is, pass in custom credentials and signatures.

## Client Type

### AccessKey Client
Setup AccessKey through [User Information Management][ak], they have full authority over the account, please keep them safe. Sometimes for security reasons, you cannot hand over a primary account AccessKey with full access to the developer of a project. You may create a sub-account [RAM Sub-account][ram] , grant its [authorization][permissions]，and use the AccessKey of RAM Sub-account to make API calls.
> Sample Code: Create a client with a certification type AccessKey.

```go
client, err := sdk.NewClientWithAccessKey("regionId", "accessKeyId", "accessKeySecret")

```


### STS Client
Create a temporary security client by applying Temporary Security Credentials (TSC) through the Security Token Service (STS).
> Sample Code: Create a client with a certification type StsToken.

```go
client, err := sdk.NewClientWithStsToken("regionId", "subaccessKeyId", "subaccessKeySecret", "stsToken")
```


### RamRoleArn Client
By specifying [RAM Role][RAM Role], the client will be able to automatically request maintenance of STS Token before making a request, and be automatically converted to a time-limited STS client. You may also apply for Token maintenance by yourself before creating `STS Client`.  
> Sample Code: Create a client with a certification type RamRoleArn.

```go
client, err := sdk.NewClientWithRamRoleArn("regionId", "subaccessKeyId", "subaccessKeySecret", "roleArn", "roleSession")
```

If you want to limit the policy([How to make a policy][policy]) of STS Token, you can create a client as following:
```go
client, err := sdk.NewClientWithRamRoleArnAndPolicy("regionId", "subaccessKeyId", "subaccessKeySecret", "roleArn", "roleSession", "policy")
```

### EcsRamRole Client
By specifying the role name, the client will be able to automatically request maintenance of STS Token before making a request, and be automatically converted to a time-limited STS client. You may also apply for Token maintenance by yourself before creating `STS Client`.  
> Sample Code: Create a client with a certification type EcsRamRole.

```go
client, err := NewClientWithEcsRamRole("regionid", "roleName")
```


### Bearer Token Client
If clients with this certification type are required by the Cloud Call Centre (CCC), please apply for Bearer Token maintenance by yourself.
> Sample Code: Create a client with a certification type Bearer Token.

```go
client, err := NewClientWithBearerToken("regionId", "bearerToken")
```


### RsaKeyPair Client
By specifying the public key ID and the private key file, the client will be able to automatically request maintenance of the AccessKey before sending the request, and be automatically converted to a time-limited AccessKey client. Only Japan station is supported. 
> Sample Code: Create a client with a certification type RsaKeyPair.

```go
client, err := NewClientWithRsaKeyPair("regionid", "publicKey", "privateKey", 3600)
```

## Create the client automatically
If no client is created before the request is sent, the client will be created using the default credential provider chain, or the program chain can be customized.

### Default Credential Provider Chain
The default credential provider chain looks for available clients, looking in the following order:

#### 1. Environment Credentials
The program first looks for environment credentials in the environment variable. If the `ALIBABA_CLOUD_ACCESS_KEY_ID` and `ALIBABA_CLOUD_ACCESS_KEY_SECRET` environment variables are defined and are not empty, the program will use them to create the default client. If the client specified by the request is not the default client, the program loads and looks for the client in the configuration file.

#### 2. Credentials File
> If there is `~/.alibabacloud/credentials` default file (Windows shows `C:\Users\USER_NAME\.alibabacloud\credentials`), the program will automatically create clients with the specified type and name. The default file may not exist, but a parse error throws an exception. The client name is case-insensitive, and if the clients have the same name, the latter will override the former. The specified files can also be loaded indefinitely: `AlibabaCloud::load('/data/credentials', 'vfs://AlibabaCloud/credentials', ...);` This configuration file can be shared between different projects and between different tools.  Because it is outside the project and will not be accidentally committed to the version control. Environment variables can be used on Windows to refer to the home directory %UserProfile%. Unix-like systems can use the environment variable $HOME or ~ (tilde). The path to the default file can be modified by defining the `ALIBABA_CLOUD_CREDENTIALS_FILE` environment variable.

```ini
[default]                          # Default client
type = access_key                  # Certification type: access_key
access_key_id = foo                # Key
access_key_secret = bar            # Secret

[client1]                          # Client that is named as `client1`
type = ecs_ram_role                # Certification type: ecs_ram_role
role_name = EcsRamRoleTest         # Role Name

[client2]                          # Client that is named as `client2` 
type = ram_role_arn                # Certification type: ram_role_arn
access_key_id = foo
access_key_secret = bar
role_arn = role_arn
role_session_name = session_name


[client3]                          # Client that is named as `client3`
type = rsa_key_pair                # Certification type: rsa_key_pair
public_key_id = publicKeyId        # Public Key ID
private_key_file = /your/pk.pem    # Private Key file

```

#### 3. Instance RAM Role
If the environment variable `ALIBABA_CLOUD_ECS_METADATA` is defined and not empty, the program will take the value of the environment variable as the role name and request `http://100.100.100.200/latest/meta-data/ram/security-credentials/` to get the temporary Security credentials, then create a default client.

### Custom Credential Provider Chain
You can replace the default order of the program chain by customizing the program chain, or you can write the closure to the provider.
```go
client, err := sdk.NewClientWithProvider("regionId", ProviderInstance, ProviderProfile, ProviderEnv)
```

***
[← Installation](1-Installation-EN.md) | Client[(中文)](2-Client-CN.md) | [SSL Verify →](3-Verify-EN.md)

[ak]: https://usercenter.console.aliyun.com/#/manage/ak
[ram]: https://ram.console.aliyun.com/users
[policy]: https://www.alibabacloud.com/help/doc-detail/28664.htm?spm=a2c63.p38356.a3.3.27a63b01khWgdh
[permissions]: https://ram.console.aliyun.com/permissions
[RAM Role]: https://ram.console.aliyun.com/#/role/list