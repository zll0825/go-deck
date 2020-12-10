package model

import (
	"go-deck/app/model/entity"
)

func (d *Dao) FindMenuIdsByRoleId(roleId int) []int {
	var menuIds []int
	err := SystemDB().Model(entity.RoleMenu{}).Where("role_id = ?", roleId).Pluck("menu_id", &menuIds).Error
	if err != nil {
		return []int{}
	}

	return menuIds
}

func (d *Dao) CreateRoleMenu(roleMenu *entity.RoleMenu) error {
	err := SystemDB().Create(roleMenu).Error
	return err
}

func (d *Dao) DeleteRoleMenuByRoleIds(roleIds []int) error {
	err := SystemDB().Delete(entity.RoleMenu{}, roleIds).Error
	return err
}

func (d *Dao) BatchCreateRoleMenu(roleMenus []*entity.RoleMenu) error {
	err := SystemDB().Create(roleMenus).Error
	return err
}


