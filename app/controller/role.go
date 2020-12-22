package controller

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller/dto"
	"go-deck/app/global"
	"go-deck/app/model/entity"
	"go-deck/app/response"
	"go-deck/app/service"
	"go.uber.org/zap"
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
// @Summary 删除Role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.DeleteRole true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /role/delete [post]
func DeleteRole(c *gin.Context) {
	var req dto.DeleteRole
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	if err := service.DeleteByIds(entity.Role{}, req.Ids); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage(c, "删除失败")
	} else {
		response.OkWithMessage(c, "删除成功")
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

	api := entity.Role{
		ID:        req.Id,
		Name:      req.Name,
		Key:       req.Key,
	}

	if err := service.UpdateById(&api, &entity.Api{}, req.Id); err != nil {
		global.Logger.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage(c, "修改失败")
	} else {
		response.OkWithMessage(c, "修改成功")
	}
}

// @Tags Role
// @Summary 分页获取Role列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.SearchRole true "分页获取Role列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /role/list [post]
func GetRoleList(c *gin.Context) {
	var req dto.SearchRole
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	if total, list, err := service.GetRoleList(req); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(c, "获取失败")
	} else {
		response.OkWithDetailed(c, response.PageResult{
			List:  list,
			Total: total,
			Page:  req.Page,
			Size:  req.Size,
		}, "获取成功")
	}
}

// @Tags Role
// @Summary 根据id获取Role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.DetailRole true "根据id获取Role"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /role/detail [post]
func GetRoleById(c *gin.Context) {
	var req dto.DetailRole
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	role := entity.Role{}
	err := service.DetailById(&role, req.Id)
	if err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(c, "获取失败")
	} else {
		response.OkWithData(c, role)
	}
}

// @Tags Role
// @Summary 获取所有的Role 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /role/all [post]
func GetAllRoles(c *gin.Context) {
	if apis, err := global.DB.AllRoles(); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(c, "获取失败")
	} else {
		response.OkWithDetailed(c, apis, "获取成功")
	}
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
