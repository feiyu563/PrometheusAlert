FROM golang:1.20.6-alpine3.18 as builder

WORKDIR $GOPATH/src/github.com/feiyu563/PrometheusAlert

RUN apk update && \
    apk add --no-cache gcc g++ sqlite-libs make git

ENV GO111MODULE on

ENV GOPROXY https://goproxy.io

COPY . $GOPATH/src/github.com/feiyu563/PrometheusAlert

RUN make build

# -----------------------------------------------------------------------------
FROM alpine:3.18

LABEL maintainer="jikun.zhang"

RUN apk update && \
    apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata && \
	mkdir -p /app/logs && \
    apk add --no-cache sqlite-libs curl sqlite

HEALTHCHECK --start-period=10s --interval=20s --timeout=3s --retries=3 \
    CMD curl -fs http://localhost:8080/health || exit 1

WORKDIR /app

COPY --from=builder /go/src/github.com/feiyu563/PrometheusAlert/PrometheusAlert .

COPY db/PrometheusAlertDB.db /opt/PrometheusAlertDB.db

COPY conf/app-example.conf conf/app.conf

COPY db db

COPY static static

COPY views views

COPY docker-entrypoint.sh docker-entrypoint.sh

ENTRYPOINT [ "/bin/sh", "/app/docker-entrypoint.sh" ]
