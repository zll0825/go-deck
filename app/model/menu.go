package model

import (
	"go-deck/app/model/entity"
)

func (d *Dao) CreateMenu(menu *entity.Menu) error {
	err := SystemDB().Create(menu).Error
	return err
}

func (d *Dao) DeleteMenu(ids []int) error {
	err := SystemDB().Delete(entity.Menu{}, ids).Error
	return err
}

func (d *Dao) AllMenus() (data []entity.Menu, err error) {
	err = SystemDB().Find(&data).Error
	return
}