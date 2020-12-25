FROM golang:1.14-alpine3.12 as builder

WORKDIR $GOPATH/src/github.com/feiyu563/PrometheusAlert

RUN apk update && apk upgrade && apk add --no-cache gcc g++ sqlite-libs

ENV GO111MODULE on

ENV GOPROXY https://goproxy.io

COPY . $GOPATH/src/github.com/feiyu563/PrometheusAlert

RUN go mod vendor && go build

# -----------------------------------------------------------------------------

FROM alpine:3.12

LABEL maintainer="jikun.zhang"

RUN apk update && apk upgrade && apk add --no-cache sqlite-libs

WORKDIR /app

COPY --from=builder /go/src/github.com/feiyu563/PrometheusAlert/PrometheusAlert .

COPY db/PrometheusAlertDB.db /opt/PrometheusAlertDB.db

COPY conf/app-example.conf conf/app.conf

COPY db db

COPY logs logs

COPY static static

COPY views views

#ENTRYPOINT [ "/app/PrometheusAlert" ]

CMD if [ ! -f /app/db/PrometheusAlertDB.db ];then cp /opt/PrometheusAlertDB.db /app/db/PrometheusAlertDB.db;echo 'init ok!';else echo 'pass!';fi && /app/PrometheusAlert