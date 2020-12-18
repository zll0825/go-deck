package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
	"go-deck/app/middleware"
)

func Init{{.StructName}}Router(Router *gin.RouterGroup) {
	router := Router.Group("{{.StructName}}").Use(middleware.OperationRecord())
	{
		router.POST("create", controller.Create{{.StructName}})  // 创建{{.StructName}}
		router.POST("delete", controller.Delete{{.StructName}})  // 删除{{.StructName}}
		router.POST("list", controller.Get{{.StructName}}List)   // 获取{{.StructName}}列表
		router.POST("detail", controller.Get{{.StructName}}ById) // 获取单条{{.StructName}}消息
		router.POST("update", controller.Update{{.StructName}})  // 更新{{.StructName}}
	}
}
