package main

import (
	"context"
	"fmt"

	//"GoFrameworkLearnProject/grpc_proto_test/proto"
	"GoFrameworkLearnProject/grpc_proto_test/proto-bak"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto_bak.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto_bak.HelloRequest{Name: "zyh"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
