package entity

import "time"

type DictData struct {
	ID          int    `gorm:"primarykey"`
	TypeID      int    `json:"type_id" form:"typeId" gorm:"column:type_id;comment:字典类型id"`
	Label       string `json:"label" form:"label" gorm:"column:label;comment:字典名（中）"`
	Value       string `json:"value" form:"value" gorm:"column:value;comment:字典名（英）"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述"`
	Status      int    `json:"status" form:"status" gorm:"column:status;comment:状态"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (DictData) TableName() string {
	return "dict_data"
}
