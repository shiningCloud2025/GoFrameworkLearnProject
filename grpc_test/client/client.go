package main

import (
	"GoFrameworkLearnProject/grpc_test/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
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
