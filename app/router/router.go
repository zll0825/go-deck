package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-deck/app/forward"
	"go-deck/app/global"
	"go-deck/app/middleware"

	_ "go-deck/docs"
)

func Routers() *gin.Engine {
	var r = gin.Default()
	// r.Use(middleware.LoadTls())  // https
	global.Logger.Info("use middleware logger")
	// 跨域
	r.Use(middleware.Cors())
	global.Logger.Info("use middleware cors")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.Logger.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	publicGroup := r.Group("")
	{
		InitBaseRouter(publicGroup) // 注册基础功能路由 不做鉴权

		InitRoleRouter(publicGroup)
		InitMenuRouter(publicGroup)
		InitApiRouter(publicGroup)
		InitDictRouter(publicGroup)
	}
	privateGroup := r.Group("")
	privateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		InitUserRouter(privateGroup) // 用户相关路由

		forward.InitForwardRouters(privateGroup) // 代理服务路由
	}
	global.Logger.Info("router register success")
	return r
}
