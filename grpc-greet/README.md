## Cmd to execute the proto file

````protoc -I./proto \
  --go_out=module=grpc-service-greet:. \
  --go-grpc_out=module=grpc-service-greet:. \
  ./proto/dummy.proto```
````
