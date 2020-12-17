GITVERSION := $(shell git describe --dirty --always --tags --long)
GOPATH ?= ${HOME}/go
PACKAGENAME := $(shell go list -m -f '{{.Path}}')
TOOLS := ${GOPATH}/src/github.com/golang/protobuf/proto \
	${GOPATH}/bin/protoc-gen-gogoslick

# Include protobuf sources
PROTOBUF_INCLUDES := -I.
PROTOBUF_INCLUDES := $(PROTOBUF_INCLUDES) -I$(shell go list -e -f '{{.Dir}}' github.com/gogo/protobuf/gogoproto)/..

# Map protobuf includes to the Go package containing the generated Go code.
PROTO_MAPPINGS :=

PROTOS := ./queryp.pb.go

.PHONY: default
default: $(PROTOS)

tools: $(TOOLS)

${GOPATH}/src/github.com/golang/protobuf/proto:
	go get github.com/golang/protobuf/proto

${GOPATH}/bin/protoc-gen-gogofaster:
	go get github.com/gogo/protobuf/protoc-gen-gogofaster

# Handle compiling the protobufs
%.pb.go: %.proto tools
	protoc ${PROTOBUF_INCLUDES} --gogofaster_out=$(PROTO_MAPPINGS),paths=source_relative:. $*.proto

.PHONY: test
test: tools ${PROTOS}
	go test -cover ./...
