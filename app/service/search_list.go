package service

import (
	"go-deck/app/controller/dto"
	"go-deck/app/global"
	"go-deck/app/model/entity"
	"go-deck/app/util"
)

func GetApiList(params dto.SearchApi) (total int64, data []entity.Api, err error) {
	db := global.DB.System.Model(entity.Api{})

	if params.Path != "" {
		db = db.Where("path LIKE ?", "%"+params.Path+"%")
	}

	if params.Description != "" {
		db = db.Where("description LIKE ?", "%"+params.Description+"%")
	}

	if params.Method != "" {
		db = db.Where("method = ?", params.Method)
	}

	if params.ApiGroup != "" {
		db = db.Where("api_group = ?", params.ApiGroup)
	}

	err = db.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	offset, limit := util.Pagination(params.Page, params.Size)
	err = db.Limit(limit).Offset(offset).Find(&data).Error

	return
}
