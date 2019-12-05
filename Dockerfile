FROM centos:7
MAINTAINER jikun.zhang <jikun.zhang>
ADD example/linux /app
WORKDIR /app
RUN chmod -R 755 /app
CMD ["/app/PrometheusAlert"]
