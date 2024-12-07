package router

import (
	"ctw-interview/controller"
	"ctw-interview/middleware"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
	apiRouter := router.Group("/api")
	// 压缩
	apiRouter.Use(gzip.Gzip(gzip.DefaultCompression))
	apiRouter.Use(middleware.GlobalAPIRateLimit())

	// 用户
	authRouter := apiRouter.Group("/auth")
	{
		authRouter.POST("/users", middleware.UserAuth(), controller.Login)
	}
}
