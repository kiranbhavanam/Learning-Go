.DEFAULT_GOAL := build
.PHONY:fmt vet build
fmt|:
	go build ./...c
vet|:fmt
	go vet ./...
build|:vet
	go build ./...
	