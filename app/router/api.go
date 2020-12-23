package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
	"go-deck/app/middleware"
)

func InitApiRouter(Router *gin.RouterGroup) {
	router := Router.Group("api").Use(middleware.OperationRecord())
	{
		router.POST("create", controller.CreateApi)  // 创建Api
		router.POST("delete", controller.DeleteApi)  // 删除Api
		router.POST("update", controller.UpdateApi)  // 更新api
		router.POST("list", controller.GetApiList)   // 获取Api列表
		router.POST("detail", controller.GetApiById) // 获取单条Api消息
		router.POST("all", controller.GetAllApis)    // 获取所有api
	}
}
