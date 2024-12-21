package main

import (
	BaseController "GrpcTest/Handler/Base/BaseController"
	"GrpcTest/MyConfig"
	"GrpcTest/MyProto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	fmt.Println("===start===")

	Listener, err := net.Listen("tcp", ":5551")
	if err != nil {
		panic("5551启动错误：" + err.Error())
	}

	MyConfig.InitConfig()

	grpcServer := grpc.NewServer()
	MyProto.RegisterViewServer(grpcServer, &BaseController.Base_Controller{})
	err = grpcServer.Serve(Listener)

	fmt.Println("Server is listening on port :50051")
	if err := grpcServer.Serve(Listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
