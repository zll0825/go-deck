package model

import "go-deck/app/model/entity"

func (d *Dao) CreateDictData(data *entity.DictData) error {
	err := SystemDB().Create(data).Error
	return err
}