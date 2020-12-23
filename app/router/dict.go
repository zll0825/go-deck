package router

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller"
	"go-deck/app/middleware"
)

func InitDictRouter(Router *gin.RouterGroup) {
	router := Router.Group("dict").Use(middleware.OperationRecord())
	{
		router.POST("type/create", controller.CreateDictType)  // 创建字典类型
		router.POST("type/delete", controller.DeleteDictType)  // 删除字典类型
		router.POST("type/update", controller.UpdateDictType)  // 更新字典类型
		router.POST("type/list", controller.GetDictTypeList)   // 获取字典类型列表
		router.POST("type/detail", controller.GetDictTypeById) // 获取单条字典类型信息
		router.POST("type/all", controller.GetAllDictType)    // 获取所有字典类型


		router.POST("data/create", controller.CreateDictData)  // 创建字典数据
		router.POST("data/delete", controller.DeleteDictData)  // 删除字典数据
		router.POST("data/update", controller.UpdateDictData)  // 更新字典数据
		router.POST("data/list", controller.GetDictDataList)   // 获取字典数据列表
		router.POST("data/detail", controller.GetDictDataById) // 获取单条字典数据信息
	}
}
