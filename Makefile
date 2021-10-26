.PHONY: build
build:
	go build cmd/srv-verification-api/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: run
run:
	go run cmd/srv-verification-api/main.go
