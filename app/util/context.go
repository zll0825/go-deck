package util

import (
	"github.com/gin-gonic/gin"
	"go-deck/pkg/jwt"
)

// 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserRoleIds(c *gin.Context) []int {
	if claims, exists := c.Get("claims"); !exists {
		return []int{}
	} else {
		waitUse := claims.(*jwt.CustomClaims)
		return waitUse.RoleIds
	}
}