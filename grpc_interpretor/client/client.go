package main

import (
	"GoFrameworkLearnProject/grpc_test/proto"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func main() {
	interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Println("耗时:%s", time.Since(start))
		return err
	}
	opt := grpc.WithUnaryInterceptor(interceptor)
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), opt)
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
