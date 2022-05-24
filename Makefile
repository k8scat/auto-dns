.PHONY: build-all
build-all: build-linux-arm64 build-linux-amd64

.PHONY: build-linux-arm64
build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -trimpath -o bin/rpi-auto-dns-linux-arm64 main.go

.PHONY: build-linux-amd64
build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -trimpath -o bin/rpi-auto-dns-linux-amd64 main.go
