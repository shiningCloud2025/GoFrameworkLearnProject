package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	// 优雅退出,当我们关闭的时候应该做的后续处理
	// 微服务 启动之前 或者启动之后会做一件事 将当前服务的ip地址和端口号注册到注册中心
	// 我们当前的服务停止了以后并没有告知注册中心

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	go func() {
		router.Run(":8083")
	}()

	// 如果想要接收到信号   kill -9是强杀命令 是接受不到的  kill是可以的哈
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 处理后续的逻辑
	fmt.Println("关闭server中")
	fmt.Println("注销服务")

}
