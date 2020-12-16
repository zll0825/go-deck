package middleware

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/global"
	"go-deck/app/model"
	"go-deck/app/response"
	"go-deck/pkg/casbin"
	"go-deck/pkg/jwt"
)

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*jwt.CustomClaims)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.Username
		e, err := casbin.NewCasbin(global.Config.CasbinConfig, model.SystemDB())
		if err != nil {
			panic(err)
		}
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if success {
			c.Next()
		} else {
			response.FailWithDetailed(c, gin.H{}, "权限不足")
			c.Abort()
			return
		}
	}
}
