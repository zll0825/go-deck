package service

import (
	"go-deck/app/global"
	"go-deck/app/util"
	"reflect"
)

func DeleteByIds(entity interface{}, ids []int) error {
	db := global.DB.System
	return db.Delete(entity, ids).Error
}

func DetailById(entity interface{}, id int) error {
	db := global.DB.System
	return db.Debug().Model(entity).Find(entity, id).Error
}

func UpdateById(new interface{}, old interface{}, id int) error {
	db := global.DB.System
	// 先根据id查出db的数据
	err := DetailById(old, id)
	if err != nil {
		return err
	}

	oldValue := reflect.ValueOf(old).Elem().Interface()
	newValue := reflect.ValueOf(new).Elem().Interface()
	_, after, err := util.DiffStruct(oldValue, newValue)
	if err != nil {
		return err
	}

	if len(after) == 0 {
		return nil
	}

	return db.Omit("").Model(new).Updates(after).Error
}
