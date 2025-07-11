## Cmd to execute the proto file

```
export PATH=$PATH:$(go env GOPATH)/bin

protoc -I./proto \
  --go_out=module=grpc-service-greet:. \
  --go-grpc_out=module=grpc-service-greet:. \
  ./proto/greet.proto
```
