package entity

import "time"

type DictType struct {
	ID     int    `gorm:"primarykey"`
	Name   string `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`
	Type   string `json:"type" form:"type" gorm:"column:type;comment:字典名（英）"`
	Desc   string `json:"desc" form:"desc" gorm:"column:desc;comment:描述"`
	Status int    `json:"status" form:"status" gorm:"column:status;comment:状态"`

	DictData []DictData `json:"sysDictionaryDetails" form:"sysDictionaryDetails"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
