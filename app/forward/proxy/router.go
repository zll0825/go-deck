package proxy

import (
	"github.com/gin-gonic/gin"
)

func InitProxyRouter(router *gin.RouterGroup) {
	r := router.Group("gw")
	{
		r.Any("/:p1", Proxy)
		r.Any("/:p1/:p2", Proxy)
		r.Any("/:p1/:p2/:p3", Proxy)
		r.Any("/:p1/:p2/:p3/:p4", Proxy)
		r.Any("/:p1/:p2/:p3/:p4/:p5", Proxy)
	}
}
