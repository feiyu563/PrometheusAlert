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
cp PrometheusAlert/conf/app.conf /etc/prometheusalert-center/

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
kubectl app -n monitoring -f https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/example/kubernetes/PrometheusAlert-Deployment.yaml

#启动后可使用浏览器打开以下地址查看：http://[YOUR-PrometheusAlert-URL]:8080
```

- 使用helm部署

```
#clone项目源代码
git clone https://github.com/feiyu563/PrometheusAlert.git
cd PrometheusAlert/example/helm/prometheusalert

#如需修改配置文件,请更新config中的app.conf
#helm部署模版支持配置Ingress域名，可在values.yaml中进行配置
#配置修改完成后，通过以下命令启动即可(注意默认的部署模版中未挂载模版数据库文件 db/PrometheusAlertDB.db，为防止模版数据丢失，请自行增加挂载配置 )
helm install -n monitoring .

#启动后可使用浏览器打开以下地址查看: http://[Ingress_url]:[Ingress_port]
```