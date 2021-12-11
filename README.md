# gRPC server example

Run protoc command:
```
protoc -I api/proto --go_out=pkg/api --go_opt=paths=source_relative --go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false api/proto/news.proto
```

Run grpc server:
`go run cmd/main.go`

GRPC client for this server is [here](https://github.com/KonstantinP85/grpc-client)

![vectorpaint](https://user-images.githubusercontent.com/74908254/145671274-f0303fe5-b132-4221-ac6d-056834ceabe6.png)
