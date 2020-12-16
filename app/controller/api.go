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

// @Tags Api
// @Summary 创建api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.CreateApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/create [post]
func CreateApi(c *gin.Context) {
	var req dto.CreateApi
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	api := entity.Api{
		Path:        req.Path,
		Description: req.Description,
		ApiGroup:    req.ApiGroup,
		Method:      req.Method,
	}

	if err := global.DB.CreateApi(&api); err != nil {
		global.Logger.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage(c, "创建失败")
	} else {
		response.OkWithMessage(c, "创建成功")
	}
}

// @Tags Api
// @Summary 删除api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.DeleteApi true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/delete [post]
func DeleteApi(c *gin.Context) {
	var req dto.DeleteApi
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	if err := service.DeleteByIds(entity.Api{}, req.Ids); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage(c, "删除失败")
	} else {
		response.OkWithMessage(c, "删除成功")
	}
}

// @Tags Api
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.SearchApi true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/list [post]
func GetApiList(c *gin.Context) {
	var req dto.SearchApi
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	if total, list, err := service.GetApiList(req); err != nil {
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

// @Tags Api
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.DetailApi true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/detail [post]
func GetApiById(c *gin.Context) {
	var req dto.DetailApi
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	api := entity.Api{}
	err := service.DetailById(&api, req.Id)
	if err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(c, "获取失败")
	} else {
		response.OkWithData(c, api)
	}
}

// @Tags Api
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.UpdateApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /api/update [post]
func UpdateApi(c *gin.Context) {
	var req dto.UpdateApi
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	api := entity.Api{
		ID:          req.Id,
		Path:        req.Path,
		Description: req.Description,
		ApiGroup:    req.ApiGroup,
		Method:      req.Method,
	}

	if err := service.UpdateById(&api, &entity.Api{}, req.Id); err != nil {
		global.Logger.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage(c, "修改失败")
	} else {
		response.OkWithMessage(c, "修改成功")
	}
}

// @Tags Api
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/all [post]
func GetAllApis(c *gin.Context) {
	if apis, err := global.DB.AllApis(); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(c, "获取失败")
	} else {
		response.OkWithDetailed(c, apis, "获取成功")
	}
}
