package main

import (
	"fmt"

	"GoFrameworkLearnProject/new_selfrpc/client_proxy"
)

func main() {
	// 1.建立连接
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	// 1.只想写业务逻辑，不下关注每个函数的名称
	// 客户端部分
	//client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	//if err != nil {
	//	panic("连接失败")
	//}
	var reply *string = new(string)
	err := client.Hello("zyh", reply)
	//client.Hello // 不用自已去封装hello方法
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(*reply)

	// 1.这些概念在grpc中都有对应
	// 2.发自灵魂的拷问: server_proxy和client_proxy能否自动生成  为多种语言生成
	// 3.能 都能满足 这个就是protobuf+grpc

}
