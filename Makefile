PATH := ${PWD}/bin:${PATH}
export PATH

.DEFAULT_GOAL := build

.PHONY: tools
tools:
	go generate ./tools.go

.PHONY: clean
clean:
	rm -rf ./bin/*

.PHONY: gen
gen: tools
	go generate ./...

.PHONY: grapi
build:
	go build -v -o ./bin/grapi ./cmd/grapi

.PHONY: lint
lint: ./bin/reviewdog ./bin/golangci-lint
ifdef CI
	reviewdog -reporter=github-pr-review
else
	reviewdog -diff="git diff master"
endif

.PHONY: test
test:
	go test -v ./...

.PHONY: cover
cover:
	go test -v -coverprofile coverage.txt -covermode atomic ./...

.PHONY: test-e2e
test-e2e: build
	go test -v -timeout 4m ./_tests/e2e --grapi=$$PWD/bin/grapi --revision="$(TARGET_REVISION)"

# linters
bin/reviewdog:
	curl -sfL https://raw.githubusercontent.com/reviewdog/reviewdog/master/install.sh | sh -s -- -b ./bin v0.9.12

bin/golangci-lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ./bin v1.17.1
