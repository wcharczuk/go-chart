all: test

test:
	@go test ./...

.PHONY: profanity
profanity:
	@profanity -include="*.go,Makefile,README.md"

cover:
	@go test -short -covermode=set -coverprofile=profile.cov
	@go tool cover -html=profile.cov
	@rm profile.cov

deps:
	@go get -u github.com/blend/go-sdk/_bin/profanity
