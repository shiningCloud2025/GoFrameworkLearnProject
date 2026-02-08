package main

import (
	"GoFrameworkLearnProject/grpc_test/proto"
	"net"

	"context"

	"google.golang.org/grpc"
)

type Server struct {
	// 隐藏方法
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello" + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("failed to listen" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc" + err.Error())
	}
}
