.DEFAULT_GOAL := build

fmt:
	go fmt ./...

lint:
	golint ./...

vet:
	go vet ./...

build:
	go build hello.go

.PHONY: fmt lint vet build
