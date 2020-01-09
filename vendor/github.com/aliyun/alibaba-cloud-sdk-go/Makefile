
all:

fmt:
	go fmt ./sdk ./integration ./services/...

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic ./sdk/...
	go tool cover -html=coverage.txt -o coverage.html
