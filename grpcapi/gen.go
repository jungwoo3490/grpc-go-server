package api

//go:generate -command compile_proto protoc -I../protos
//go:generate compile_proto todo.proto --proto_path=. --go_out=. --go-grpc_out=.
