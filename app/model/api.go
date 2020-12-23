package model

import (
	"go-deck/app/model/entity"
)

func (d *Dao) CreateApi(api *entity.Api) error {
	err := SystemDB().Create(api).Error
	return err
}

func (d *Dao) AllApis() (data []entity.Api, err error) {
	err = SystemDB().Find(&data).Error
	return
}

func (d *Dao) FindApisByIds(ids []int) (data []entity.Api, err error) {
	err = SystemDB().Where("id in (?)", ids).Find(&data).Error
	return
}