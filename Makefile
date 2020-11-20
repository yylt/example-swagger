ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

all: build

# Build manager binary
build: vendor fmt vet
	mkdir -p bin
	go build -o bin/example server.go

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vendor against code
vendor:
	go mod vendor

# Run go vet against code
vet:
	go vet ./...

# Generate code
generate:
	@go generate ./...

