package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
	"go-deck/app/middleware"
)

func InitDictRouter(router *gin.RouterGroup) {
	r := router.Group("dict").Use(middleware.OperationRecord())
	{
		r.POST("type/create", controller.CreateDictType)  // 创建字典类型
		r.POST("type/delete", controller.DeleteDictType)  // 删除字典类型
		r.POST("type/update", controller.UpdateDictType)  // 更新字典类型
		r.POST("type/list", controller.GetDictTypeList)   // 获取字典类型列表
		r.POST("type/detail", controller.GetDictTypeById) // 获取单条字典类型信息
		r.POST("type/all", controller.GetAllDictType)    // 获取所有字典类型


		r.POST("data/create", controller.CreateDictData)  // 创建字典数据
		r.POST("data/delete", controller.DeleteDictData)  // 删除字典数据
		r.POST("data/update", controller.UpdateDictData)  // 更新字典数据
		r.POST("data/list", controller.GetDictDataList)   // 获取字典数据列表
		r.POST("data/detail", controller.GetDictDataById) // 获取单条字典数据信息
	}
}
