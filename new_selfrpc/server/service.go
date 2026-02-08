package main

import (
	"GoFrameworkLearnProject/new_selfrpc/handler"
	"net"
	"net/rpc"

	"GoFrameworkLearnProject/new_selfrpc/server_proxy"
)

func main() {
	// 1.实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	// 2.注册处理逻辑 handler
	_ = server_proxy.RegisterHelloService(&handler.NewHelloService{})
	//_ = rpc.RegisterName(handler.HelloServiceName, &handler.HelloService{})
	// 3.启动服务
	for {
		conn, _ := listener.Accept() // 当一个新连接进来的适合
		go rpc.ServeConn(conn)
	}

	// 一连串的代码大部分都是net包好像和rpc没有关系
	// 不行 rpc调用中有几个问题需要解决:1.call id 2.序列化和反序列化
	// 相较python下的开发而言，这个显得不是很好用
	// 可以跨语言调用吗 1.go语言rpc的序列化和反序列化协议是什么 Gob 2.能否替换成常见的序列化

}
