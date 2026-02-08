package main

import (
	"GoFrameworkLearnProject/selfrpc/proto"
	"fmt"

	"github.com/golang/protobuf/proto"
)

type Hello struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"course"`
}

func main() {
	req := helloworld.HelloRequest{
		Name:    "zyh",
		Age:     18,
		Courses: []string{"go", "gin", "微服务"},
	}
	//jsonStruct := Hello{Name: "zyh",
	//	Age:     18,
	//	Courses: []string{"go", "gin", "微服务"}}
	//jsonRsp, _ := json.Marshal(jsonStruct)
	//fmt.Println(string(jsonRsp))
	//fmt.Println(len(jsonRsp))
	rsp, _ := proto.Marshal(&req) //具体的编码是如何做到的，可以百度一下protobuf的原理 varint
	newReq := helloworld.HelloRequest{}
	_ = proto.Unmarshal(rsp, &newReq)
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses)
	//fmt.Println(rsp)
	//fmt.Println(string(rsp))
	//fmt.Println(len(rsp))
}
