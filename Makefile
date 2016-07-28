all: test

test:
	@go test ./...

cover:
	@go test -short -covermode=set -coverprofile=profile.cov
	@go tool cover -html=profile.cov
	@rm profile.cov