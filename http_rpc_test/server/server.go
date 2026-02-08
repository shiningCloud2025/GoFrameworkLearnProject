package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	// 返回值是通过修改reply的值
	*reply = "hello," + request
	return nil
}

func main() {
	// 1.实例化一个server
	_ = rpc.RegisterName("HelloService", &HelloService{})
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234", nil)

}

// 一连串的代码大部分都是net包好像和rpc没有关系
// 不行 rpc调用中有几个问题需要解决:1.call id 2.序列化和反序列化
// 相较python下的开发而言，这个显得不是很好用
// 可以跨语言调用吗 1.go语言rpc的序列化和反序列化协议是什么 Gob 2.能否替换成常见的序列化
