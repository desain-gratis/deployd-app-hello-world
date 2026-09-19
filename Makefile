ifneq (,$(wildcard .env))
include .env
export
endif

test:
	go test -v ./internal/...

build: test
	go build -o dist/hello-world cmd/hello-world/*.go
	chmod u+x dist/hello-world

run: test
	go run ./cmd/hello-world/*.go
