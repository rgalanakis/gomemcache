.DEFAULT_GOAL := test

fmt:
	go fmt ./...

vet:
	go vet

test:
	go test -race ./...

coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...
