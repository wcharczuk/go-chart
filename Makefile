all: test

tools:
	@go get -u github.com/blend/go-sdk/_bin/coverage
	@go get -u github.com/blend/go-sdk/_bin/profanity

test:
	@go test ./...

.PHONY: profanity
profanity:
	@profanity -include="*.go,Makefile,README.md"

cover:
	@coverage