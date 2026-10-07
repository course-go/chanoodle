.PHONY: build
build:
	go build -o bin/chanoodle cmd/chanoodle/main.go

.PHONY: dev
dev:
	go run cmd/chanoodle/main.go

.PHONY: test
test:
	go test -cover -race -timeout 30s ./...

.PHONY: lint
lint:
	go tool golangci-lint run

.PHONY: format
format:
	go tool golangci-lint fmt

.PHONY: format-diff
format-diff:
	go tool golangci-lint fmt --diff-colored
