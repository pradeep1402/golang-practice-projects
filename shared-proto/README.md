- To genrate for auth
  ```
  protoc \
  --proto_path=auth \
  --go_out=gen/auth \
  --go-grpc_out=gen/auth \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  auth/auth.proto
  ```
