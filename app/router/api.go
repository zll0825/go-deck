package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
	"go-deck/app/middleware"
)

func InitApiRouter(router *gin.RouterGroup) {
	r := router.Group("api").Use(middleware.OperationRecord())
	{
		r.POST("create", controller.CreateApi)  // 创建Api
		r.POST("delete", controller.DeleteApi)  // 删除Api
		r.POST("update", controller.UpdateApi)  // 更新api
		r.POST("list", controller.GetApiList)   // 获取Api列表
		r.POST("detail", controller.GetApiById) // 获取单条Api消息
		r.POST("all", controller.GetAllApis)    // 获取所有api
	}
}
