package forward

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/forward/proxy"
	"go-deck/app/forward/transport"
)

func InitForwardRouters(router *gin.RouterGroup)  {
	// 通配路由
	proxy.InitProxyRouter(router)

	// 代理路由
	transport.InitTransportRouter(router)
}
