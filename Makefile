all: test

ci: profanity coverage

new-install:
	@go get -v -u ./...
	@go get -v -u github.com/blend/go-sdk/cmd/coverage
	@go get -v -u github.com/blend/go-sdk/cmd/profanity

generate:
	@go generate ./...

test:
	@go test ./...

.PHONY: profanity
profanity:
	@profanity

coverage:
	@coverage
