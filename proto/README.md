# Proto files for all microservices

This dir contains the proto-files for all microservices.

## Prerequisites

Before you can build and run a project using grpc, go and proto, you'll need to install the following tools:

- [Protocol Buffers Compiler (protoc)](https://github.com/protocolbuffers/protobuf)
- Go (Go 1.16 or higher is recommended)

You will also need to install the Go-specific Protocol Buffers and gRPC plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
Make sure that $GOBIN and $GOPATH are in your PATH.

Generate Go code from the .proto file:
bash
Copy code
````bash

protoc --proto_path=.<proto dir> --go_out=. --go-grpc_out=.<proto dir> .<proto file path>

