package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	goodsGroup := router.Group("/goods")
	{
		goodsGroup.GET("/list", goodList)
		//goodsGroup.GET("/:id/:action", goodsDetail) // 获取商品id为1的详细信息 格式
		// 这种*一般用在文件业务上，其他业务基本用不到
		goodsGroup.GET("/:id/*action", goodsDetail) // 获取商品id为1的详细信息 格式
		goodsGroup.POST("/add", createGoods)
	}

	//router.GET("/goods",goodList)
	//router.GET("goods/1",goodsDetail)
	//router.GET("goods/add",createGoods)

	router.Run(":8083")
}

func createGoods(context *gin.Context) {

}

func goodsDetail(context *gin.Context) {
	id := context.Param("id")
	action := context.Param("action")
	context.JSON(http.StatusOK, gin.H{"id": id, "action": action})
}

func goodList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"name": "goodsList",
	})
}
