# Building stage
FROM golang:1.13-alpine3.10  AS builder

WORKDIR /build/src/PrometheusAlert
RUN adduser -u 10001 -D app-runner

ENV GO111MODULE off
ENV GOPATH /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o PrometheusAlert  main.go

# Production stage
FROM alpine:3.10 AS final

WORKDIR /app
MAINTAINER jikun.zhang <jikun.zhang>

COPY --from=builder /build/src/PrometheusAlert/example/linux /app
COPY --from=builder /build/src/PrometheusAlert/example/linux/db/PrometheusAlertDB.db /opt/PrometheusAlertDB.db
COPY --from=builder /build/src/PrometheusAlert/PrometheusAlert /app
COPY --from=builder /build/src/PrometheusAlert/conf /app/conf

RUN adduser -u 10001 -D app-runner
RUN chmod -R 755 /app

CMD if [ ! -f /app/db/PrometheusAlertDB.db ];then cp /opt/PrometheusAlertDB.db /app/db/PrometheusAlertDB.db;echo 'init ok!';else echo 'pass!';fi

ENTRYPOINT ["/app/PrometheusAlert"]