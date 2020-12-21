package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
	"go-deck/app/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	router := Router.Group("user").Use(middleware.OperationRecord())
	{
		router.POST("create", controller.CreateUser)  // 创建用户
		router.POST("delete", controller.DeleteUser)  // 删除用户
		router.POST("update", controller.UpdateUser)  // 修改用户
		router.POST("list", controller.GetUserList)   // 用户列表
		router.POST("detail", controller.GetUserById) // 用户详情
		router.POST("all", controller.GetAllUsers)    // 所有用户

		router.POST("bindRole", controller.BindRole) // 给用户分配角色
	}
}
