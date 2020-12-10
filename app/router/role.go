package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
	"go-deck/app/middleware"
)

func InitRoleRouter(Router *gin.RouterGroup) {
	router := Router.Group("role").Use(middleware.OperationRecord())
	{
		router.POST("create", controller.CreateRole) // 新增角色
		router.POST("update", controller.UpdateRole) // 修改角色

		router.POST("bindMenu", controller.BindMenu) // 绑定菜单
	}
}
