.PHONY: all test run

all: run

fmt:
	gofmt -w .

test:
	go test -v -tags=unit $$(go list ./... | grep -v '/vendor/')

run:
	go run main.go
