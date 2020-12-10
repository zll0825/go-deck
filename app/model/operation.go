package model

import "go-deck/app/model/entity"

func (d *Dao)CreateOperationLog(operationLog entity.OperationLog) (err error) {
	err = SystemDB().Create(&operationLog).Error
	return err
}