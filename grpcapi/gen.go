package grpcapi

//go:generate -command compile_proto protoc -I../protos
//go:generate compile_proto todo.proto --proto_path=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative
