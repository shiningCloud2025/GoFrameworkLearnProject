package main

import (
	"net/http"

	"GinStart/chapter06/proto"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 最常用的就是
	/**
	TODO:
		json
		protobuf
		purejson:为了不让html的<>等被解析为二进制
	*/
	router.GET("/moreJSON", moreJSON)
	router.GET("/someProtobuf", returnProto)

	router.Run(":8083")
}

func returnProto(context *gin.Context) {
	course := []string{"python", "go", "微服务"}
	user := &proto.Teacher{
		Name:   "zyh",
		Course: course,
	}
	context.ProtoBuf(http.StatusOK, user)
}

func moreJSON(context *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "zyh"
	msg.Message = "这是一个测试json"
	msg.Number = 25
	context.JSON(http.StatusOK, msg)
}
