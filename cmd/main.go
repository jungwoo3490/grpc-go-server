package main

import (
	"grpc-go-server/grpcapi"
	"grpc-go-server/handler"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":19001")
	if err != nil {
		log.Fatal("서버 시작 중 오류 발생: ", err)
	}

	grpcServer := grpc.NewServer()
	grpcapi.RegisterTodoServiceServer(grpcServer, handler.NewTodoHandler())

	log.Println("gRPC 서버 시작 중... (port: 19001)")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("서버 실행 중 오류 발생: ", err)
	}
}
