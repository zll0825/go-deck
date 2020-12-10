package model

import (
	"go-deck/app/model/entity"
)

func (d *Dao) CreateRole(role *entity.Role) error {
	err := SystemDB().Create(role).Error
	return err
}