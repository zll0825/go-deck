package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-deck/app/global"
	"go-deck/app/middleware"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	// Router.Use(middleware.LoadTls())  // https
	global.Logger.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	global.Logger.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.Logger.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权

		InitRoleRouter(PublicGroup)
		InitMenuRouter(PublicGroup)
		InitApiRouter(PublicGroup)
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		InitUserRouter(PrivateGroup) // 用户相关路由
	}
	global.Logger.Info("router register success")
	return Router
}
