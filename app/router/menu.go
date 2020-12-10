package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
	"go-deck/app/middleware"
)

func InitMenuRouter(Router *gin.RouterGroup) {
	router := Router.Group("menu").Use(middleware.OperationRecord())
	{
		router.POST("create", controller.CreateMenu) // 新增菜单
		router.POST("update", controller.UpdateMenu) // 修改菜单
	}
}
