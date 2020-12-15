set GO11MODULE=on
set GO111MODULE=on
set GOPROXY=https://goproxy.io
:: go mod init PrometheusAlert
go mod vendor
go build