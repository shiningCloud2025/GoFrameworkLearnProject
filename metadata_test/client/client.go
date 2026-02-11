package main

import (
	"GoFrameworkLearnProject/grpc_test/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	//md := metadata.Pairs("timestamp",time.Now().Format(timestampFormat))
	md := metadata.New(map[string]string{
		"name":     "zyh",
		"password": "123456",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "zyh"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
