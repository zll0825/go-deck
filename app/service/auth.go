package service

import (
	"go-deck/app/global"
	"go-deck/app/model/entity"
	"go-deck/app/util"
	"go-deck/pkg/casbin"
	"gorm.io/gorm"
)

func BindRoleMenu(roleId int, menuIds []int) error {
	// 查出现有roleId下的menuIds
	dbMenuIds := global.DB.FindMenuIdsByRoleId(roleId)

	// 对比menuIds和dbMenuIds得出要删的和要新增的
	del, add := util.GetIntDifference(dbMenuIds, menuIds)
	addItems := make([]*entity.RoleMenu, 0)
	for _, menuId := range add {
		addItems = append(addItems, &entity.RoleMenu{
			RoleID: roleId,
			MenuID: menuId,
		})
	}

	// 开启事务
	err := global.DB.System.Transaction(func(tx *gorm.DB) error {
		if len(del) > 0 {
			// 删除权限
			if err := tx.Where("role_id = ? and menu_id in (?)", roleId, del).Delete(entity.RoleMenu{}).Error;
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
	// 查出roleInfo
	role, err := global.DB.FindRoleByRoleId(roleId)
	if err != nil {
		return err
	}

	// 查出apiInfo
	apiInfos, err := global.DB.FindApisByIds(apiIds)
	if err != nil {
		return err
	}
	var apis []string
	for _, api := range apiInfos {
		apis = append(apis, api.Path)
	}

	// 查出现有角色下的api权限
	dbApis := global.DB.FindApisByRoleName(role.Name)

	// 对比menuIds和dbMenuIds得出要删的和要新增的
	del, add := util.GetStringDifference(dbApis, apis)

	// 开启事务
	err = global.DB.System.Transaction(func(tx *gorm.DB) error {
		var rules [][]string

		// 写入casbin
		cas, err := casbin.NewCasbin(global.Config.CasbinConfig, tx)
		if err != nil {
			return err
		}

		// 移除api
		if len(del) > 0 {
			for _, path := range del {
				rules = append(rules, []string{role.Key, path, "POST"})
			}
			_, err := cas.RemovePolicies(rules)
			if err != nil {
				return err
			}
		}

		// 增加api
		if len(add) > 0 {
			for _, path := range add {
				rules = append(rules, []string{role.Key, path, "POST"})
			}

			_, err = cas.AddPolicies(rules)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func BindUserRole(userId int, roleIds []int) error {
	// 查出用户
	user, err := global.DB.FindUserByUserId(userId)
	if err != nil {
		return err
	}

	// 查出现有username下的roleKeys
	dbRoleKeys := global.DB.FindRoleKeysByUsername(user.Username)
	roleKeys := global.DB.FindRoleKeysByRoleIds(roleIds)

	// 对比menuIds和dbMenuIds得出要删的和要新增的
	del, add := util.GetStringDifference(dbRoleKeys, roleKeys)

	// 开启事务
	err = global.DB.System.Transaction(func(tx *gorm.DB) error {
		var rules [][]string

		// 写入casbin
		cas, err := casbin.NewCasbin(global.Config.CasbinConfig, tx)
		if err != nil {
			return err
		}

		// 删除角色
		if len(del) > 0 {
			for _, delKey := range del {
				rules = append(rules, []string{user.Username, delKey})
			}
			if _, err := cas.RemoveNamedGroupingPolicies("g", rules); err != nil {
				return err
			}
		}

		// 增加角色
		if len(add) > 0 {
			for _, addKey := range add {
				rules = append(rules, []string{user.Username, addKey, ""})
			}
			if _, err := cas.AddNamedGroupingPolicies("g", rules); err != nil {
				return err
			}
		}

		// 返回 nil 提交事务
		return nil
	})

	return err
}
