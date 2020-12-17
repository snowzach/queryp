GITVERSION := $(shell git describe --dirty --always --tags --long)
GOPATH ?= ${HOME}/go
PACKAGENAME := $(shell go list -m -f '{{.Path}}')
TOOLS := ${GOPATH}/src/github.com/golang/protobuf/proto \
	${GOPATH}/bin/protoc-gen-go

PROTOS := ./queryp.pb.go

.PHONY: default
default: $(PROTOS)

tools: $(TOOLS)

${GOPATH}/src/github.com/golang/protobuf/proto:
	go get github.com/golang/protobuf/proto

${GOPATH}/bin/protoc-gen-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go

# Handle compiling the protobufs
%.pb.go: %.proto tools
	protoc --go_out=$(PROTO_MAPPINGS),paths=source_relative:. $*.proto

.PHONY: test
test: tools ${PROTOS}
	go test -cover ./...
