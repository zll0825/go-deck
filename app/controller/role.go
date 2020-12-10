package controller

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller/dto"
	"go-deck/app/global"
	"go-deck/app/model/entity"
	"go-deck/app/response"
	"go-deck/app/service"
)

// @Tags Authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authority/createAuthority [post]
func CreateRole(c *gin.Context) {
	var req dto.CreateRole

	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage("参数验证失败", c)
		c.Abort()
		return
	}

	role := entity.Role{
		Name: req.RoleName,
	}
	if err := global.DB.CreateRole(&role); err != nil {
		response.FailWithMessage("创建角色失败", c)
	} else {
		response.OkWithMessage( "创建成功", c)
	}
}

func UpdateRole(c *gin.Context) {
	var req dto.UpdateRole

	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage("参数验证失败", c)
		c.Abort()
		return
	}

	// todo:更新角色
}

// 给角色绑定菜单权限
func BindMenu(c *gin.Context)  {
	var req dto.AddRoleMenu

	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage("参数验证失败", c)
		c.Abort()
		return
	}

	// 查出现有角色下的菜单
	if err := service.BindRoleMenu(req.RoleId, req.MenuIds); err != nil {
		response.FailWithMessage("绑定权限失败", c)
		c.Abort()
		return
	}

	response.OkWithMessage("绑定权限成功", c)
	return
}
