all: build

build:
	go build

install:
	go install

test:
	./run_tests.sh

fmt:
	go fmt `go list ./... | grep -v -vender`

lint:
	golint `go list ./... | grep -v -vender`

vet:
	go vet `go list ./... | grep -v -vender`

clean:
	go clean `go list ./... | grep -v -vender`
	rm -f coverage.txt profile.out

.PHONY: all build install test fmt lint vet clean
