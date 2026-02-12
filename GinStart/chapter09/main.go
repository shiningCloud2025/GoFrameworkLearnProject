package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// LoadHTMLFiles会制定目录下的文件加载好，相对目录
	// 为什么我们通过goland运行main.go的时候没有生成exe文件（goland会生成临时地址 可以参考tomcat）
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(dir)
	//router.LoadHTMLFiles("templates/index.html", "templates/goods.html")
	router.LoadHTMLGlob("templates/**/*")

	// 如果没有在模版中使用define定义,那么我们就可以使用默认的文件名来找
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "zyh",
		})
	})

	router.GET("/goods/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods/list.html", gin.H{
			"title": "zyh",
		})
	})

	router.GET("/users/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/list.html", gin.H{
			"title": "zyh",
		})
	})

	router.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"name": "微服务开发",
		})
	})
	router.Run(":8083")
}
