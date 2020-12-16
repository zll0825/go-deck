package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
	"go-deck/app/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	router := Router.Group("user").Use(middleware.OperationRecord())
	{
		router.POST("create", controller.CreateUser) // 创建用户
		router.POST("bindRole", controller.BindRole) // 给用户分配角色
		//UserRouter.POST("changePassword", controllers.ChangePassword)     // 修改密码
		//UserRouter.POST("getUserList", controllers.GetUserList)           // 分页获取用户列表
		//UserRouter.POST("setUserAuthority", controllers.SetUserAuthority) // 设置用户权限
		//UserRouter.DELETE("deleteUser", controllers.DeleteUser)           // 删除用户
		//UserRouter.PUT("setUserInfo", controllers.SetUserInfo)            // 设置用户信息
	}
}
