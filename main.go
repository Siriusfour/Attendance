package main

import (
	"GrpcTest/Server"
	"GrpcTest/proto"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"net"
)

func main() {

	fmt.Println("===start===")
	Listener, err := net.Listen("tcp", ":5551")
	if err != nil {
		panic("5551启动错误：" + err.Error())
	}

	grpcServer := grpc.NewServer()

	proto.RegisterHelloServer(grpcServer, &Server.Server{})
	err = grpcServer.Serve(Listener)

	fmt.Println("Server is listening on port :50051")
	if err := grpcServer.Serve(Listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
