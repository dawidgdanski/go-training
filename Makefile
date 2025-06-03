.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt:
		go fmt ./...

vet: fmt
		go vet ./...

generate: vet
		go generate ./...

build: generate
		go build

clean:
		go clean

test:
		go test ./...
