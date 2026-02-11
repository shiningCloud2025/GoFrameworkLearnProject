package main

import (
	"GoFrameworkLearnProject/grpc_error_test/proto"
	"net"
	"time"

	"context"

	"google.golang.org/grpc"
)

type Server struct {
	// 隐藏方法
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	time.Sleep(time.Second * 5)
	return &proto.HelloReply{
		Message: "hello" + request.Name,
	}, nil

	//return nil, status.Errorf(codes.NotFound, "记录未找到:%s", request.Name)
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
