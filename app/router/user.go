package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
	"go-deck/app/middleware"
)

func InitUserRouter(router *gin.RouterGroup) {
	r := router.Group("user").Use(middleware.OperationRecord())
	{
		r.POST("create", controller.CreateUser)  // 创建用户
		r.POST("delete", controller.DeleteUser)  // 删除用户
		r.POST("update", controller.UpdateUser)  // 修改用户
		r.POST("list", controller.GetUserList)   // 用户列表
		r.POST("detail", controller.GetUserById) // 用户详情
		r.POST("all", controller.GetAllUsers)    // 所有用户

		r.POST("bindRole", controller.BindRole) // 给用户分配角色
	}
}
