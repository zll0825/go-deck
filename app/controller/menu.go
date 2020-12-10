package controller

import (
	"github.com/gin-gonic/gin"
	"go-deck/app/controller/dto"
	"go-deck/app/global"
	"go-deck/app/model/entity"
	"go-deck/app/response"
)

// 获取用户授权菜单
func GetUserMenu(c *gin.Context) {

}

// 分页获取所有菜单
func GetMenuList()  {

}

// 新增菜单
func CreateMenu(c *gin.Context)  {
	var req dto.CreateMenu

	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage("参数验证失败", c)
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
		response.FailWithMessage("创建菜单失败", c)
	} else {
		response.OkWithMessage( "创建成功", c)
	}
}

// 更新菜单
func UpdateMenu(c *gin.Context)  {

}

// 删除菜单
func DeleteMenu()  {

}