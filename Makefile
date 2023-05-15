BINARY_NAME=BWINF


default: build

build:
	go build -o bin/$(BINARY_NAME)

run:
	go run main.go

clean:
	rm -rf bin/

test:
	go test -v ./...

# Cross compilation for amd64
linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)-linux-amd64

windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY_NAME)-windows-amd64.exe

darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BINARY_NAME)-darwin-amd64

# Cross compilation for arm64
linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o bin/$(BINARY_NAME)-linux-arm64

windows-arm64:
	GOOS=windows GOARCH=arm64 go build -o bin/$(BINARY_NAME)-windows-arm64.exe

darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o bin/$(BINARY_NAME)-darwin-arm64

# Cross compilation for all
linux: linux-amd64 linux-arm64

windows: windows-amd64 windows-arm64

darwin: darwin-amd64 darwin-arm64


all: linux windows darwin