package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
)

func InitProxyRouter(Router *gin.RouterGroup) {
	router := Router.Group("gw")
	{
		router.Any("/:p1", controller.Proxy)
		router.Any("/:p1/:p2", controller.Proxy)
		router.Any("/:p1/:p2/:p3", controller.Proxy)
		router.Any("/:p1/:p2/:p3/:p4", controller.Proxy)
		router.Any("/:p1/:p2/:p3/:p4/:p5", controller.Proxy)
	}
}