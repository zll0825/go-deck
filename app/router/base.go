package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	router := Router.Group("base")
	{
		router.POST("captcha", controller.Captcha)
		router.POST("login", controller.Login)
	}
}
