.PHONY: build

build:
	GOOS=linux GOARCH=amd64 go build main.go