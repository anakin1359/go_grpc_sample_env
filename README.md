# go_grpc_sample_env

https://qiita.com/gp333/items/24733ff4b95d9a4bb977


```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative proto/cat.proto
```
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
```
go run server/main.go
```
```
go run client/main.go
```