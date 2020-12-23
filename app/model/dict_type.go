package model

import "go-deck/app/model/entity"

func (d *Dao) CreateDictType(data *entity.DictType) error {
	err := SystemDB().Create(data).Error
	return err
}

func (d *Dao) AllDictTypes() (data []entity.DictType, err error) {
	err = SystemDB().Find(&data).Error
	return
}