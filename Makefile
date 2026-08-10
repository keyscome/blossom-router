PREFIX ?= $(HOME)/.local

.PHONY: build test install

build:
	go build -o bin/bloom ./cmd/bloom

test:
	go test ./...

install: build
	mkdir -p $(PREFIX)/bin
	install -m 755 bin/bloom $(PREFIX)/bin/bloom
	@echo "Installed bloom to $(PREFIX)/bin/bloom"
	@case ":$$PATH:" in *":$(PREFIX)/bin:"*) ;; *) echo 'Add it to this shell with: export PATH="$(PREFIX)/bin:$$PATH"' ;; esac
