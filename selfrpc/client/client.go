package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1.建立连接
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		panic("连接失败")
	}
	var reply *string = new(string)
	err = client.Call("HelloService.Hello", "zyh", reply)
	//client.Hello // 不用自已去封装hello方法
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(*reply)
}
