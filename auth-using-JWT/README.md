### To Genrate from proto file to gen dir

```
protoc \
  --proto_path=proto \
  --go_out=gen \
  --go-grpc_out=gen \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  proto/auth.proto
```
