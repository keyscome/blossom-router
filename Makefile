.PHONY: build test install

build:
	go build -o bin/bloom ./cmd/bloom

test:
	go test ./...

install:
	go install ./cmd/bloom
