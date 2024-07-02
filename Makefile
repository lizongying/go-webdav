.PHONY: all

all: tidy webdav_server


bin := ./releases/webdav_server
path := ./cmd/webdav_server

tidy:
	go mod tidy

webdav_server:
	go vet $(path)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o $(bin)_linux_amd64 $(path)
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-s -w" -o $(bin)_linux_arm64 $(path)
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o $(bin)_darwin_amd64 $(path)
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-s -w" -o $(bin)_darwin_arm64 $(path)
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o $(bin)_windows_amd64.exe $(path)
