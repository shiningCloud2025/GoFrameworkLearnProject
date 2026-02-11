package main

import (
	"GoFrameworkLearnProject/grpc_error_test/proto"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "zyh"})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			panic("解析error失败")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}
	fmt.Println(r.Message)
}
