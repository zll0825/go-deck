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

// @Tags Menu
// @Summary 创建菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.CreateMenu true "菜单名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /menu/create [post]
func CreateMenu(c *gin.Context) {
	var req dto.CreateMenu

	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	menu := entity.Menu{
		MenuLevel:   req.MenuLevel,
		ParentId:    req.ParentId,
		Path:        req.Path,
		Name:        req.Name,
		Hidden:      req.Hidden,
		Component:   req.Component,
		Sort:        req.Sort,
		KeepAlive:   req.KeepAlive,
		DefaultMenu: req.DefaultMenu,
		Title:       req.Title,
		Icon:        req.Icon,
	}
	if err := global.DB.CreateMenu(&menu); err != nil {
		response.FailWithMessage(c, "创建菜单失败")
	} else {
		response.OkWithMessage(c, "创建成功")
	}
}

// @Tags Menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.DeleteMenu true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /menu/delete [post]
func DeleteMenu(c *gin.Context) {
	var req dto.DeleteMenu
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	if err := service.DeleteByIds(entity.Menu{}, req.Ids); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage(c, "删除失败")
	} else {
		response.OkWithMessage(c, "删除成功")
	}
}

// @Tags Menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.UpdateApi true "菜单名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /menu/update [post]
func UpdateMenu(c *gin.Context) {
	var req dto.UpdateMenu
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	menu := entity.Menu{
		ID:          req.Id,
		MenuLevel:   req.MenuLevel,
		ParentId:    req.ParentId,
		Path:        req.Path,
		Name:        req.Name,
		Hidden:      req.Hidden,
		Component:   req.Component,
		Sort:        req.Sort,
		KeepAlive:   req.KeepAlive,
		DefaultMenu: req.DefaultMenu,
		Title:       req.Title,
		Icon:        req.Icon,
	}

	if err := service.UpdateById(&menu, &entity.Menu{}, req.Id); err != nil {
		global.Logger.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage(c, "修改失败")
	} else {
		response.OkWithMessage(c, "修改成功")
	}
}

// @Tags Menu
// @Summary 分页获取菜单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/tree [post]
func GetMenuTree(c *gin.Context) {
	var req dto.SearchMenu
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	if total, list, err := service.GetMenuTree(); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(c, "获取失败")
	} else {
		response.OkWithDetailed(c, response.PageResult{
			List:  list,
			Total: int64(total),
			Page:  req.Page,
			Size:  req.Size,
		}, "获取成功")
	}
}

// @Tags Menu
// @Summary 根据id获取菜单信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.DetailMenu true "id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/detail [post]
func GetMenuById(c *gin.Context) {
	var req dto.DetailMenu
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	api := entity.Menu{}
	err := service.DetailById(&api, req.Id)
	if err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(c, "获取失败")
	} else {
		response.OkWithData(c, api)
	}
}

// @Tags Menu
// @Summary 获取所有的菜单 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/all [post]
func GetAllMenus(c *gin.Context) {
	if menus, err := global.DB.AllMenus(); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(c, "获取失败")
	} else {
		response.OkWithDetailed(c, menus, "获取成功")
	}
}
