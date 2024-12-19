package Server

import (
	"GrpcTest/proto"
	"context"
)

type Server struct {
	proto.UnimplementedHelloServer
}

func (s *Server) Say(ctx context.Context, request *proto.SayRequest) (*proto.SayResponse, error) {

	var SayResponse *proto.SayResponse
	SayResponse.Message = "Hello(Say) !" + request.Name

	return SayResponse, nil
}

func (s *Server) SayYouAagain(ctx context.Context, request *proto.SayRequest) (*proto.SayResponse, error) {
	var SayResponse *proto.SayResponse
	SayResponse.Message = "Hello(SayYouAagain) !" + request.Name

	return SayResponse, nil

}
