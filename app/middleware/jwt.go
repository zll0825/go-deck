package middleware

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/global"
	"go-deck/app/response"
	"go-deck/pkg/jwt"
	"strconv"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(c, gin.H{"reload": true}, "未登录或非法访问")
			c.Abort()
			return
		}
		j := jwt.NewJWT(global.Config.JwtConfig)
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				response.FailWithDetailed(c, gin.H{"reload": true}, "授权已过期")
				c.Abort()
				return
			}
			response.FailWithDetailed(c, gin.H{"reload": true}, err.Error())
			c.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + 60*60*24*7
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}
