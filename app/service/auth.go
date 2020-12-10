package service

import (
	"go-deck/app/global"
	"go-deck/app/model/entity"
	"go-deck/app/util"
	"gorm.io/gorm"
)

func BindRoleMenu(roleId int, menuIds []int) error {
	// 查出现有roleId下的menuIds
	dbMenuIds := global.DB.FindMenuIdsByRoleId(roleId)

	// 对比menuIds和dbMenuIds得出要删的和要新增的
	delIds, addIds := util.GetDifference(dbMenuIds, menuIds)
	addItems := make([]*entity.RoleMenu, 0)
	for _, menuId := range addIds {
		addItems = append(addItems, &entity.RoleMenu{
			RoleID: roleId,
			MenuID: menuId,
		})
	}

	// 开启事务
	err := global.DB.System.Transaction(func(tx *gorm.DB) error {
		if len(delIds) > 0 {
			// 删除权限
			if err := tx.Where("role_id = ? and menu_id in (?)", roleId, delIds).Delete(entity.RoleMenu{}).Error;
			err != nil {
				return err
			}
		}

		if len(addItems) > 0 {
			// 增加权限
			if err := tx.Create(&addItems).Error; err != nil {
				return err
			}
		}

		// 返回 nil 提交事务
		return nil
	})

	return err
}
