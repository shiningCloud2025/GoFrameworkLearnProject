package client_proxy

import (
	"net/rpc"

	"GoFrameworkLearnProject/new_selfrpc/handler"
)

type HelloServiceSub struct {
	*rpc.Client
}

// 在go语言中没有类、对象，就意味着没有初始化
func NewHelloServiceClient(protol, address string) HelloServiceSub {
	conn, err := rpc.Dial(protol, address)
	if err != nil {
		panic("connect fail")
	}
	return HelloServiceSub{conn}
}

func (c *HelloServiceSub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
