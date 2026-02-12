package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/welcome", welcome)
	router.POST("/form", formPost)
	router.POST("/post", getPost)

	router.Run(":8083")
}

func getPost(context *gin.Context) {
	id := context.Query("id")
	page := context.DefaultQuery("page", "0")
	name := context.PostForm("name")
	message := context.DefaultPostForm("message", "信息")
	context.JSON(http.StatusOK, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})
}

func formPost(context *gin.Context) {
	message := context.PostForm("message")
	nick := context.DefaultPostForm("nick", "nimingyonghu")
	context.JSON(http.StatusOK, gin.H{
		"message": message,
		"nick":    nick,
	})
}

func welcome(context *gin.Context) {
	firstName := context.DefaultQuery("firstname", "zyh")
	lastName := context.DefaultQuery("lastname", "zyh1")
	context.JSON(http.StatusOK, gin.H{
		"firstName": firstName,
		"lastName":  lastName,
	})
}
