package router

import (	
	"booksapi/src/controllers"
	"booksapi/src/middleware"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Options)
	var wechat controllers.Wechat

	r.GET("/", controllers.SayHello)
	r.GET("/api", controllers.SayHello)
	r.GET("/api/books", controllers.GetAllBooks)

	v1 := r.Group("/api/v1")
	v1.Use(middleware.Authorization())
	
	v1.GET("/wechat/login", wechat.WeChatLogin)
	v1.POST("/wechat/getuserinfo", wechat.WechatGetUserInfo)
	v1.GET("/users/myuserinfo", controllers.GetMyUserInfo)

	return r
}