package main

import (
	"GoFrameworkLearnProject/grpc_test/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type customCredential struct {
}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "zyh",
		"appkey": "i am zyh",
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security.
func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {
	// 拦截器方式1
	//interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//	start := time.Now()
	//	md := metadata.New(map[string]string{
	//		"appid":  "zyh",
	//		"appkey": "i am zyh",
	//	})
	//	ctx = metadata.NewOutgoingContext(context.Background(), md)
	//	err := invoker(ctx, method, req, reply, cc, opts...)
	//	fmt.Println("耗时:%s", time.Since(start))
	//	return err
	//}
	// 拦截器方式2
	credentials := grpc.WithPerRPCCredentials(customCredential{})

	//opt := grpc.WithUnaryInterceptor(interceptor)
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), credentials)
	//conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), opt)
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
