.PHONY: build test install

build:
	go build -o bin/dd ./cmd/dd

test:
	go test ./...

install:
	go install ./cmd/dd
