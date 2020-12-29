package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
	"go-deck/app/middleware"
)

func InitRoleRouter(router *gin.RouterGroup) {
	r := router.Group("role").Use(middleware.OperationRecord())
	{
		r.POST("create", controller.CreateRole)  // 新增角色
		r.POST("delete", controller.DeleteRole)  // 删除角色
		r.POST("update", controller.UpdateRole)  // 修改角色
		r.POST("list", controller.GetRoleList)   // 角色列表
		r.POST("detail", controller.GetRoleById) // 角色详情
		r.POST("all", controller.GetAllRoles)    // 所有角色

		r.POST("bindMenu", controller.BindMenu) // 绑定菜单
		r.POST("bindApi", controller.BindApi)   // 绑定api
	}
}
