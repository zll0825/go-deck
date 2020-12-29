package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
	"go-deck/app/middleware"
)

func InitMenuRouter(router *gin.RouterGroup) {
	r := router.Group("menu").Use(middleware.OperationRecord())
	{
		r.POST("create", controller.CreateMenu) // 新增菜单
		r.POST("delete", controller.DeleteMenu) // 删除菜单
		r.POST("tree", controller.GetMenuTree)   // 获取菜单树
		r.POST("detail", controller.GetMenuById) // 获取菜单信息
		r.POST("update", controller.UpdateMenu)  // 更新菜单
		r.POST("all", controller.GetAllMenus)    // 获取所有菜单
	}
}
