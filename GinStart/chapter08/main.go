package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "123456")
		// 让原本该执行的逻辑继续执行
		c.Next()

		end := time.Since(t)
		fmt.Printf("耗时:%V\n", end)
		status := c.Writer.Status()
		fmt.Println("状态:", status)
	}
}

func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for k, v := range c.Request.Header {
			if k == "X-Token" {
				token = v[0]
			} else {
				fmt.Println(k, v)
			}
			fmt.Println(k, v, token)
		}
		if token != "zyh" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "未登录",
			})
			//return
			// 挑战 为什么连return都阻止不了后续逻辑的执行
			c.Abort()
		}
		c.Next()
	}
}

func main() {
	router := gin.New()
	// 使用logger和recovery中间件,全局使用
	//router.Use(gin.Logger(), gin.Recovery())
	//authrized := router.Group("/goods")
	router.Use(MyLogger())
	router.Use(TokenRequired())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8083")

}
