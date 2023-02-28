FROM golang:1.18-alpine3.17 as builder

WORKDIR $GOPATH/src/github.com/feiyu563/PrometheusAlert

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && apk upgrade && \
    apk add --no-cache gcc g++ sqlite-libs make git

ENV GO111MODULE on

ENV GOPROXY https://goproxy.io

COPY . $GOPATH/src/github.com/feiyu563/PrometheusAlert

RUN make build

# -----------------------------------------------------------------------------
FROM alpine:3.17

LABEL maintainer="jikun.zhang"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata && \
	mkdir -p /app/logs && \
    apk update && apk upgrade && apk add --no-cache sqlite-libs curl sqlite

HEALTHCHECK --start-period=10s --interval=20s --timeout=3s --retries=3 \
    CMD curl -fs http://localhost:8080/health || exit 1

WORKDIR /app

COPY --from=builder /go/src/github.com/feiyu563/PrometheusAlert/PrometheusAlert .

COPY conf/app-example.conf conf/app.conf

COPY db db

COPY static static

COPY views views

COPY docker-entrypoint.sh docker-entrypoint.sh

ENTRYPOINT [ "/bin/sh", "/app/docker-entrypoint.sh" ]
