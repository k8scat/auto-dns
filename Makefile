.PHONY: build
build:
	GOOS=linux GOARCH=arm64 go build -o bin/rpi-auto-dns main.go
