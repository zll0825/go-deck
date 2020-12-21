package model

import (
	"go-deck/app/model/entity"
)

func (d *Dao) AllRoles() (data []entity.Role, err error) {
	err = SystemDB().Find(&data).Error
	return
}

func (d *Dao) FindRoleByRoleId(roleId int) (entity.Role, error) {
	var role entity.Role

	err := SystemDB().Model(entity.Role{}).Where("id = ?", roleId).Find(&role).Error

	return role, err
}

func (d *Dao) FindRoleKeyByRoleId(roleId int) string {
	var roleKey string

	err := SystemDB().Model(entity.Role{}).Where("id = ?", roleId).Pluck("key", &roleKey).Error

	if err != nil {
		return ""
	}

	return roleKey
}

func (d *Dao) FindRoleKeysByRoleIds(roleIds []int) []string {
	var roleKeys []string

	err := SystemDB().Model(entity.Role{}).Where("id in (?)", roleIds).Pluck("key", &roleKeys).Error

	if err != nil {
		return []string{}
	}

	return roleKeys
}

func (d *Dao) CreateRole(role *entity.Role) error {
	err := SystemDB().Create(role).Error
	return err
}
