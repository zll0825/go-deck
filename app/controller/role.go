package controller

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller/dto"
	"go-deck/app/global"
	"go-deck/app/model/entity"
	"go-deck/app/response"
	"go-deck/app/service"
)

// @Tags Role
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.CreateRole true "角色名，角色key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /role/create [post]
func CreateRole(c *gin.Context) {
	var req dto.CreateRole

	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	role := entity.Role{
		Name: req.Name,
		Key:  req.Key,
	}
	if err := global.DB.CreateRole(&role); err != nil {
		response.FailWithMessage(c, "创建角色失败")
	} else {
		response.OkWithMessage(c, "创建成功")
	}
}

// @Tags Role
// @Summary 更新角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.UpdateRole true "角色名，角色key"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /role/update [post]
func UpdateRole(c *gin.Context) {
	var req dto.UpdateRole

	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	// todo:更新角色
}

// @Tags Role
// @Summary 给角色绑定菜单权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.BindRoleMenu true "角色id, 菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /role/bindMenu [post]
func BindMenu(c *gin.Context) {
	var req dto.BindRoleMenu

	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	// 绑定角色菜单
	if err := service.BindRoleMenu(req.RoleId, req.MenuIds); err != nil {
		response.FailWithMessage(c, "绑定权限失败")
		c.Abort()
		return
	}

	response.OkWithMessage(c, "绑定权限成功")
	return
}

// @Tags Role
// @Summary 给角色绑定接口权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.BindRoleApi true "角色id, apiId"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /role/bindApi [post]
func BindApi(c *gin.Context) {
	var req dto.BindRoleApi

	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	// 绑定角色api
	if err := service.BindRoleApi(req.RoleId, req.ApiIds); err != nil {
		response.FailWithMessage(c, "绑定权限失败")
		c.Abort()
		return
	}

	response.OkWithMessage(c, "绑定权限成功")
	return
}
