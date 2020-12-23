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

func GetRoleList(params dto.SearchRole) (total int64, data []entity.Role, err error) {
	db := global.DB.System.Model(entity.Role{})

	if params.Name != "" {
		db = db.Where("name LIKE ?", "%"+params.Name+"%")
	}

	if params.Key != "" {
		db = db.Where("key LIKE ?", "%"+params.Key+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	offset, limit := util.Pagination(params.Page, params.Size)
	err = db.Limit(limit).Offset(offset).Find(&data).Error

	return
}

func GetUserList(params dto.SearchUser) (total int64, data []entity.User, err error) {
	db := global.DB.System.Model(entity.Role{})

	if params.Username != "" {
		db = db.Where("username LIKE ?", "%"+params.Username+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	offset, limit := util.Pagination(params.Page, params.Size)
	err = db.Limit(limit).Offset(offset).Find(&data).Error

	return
}

func GetDictTypeList(params dto.SearchDictType)(total int64, data []entity.DictType, err error) {
	db := global.DB.System.Model(entity.Role{})

	if params.Name != "" {
		db = db.Where("name like ?", "%"+params.Name+"%")
	}
	if params.Type != "" {
		db = db.Where("type like ?", "%"+params.Type+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	offset, limit := util.Pagination(params.Page, params.Size)
	err = db.Limit(limit).Offset(offset).Find(&data).Error

	return
}

func GetDictDataList(params dto.SearchDictData)(total int64, data []entity.DictData, err error) {
	db := global.DB.System.Model(entity.Role{})

	if params.TypeID > 0 {
		db = db.Where("type_id = ?", params.TypeID)
	}

	err = db.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	offset, limit := util.Pagination(params.Page, params.Size)
	err = db.Limit(limit).Offset(offset).Find(&data).Error

	return
}