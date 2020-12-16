package model

import (
	"go-deck/app/model/entity"
)

func (d *Dao) FindRoleKeysByUsername(username string) []string {
	var roleKeys []string
	err := SystemDB().Model(entity.CasbinRule{}).Where("p_type = 'g' and v0 = ?", username).Pluck("v1", &roleKeys).Error
	if err != nil {
		return []string{}
	}

	return roleKeys
}

func (d *Dao) FindApisByRoleName(roleName string) []string {
	var apis []string
	err := SystemDB().Model(entity.CasbinRule{}).Where("p_type = 'p' and v0 = ?", roleName).Pluck("v1", &apis).Error
	if err != nil {
		return []string{}
	}

	return apis
}

func (d *Dao) CreateUserRole(roleMenu *entity.RoleMenu) error {
	err := SystemDB().Create(roleMenu).Error
	return err
}

func (d *Dao) DeleteUserRoleByUserIds(roleIds []int) error {
	err := SystemDB().Delete(entity.RoleMenu{}, roleIds).Error
	return err
}

func (d *Dao) BatchCreateUserRole(roleMenus []*entity.RoleMenu) error {
	err := SystemDB().Create(roleMenus).Error
	return err
}


