package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
)

func InitBaseRouter(router *gin.RouterGroup) {
	r := router.Group("base")
	{
		r.POST("captcha", controller.Captcha)
		r.POST("login", controller.Login)
	}
}
