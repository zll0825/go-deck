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
		router.POST("delete", controller.DeleteMenu) // 删除菜单
		router.POST("tree", controller.GetMenuTree)   // 获取菜单树
		router.POST("detail", controller.GetMenuById) // 获取菜单信息
		router.POST("update", controller.UpdateMenu)  // 更新菜单
		router.POST("all", controller.GetAllMenus)    // 获取所有菜单
	}
}
