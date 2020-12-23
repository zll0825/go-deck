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

// @Tags DictType
// @Summary 创建dictType
// @Security DictTypeKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.CreateDictType true "dictType路径, dictType中文描述, dictType组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dictType/create [post]
func CreateDictType(c *gin.Context) {
	var req dto.CreateDictType
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	dictType := entity.DictType{
		Name:      req.Name,
		Type:      req.Type,
		Desc:      req.Desc,
		Status:    0,
	}

	if err := global.DB.CreateDictType(&dictType); err != nil {
		global.Logger.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage(c, "创建失败")
	} else {
		response.OkWithMessage(c, "创建成功")
	}
}

// @Tags DictType
// @Summary 删除dictType
// @Security DictTypeKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.DeleteDictType true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dictType/delete [post]
func DeleteDictType(c *gin.Context) {
	var req dto.DeleteDictType
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	if err := service.DeleteByIds(entity.DictType{}, req.Ids); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage(c, "删除失败")
	} else {
		response.OkWithMessage(c, "删除成功")
	}
}

// @Tags DictType
// @Summary 更新dictType
// @Security DictTypeKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.UpdateDictType true "dictType路径, dictType中文描述, dictType组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /dictType/update [post]
func UpdateDictType(c *gin.Context) {
	var req dto.UpdateDictType
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	dictType := entity.DictType{
		ID:        req.Id,
		Name:      req.Name,
		Type:      req.Type,
		Desc:      req.Desc,
		Status:    0,
	}

	if err := service.UpdateById(&dictType, &entity.DictType{}, req.Id); err != nil {
		global.Logger.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage(c, "修改失败")
	} else {
		response.OkWithMessage(c, "修改成功")
	}
}

// @Tags DictType
// @Summary 分页获取API列表
// @Security DictTypeKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.SearchDictType true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dictType/list [post]
func GetDictTypeList(c *gin.Context) {
	var req dto.SearchDictType
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	if total, list, err := service.GetDictTypeList(req); err != nil {
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

// @Tags DictType
// @Summary 根据id获取dictType
// @Security DictTypeKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.DetailDictType true "根据id获取dictType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dictType/detail [post]
func GetDictTypeById(c *gin.Context) {
	var req dto.DetailDictType
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	dictType := entity.DictType{}
	err := service.DetailById(&dictType, req.Id)
	if err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(c, "获取失败")
	} else {
		response.OkWithData(c, dictType)
	}
}

// @Tags DictType
// @Summary 获取所有的DictType 不分页
// @Security DictTypeKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dictType/all [post]
func GetAllDictType(c *gin.Context) {
	if dictType, err := global.DB.AllDictTypes(); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(c, "获取失败")
	} else {
		response.OkWithDetailed(c, dictType, "获取成功")
	}
}
