GITVERSION := $(shell git describe --dirty --always --tags --long)
GOPATH ?= ${HOME}/go
PACKAGENAME := $(shell go list -m -f '{{.Path}}')
TOOLS := ${GOPATH}/src/github.com/golang/protobuf/proto \
	${GOPATH}/bin/protoc-gen-gogoslick
export PROTOBUF_INCLUDES = -I. -I/usr/include -I${GOPATH}/src

PROTOS := ./queryp.pb.go

.PHONY: default
default: ${PROTOS}

tools: 

${GOPATH}/src/github.com/golang/protobuf/proto:
	go get github.com/golang/protobuf/proto

${GOPATH}/bin/protoc-gen-gogoslick:
	go get github.com/gogo/protobuf/protoc-gen-gogoslick

# Handle any non-specific protobufs
%.pb.go: %.proto ${TOOLS}
	protoc ${PROTOBUF_INCLUDES} --gogoslick_out=paths=source_relative,plugins=grpc:. $*.proto

.PHONY: test
test: tools ${PROTOS}
	go test -cover ./...
