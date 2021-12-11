#grpc-server

Run protoc command:
```
protoc -I api/proto --go_out=pkg/api --go_opt=paths=source_relative --go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false api/proto/news.proto
```

Run grpc server:
`go run cmd/main.go`

GRPC client for this server is [here]()
