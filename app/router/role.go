package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
	"go-deck/app/middleware"
)

func InitRoleRouter(Router *gin.RouterGroup) {
	router := Router.Group("role").Use(middleware.OperationRecord())
	{
		router.POST("create", controller.CreateRole)  // 新增角色
		router.POST("delete", controller.DeleteRole)  // 删除角色
		router.POST("update", controller.UpdateRole)  // 修改角色
		router.POST("list", controller.GetRoleList)   // 角色列表
		router.POST("detail", controller.GetRoleById) // 角色详情
		router.POST("all", controller.GetAllRoles)    // 所有角色

		router.POST("bindMenu", controller.BindMenu) // 绑定菜单
		router.POST("bindApi", controller.BindApi)   // 绑定api
	}
}
