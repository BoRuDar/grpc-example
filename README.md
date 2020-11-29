# grpc-example

1. Install `protoc` from https://github.com/protocolbuffers/protobuf/releases
2. Install the protocol compiler plugins for Go: 
```bash
go get google.golang.org/protobuf/cmd/protoc-gen-go \
            google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go \
            google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

# Doc
https://developers.google.com/protocol-buffers/docs/proto3