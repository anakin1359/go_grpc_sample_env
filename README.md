# go_grpc_sample_env

### exec protoc
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative proto/cat.proto
```

### library
```
$ tree
.
├── README.md
├── client
│   └── main.go       
├── go.mod
├── go.sum
├── proto
│   ├── cat.pb.go     
│   ├── cat.proto     
│   └── cat_grpc.pb.go
├── server
│   └── main.go
└── service
    └── main.go

4 directories, 9 files
```

### gRPC exec
```
go run server/main.go
```
```
go run client/main.go
```