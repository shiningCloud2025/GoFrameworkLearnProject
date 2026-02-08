package main

import "fmt"

func Add(a, b int) int {
	total := a + b
	return total
}

type Company struct {
	Name    string
	Address string
}

type Employee struct {
	Name    string
	company Company
}

type PrintResult struct {
	Info string
	Err  error
}

func RpcPrintln(employee Employee) {
	/**
	RPC中的两个点:传输协议、数据编码协议
	传输协议:
	http1.x,http2.0
	http协议底层也使用的tcp协议，http现在主流的是1.x，这种协议有性能权问题，一次性问题，一旦结果返回，连接断开
	1.直接自已基于tcp/udp协议去封装一层协议  mhttp  没有通用性 http2.0 既有http的特性，也有长连接的特性  grpc:http2.0




	数据编码协议:
	客户端
		1.建立连接,tcp/http
		2.将employee对象序列化成json字符串-序列化
		3.发送json字符串-调用成功后实际上接收到的是一个二进制数据
		4.等待服务器发送结果
		5.将服务器返回的数据解析成PrintResult对象-反序列号
	服务端:
		1.监听网络端口 80
		2.读取数据-二进制的json数据
		3.反序列化数据，变成Employee对象
		4.开始处理业务逻辑
		5.将处理结果序列号成JSON二进制数据
		6.将数据返回
	序列化和反序列化是可选的，不一定要采用json、xml、protobuf、msgpack
	http协议来说，有一个问题:一次性，对方返回了结果，链接断开 http2.0 长连接


	*/

}

func main() {

	/**
		现在我们想把Add函数变成一个远程的函数调用，也就意味着要把Add函数放到远程服务器上
		我们最原本的电商系统，这地方有一段逻辑，这个逻辑是扣减库存，但是库存服务是一个独立的系统，reduce,那么如何调用
		一定会牵扯到网络问题，做成一个web服务(gin、beego、net/httpserver)

	1.这个函数的调用参数如何传递-Json(json是一种数据格式的协议)/xml/protobuf/msgpack-编码协议，json并不是一个高性能的编码协议
		现在的网络调用有两个端-客户端、应该干嘛？将数据传输到gin
		gin-服务端，服务端负责解析数据



	*/
	fmt.Println(Add(1, 2))
	// 将这个打印的工作放在另一台服务器上，我需要将本地的内存对象strctu传递过去，可行的方式是序列化为json

	fmt.Println(
		Employee{
			Name: "zyh",
			company: Company{
				Name:    "zyh",
				Address: "北京市",
			},
		})

	// 远程服务需要将二进制对象反解成strcut对象
	// 搞这么麻烦，直接全部使用json去序列号不香吗?这样做法在浏览器和gin服务之间是可行的，但是如果你是一个大型的分布式系统
}
