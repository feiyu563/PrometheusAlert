set GO11MODULE=on
set GO111MODULE=on
set GOPROXY=https://goproxy.io
# go mod init PrometheusAlert
cd ..
go mod vendor
go build