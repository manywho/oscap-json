.PHONY: build

VERSION = 0.1.0

build:
	go build -ldflags "-X main.version=$(VERSION) -X main.commitHash=$$(git rev-parse --short HEAD)"

release:
	GOARCH=amd64 GOOS=linux go build -ldflags "-X main.version=$(VERSION) -X main.commitHash=$$(git rev-parse --short HEAD)" -o oscap-json-linux-amd64
	GOARCH=386 GOOS=linux go build -ldflags "-X main.version=$(VERSION) -X main.commitHash=$$(git rev-parse --short HEAD)" -o oscap-json-linux-386