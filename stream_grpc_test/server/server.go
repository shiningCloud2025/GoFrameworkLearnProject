package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"GoFrameworkLearnProject/stream_grpc_test/proto"

	"google.golang.org/grpc"
)

const PORT = ":50052"

type server struct {
	proto.UnimplementedGreeterServer
}

// 服务端流模式(服务端一直发)
func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("req:%s--%v", req.Data, time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	return nil
}

// 客户端流模式(客户端一直发)
func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		if a, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Printf(a.Data)
		}

	}
	return nil
}

func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	// 不能先接受再传输，这是会有问题的 万一阻塞住了 就变成单向的了
	// 也就是说需要两者并行 这个时候就要用协程了
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息:" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			allStr.Send(&proto.StreamResData{Data: "你好，我是服务器"})
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
