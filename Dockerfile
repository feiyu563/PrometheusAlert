FROM centos:7
MAINTAINER jikun.zhang <jikun.zhang>
ADD example/linux /app
ADD example/linux/db/PrometheusAlertDB.db /opt/PrometheusAlertDB.db
WORKDIR /app
RUN chmod -R 755 /app
CMD if [ ! -f /app/db/PrometheusAlertDB.db ];then cp /opt/PrometheusAlertDB.db /app/db/PrometheusAlertDB.db;echo 'init ok!';else echo 'pass!';fi && /app/PrometheusAlert
