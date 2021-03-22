## 安装部署PrometheusAlert

部署方式
----

PrometheusAlert可以部署在本地和云平台上，支持windows、linux、公有云、私有云、混合云、容器和kubernetes。你可以根据实际场景或需求，选择相应的方式来部署PrometheusAlert：

- 使用容器部署

```
#clone项目源代码
git clone https://github.com/feiyu563/PrometheusAlert.git

#创建配置文件
mkdir /etc/prometheusalert-center/
cp PrometheusAlert/conf/app-example.conf /etc/prometheusalert-center/

#启动PrometheusAlert并挂载配置文件
docker run -d -p 8080:8080 -v /etc/prometheusalert-center:/app/conf --name prometheusalert-center feiyu563/prometheus-alert:latest

#启动后可使用浏览器打开以下地址查看：http://127.0.0.1:8080
```

- 在linux系统中部署

```
#clone项目源代码
git clone https://github.com/feiyu563/PrometheusAlert.git

#进入程序目录并运行PrometheusAlert
cd PrometheusAlert/example/linux/
./PrometheusAlert #后台运行请执行nohup ./PrometheusAlert &

#启动后可使用浏览器打开以下地址查看：http://127.0.0.1:8080
```

- 在windows系统中运行

```
#clone项目源代码
git clone https://github.com/feiyu563/PrometheusAlert.git

#进入程序目录并双击运行 PrometheusAlert.exe即可
cd PrometheusAlert/example/windows/

#启动后可使用浏览器打开测试地址：http://127.0.0.1:8080
```

- 在kubernetes中运行

```
#Kubernetes中运行可以直接执行以下命令行即可(注意默认的部署模版中未挂载模版数据库文件 db/PrometheusAlertDB.db，为防止模版数据丢失，请自行增加挂载配置 )
kubectl apply -n monitoring -f https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/example/kubernetes/PrometheusAlert-Deployment.yaml

#启动后可使用浏览器打开以下地址查看：http://[YOUR-PrometheusAlert-URL]:8080
```

- 使用helm部署

```
#clone项目源代码
git clone https://github.com/feiyu563/PrometheusAlert.git
cd PrometheusAlert/example/helm

#如需修改配置文件,请更新config中的app.conf
#helm部署模版支持配置Ingress域名，可在values.yaml中进行配置
#配置修改完成后，通过以下命令启动即可(注意默认的部署模版中未挂载模版数据库文件 db/PrometheusAlertDB.db，为防止模版数据丢失，请自行增加挂载配置 )
helm upgrade --install monitor prometheusalert -n monitoring

#启动后可使用浏览器打开以下地址查看: http://[Ingress_url]:[Ingress_port]
```
--------------------------------------------------------------------

配置PrometheusAlert使用mysql作为后端数据存储
----
- PrometheusAlert默认使用sqlite3作为后端自定义模板的存储，这种方式适合于单机部署，满足绝大部分生产场景使用。考虑到部分企业对于服务的高可用要求较高，同时也为了让PrometheusAlert更易于横向扩展，用户可以更改PrometheusAlert的默认存储为mysql。（推荐使用mysql 5.7及以上版本）
- 1.创建数据库
    ```
    CREATE DATABASE prometheusalert CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
    ```
- 2.利用`Navicat`或命令行将`db目录`中的 `prometheusalert.sql` 导入数据库`prometheusalert`
    ```
    use prometheusalert
    source prometheusalert.sql
    ```
- 3.开启PrometheusAlert配置文件中关于mysql的配置 conf/app.conf，数据库名称与上面创建的数据一致

    ```
    #数据库驱动，支持sqlite3，mysql,如使用mysql，请开启db_host,db_user,db_password,db_name的注释
    db_driver=mysql
    db_host=127.0.0.1:3306
    db_user=root
    db_password=root
    db_name=prometheusalert
    ```
- 重启PrometheusAlert，这样即完成配置PrometheusAlert使用mysql数据库作为默认后端存储。
