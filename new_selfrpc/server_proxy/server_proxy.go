package server_proxy

import (
	"GoFrameworkLearnProject/new_selfrpc/handler"
	"net/rpc"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}

// 如何做到解耦-我们关心的是函数 鸭子类型
func RegisterHelloService(srv HelloServicer) error {
	return rpc.RegisterName(handler.HelloServiceName, srv)

}
