package service

import (
	"go-deck/app/global"
	"go-deck/app/model/entity"
	"go-deck/app/util"
	"go-deck/pkg/casbin"
	"gorm.io/gorm"
	"strconv"
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

func BindRoleApi(roleId int, apiIds []int) error {
	// 查出apiInfo
	apiInfos, err := global.DB.FindApisByIds(apiIds)
	if err != nil {
		return err
	}

	roleIdString := strconv.Itoa(roleId)
	// 开启事务
	err = global.DB.System.Transaction(func(tx *gorm.DB) error {
		// 写入casbin
		cas, err := casbin.NewCasbin(global.Config.CasbinConfig, tx)
		if err != nil {
			return err
		}

		// 移除所有权限
		res, err := cas.RemoveFilteredPolicy(0, roleIdString)
		if err != nil {
			return err
		}

		var rules [][]string
		for _, info := range apiInfos {
			rules = append(rules, []string{roleIdString, info.Path, info.Method})
		}

		res, err = cas.AddPolicies(rules)
		if err != nil || res != true {
			return err
		}

		return nil
	})

	return err
}
