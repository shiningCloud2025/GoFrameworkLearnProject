package main

import (
	"GoFrameworkLearnProject/stream_grpc_test/proto"
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// 服务端流模式
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	res, err := c.GetStream(context.Background(), &proto.StreamReqData{Data: "帅哥"})
	for {
		a, err := res.Recv() // 如果懂socket编程的话就明白了 send recv
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a)
	}

	// 客户端流模式
	putS, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		putS.Send(&proto.StreamReqData{Data: fmt.Sprintf("帅哥%d\n", i)})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	// 双向流模式
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到服务端消息:" + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			allStr.Send(&proto.StreamReqData{Data: "你好，我是客户端啊"})
		}
	}()
	wg.Wait()

}
