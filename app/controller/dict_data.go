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

// @Tags DictData
// @Summary 创建dictData
// @Security DictDataKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.CreateDictData true "dictData路径, dictData中文描述, dictData组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dictData/create [post]
func CreateDictData(c *gin.Context) {
	var req dto.CreateDictData
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	dictData := entity.DictData{
		TypeID:      req.TypeId,
		Label:       req.Label,
		Value:       req.Value,
		Description: req.Description,
		Status:      0,
	}

	if err := global.DB.CreateDictData(&dictData); err != nil {
		global.Logger.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage(c, "创建失败")
	} else {
		response.OkWithMessage(c, "创建成功")
	}
}

// @Tags DictData
// @Summary 删除dictData
// @Security DictDataKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.DeleteDictData true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dictData/delete [post]
func DeleteDictData(c *gin.Context) {
	var req dto.DeleteDictData
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	if err := service.DeleteByIds(entity.DictData{}, req.Ids); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage(c, "删除失败")
	} else {
		response.OkWithMessage(c, "删除成功")
	}
}

// @Tags DictData
// @Summary 更新dictData
// @Security DictDataKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.UpdateDictData true "dictData路径, dictData中文描述, dictData组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /dictData/update [post]
func UpdateDictData(c *gin.Context) {
	var req dto.UpdateDictData
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	dictData := entity.DictData{
		ID:          req.Id,
		TypeID:      req.TypeId,
		Label:       req.Label,
		Value:       req.Value,
		Description: req.Description,
		Status:      0,
	}

	if err := service.UpdateById(&dictData, &entity.DictData{}, req.Id); err != nil {
		global.Logger.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage(c, "修改失败")
	} else {
		response.OkWithMessage(c, "修改成功")
	}
}

// @Tags DictData
// @Summary 分页获取API列表
// @Security DictDataKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.SearchDictData true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dictData/list [post]
func GetDictDataList(c *gin.Context) {
	var req dto.SearchDictData
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	if total, list, err := service.GetDictDataList(req); err != nil {
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

// @Tags DictData
// @Summary 根据id获取dictData
// @Security DictDataKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.DetailDictData true "根据id获取dictData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dictData/detail [post]
func GetDictDataById(c *gin.Context) {
	var req dto.DetailDictData
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	dictData := entity.DictData{}
	err := service.DetailById(&dictData, req.Id)
	if err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(c, "获取失败")
	} else {
		response.OkWithData(c, dictData)
	}
}
