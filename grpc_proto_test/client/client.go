package main

import (
	"GoFrameworkLearnProject/grpc_proto_test/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "zyh"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
