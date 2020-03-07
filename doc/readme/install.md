PrometheusAlert全家桶部署说明

-----------------

部署方式
----

PrometheusAlert可以部署在本地和云平台上，支持windows、linux、公有云、私有云、混合云、容器和kubernetes。你可以根据实际场景或需求，选择相应的方式来部署PrometheusAlert：

 - 使用容器部署
```
git clone https://github.com/feiyu563/PrometheusAlert.git
mkdir /etc/prometheusalert-center/
cp PrometheusAlert/conf/app.conf /etc/prometheusalert-center/
docker run -d -p 8080:8080 -v /etc/prometheusalert-center:/app/conf --name prometheusalert-center feiyu563/prometheus-alert:latest
```
 - 在linux系统中部署
```
git clone https://github.com/feiyu563/PrometheusAlert.git
cd PrometheusAlert/example/linux/
./PrometheusAlert #后台运行请执行nohup ./PrometheusAlert &
```
- 在windows系统中运行
```
git clone https://github.com/feiyu563/PrometheusAlert.git
cd PrometheusAlert/example/windows/
双击运行 PrometheusAlert.exe即可
```
- 在kubernetes中运行
```
kubectl app -n monitoring -f https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/example/kubernetes/PrometheusAlert-Deployment.yaml
```
- 使用helm部署
```
git clone https://github.com/feiyu563/PrometheusAlert.git
cd PrometheusAlert/example/helm/prometheusalert
#如需修改配置文件,请更新config中的app.conf
helm install -n monitoring .
```
```
启动后可使用浏览器打开测试地址: http://127.0.0.1:8080
```